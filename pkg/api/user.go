package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/sekthor/songbird-backend/pkg/model"
)

func (api *api) Register(c *gin.Context) {

	var registerRequest struct {
		model.User
		Name        string `json:"name"`
		Address     string `json:"address"`
		ZipCode     int    `json:"zip"`
		City        string `json:"city"`
		Phone       string `json:"phone"`
		Genere      string `json:"genere"`
		Description string `json:"description"`
	}

	if err := c.BindJSON(&registerRequest); err != nil {
		log.Debug().Err(err).Msgf("could not unmarshall user in register request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data in login form"})
		return
	}

	if registerRequest.Name == "" {
		log.Debug().Msgf("missing field: name")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing field: name"})
		return
	}

	if registerRequest.Phone == "" {
		log.Debug().Msgf("missing field: phone")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing field: phone"})
		return
	}

	invite := c.Query("invite")

	log.Trace().Msg("attempting to register user")
	user, err := api.userService.Register(registerRequest.User, invite)

	if err != nil {
		log.Debug().Err(err).Msgf("could not register user")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Trace().Msg("attempting to determine user type")
	switch user.Type {
	case 1:
		log.Trace().Msg("user is of type 'artist'")
		artist := model.Artist{}
		artist.ID = user.ID
		artist.Name = registerRequest.Name
		artist.User = user
		artist.Contact = user.Email
		artist.Phone = registerRequest.Phone
		artist.Description = registerRequest.Description
		artist.Genere = registerRequest.Genere
		log.Trace().Msgf("attempting to create artist for user '%d'", user.ID)
		artist, err = api.artistService.Create(artist)
		if err != nil {
			log.Debug().Err(err).Msgf("could not register artist for user '%d'", user.ID)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, &artist)
	case 2:
		log.Trace().Msg("user is of type 'venue'")
		venue := model.Venue{}
		venue.ID = user.ID
		venue.Name = registerRequest.Name
		venue.User = user
		venue.Address = registerRequest.Address
		venue.ZipCode = registerRequest.ZipCode
		venue.City = registerRequest.City
		venue.Phone = registerRequest.Phone
		venue.Contact = user.Email
		venue.Description = registerRequest.Description
		log.Trace().Msgf("attempting to create venue for user '%d'", user.ID)
		venue, err = api.venueService.Create(venue)
		if err != nil {
			log.Debug().Err(err).Msgf("could not register venue for user '%d'", user.ID)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, &venue)
	default:
		log.Debug().Err(err).Msgf("invalid user type '%s'", user.Type)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user type"})
	}
}

func (api *api) Login(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		log.Debug().Err(err).Msgf("could not unmarshall login request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data in login form"})
		return
	}

	token, err := api.userService.Login(user)

	if err != nil {
		log.Debug().Err(err).Msgf("could not login user '%d'", user.ID)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24, "", "", false, false)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (api *api) GetProfile(c *gin.Context) {

	// get the userid
	id, err := api.getUserIdFromContext(c)
	if err != nil {
		log.Debug().Err(err).Msgf("no userid found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response struct {
		User   model.UserDTO `json:"user"`
		Artist model.Artist  `json:"artist,omitempty"`
		Venue  model.Venue   `json:"venue,omitempty"`
	}

	user, err := api.userService.GetById(id)
	if err != nil {
		log.Debug().Err(err).Msgf("user '%d' not found", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "no such user"})
		return
	}

	response.User = user.DTO()

	switch user.Type {
	case 1:
		artist, err := api.artistService.GetById(id)
		if err != nil {
			log.Debug().Err(err).Msgf("artist '%d' not found", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "artist not found"})
			return
		}
		response.Artist = artist
	case 2:
		venue, err := api.venueService.GetById(id)
		if err != nil {
			log.Debug().Err(err).Msgf("venue '%d' not found", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "venue not found"})
			return
		}
		response.Venue = venue
	}

	c.JSON(http.StatusOK, &response)
}

func (api *api) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userid"))

	if err != nil {
		log.Debug().Err(err).Msgf("invalid user id: '%s'", c.Param("userId"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	var user model.User
	if c.BindJSON(&user) != nil {
		log.Debug().Err(err).Msgf("could not unmarshall user")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data format"})
		return
	}

	user, err = api.userService.Update(id, user)

	if err != nil {
		log.Debug().Err(err).Msgf("could not update user '%d'", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not update user"})
		return
	}

	c.JSON(http.StatusAccepted, &user)
}
