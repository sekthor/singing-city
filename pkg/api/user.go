package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sekthor/songbird-backend/pkg/model"
)

func (api *api) Register(c *gin.Context) {

	var registerRequest struct {
		model.User
		Name    string `json:"name"`
		Address string `json:"address"`
		ZipCode int    `json:"zip"`
		City    string `json:"city"`
	}

	if c.BindJSON(&registerRequest) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data in login form"})
		return
	}

	user, err := api.userService.Register(registerRequest.User)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch user.Type {
	case 1:
		artist := model.Artist{}
		artist.ID = user.ID
		artist.Name = registerRequest.Name
		artist.User = user
		artist.Contact = user.Email
		artist, err = api.artistService.Create(artist)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, &artist)
	case 2:
		venue := model.Venue{}
		venue.ID = user.ID
		venue.Name = registerRequest.Name
		venue.User = user
		venue.Address = registerRequest.Address
		venue.ZipCode = registerRequest.ZipCode
		venue.City = registerRequest.City
		venue.Contact = user.Email
		venue, err = api.venueService.Create(venue)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, &venue)
	default:
		c.JSON(http.StatusBadRequest, gin.H{})
	}
}

func (api *api) Login(c *gin.Context) {
	var user model.User

	if c.BindJSON(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data in login form"})
		return
	}

	token, err := api.userService.Login(user)

	if err != nil {
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
		c.JSON(http.StatusNotFound, gin.H{"error": "no such user"})
		return
	}

	response.User = user.DTO()

	switch user.Type {
	case 1:
		artist, err := api.artistService.GetById(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "artist not found"})
			return
		}
		response.Artist = artist
	case 2:
		venue, err := api.venueService.GetById(id)
		if err != nil {
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	var user model.User
	if c.BindJSON(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data format"})
		return
	}

	user, err = api.userService.Update(id, user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not update artist"})
		return
	}

	c.JSON(http.StatusAccepted, &user)
}
