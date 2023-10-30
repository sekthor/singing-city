package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sekthor/songbird-backend/pkg/model"
)

func (api *api) Register(c *gin.Context) {

	var registerRequest struct {
		model.User
		Name string `json:"name"`
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
		venue.Contact = user.Email
		venue, err = api.venueService.Create(venue)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, &venue)
	default:
		c.JSON(http.StatusAccepted, gin.H{})
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
