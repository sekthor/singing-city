package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sekthor/songbird-backend/pkg/model"
)

func (api *api) Register(c *gin.Context) {

	var user model.User

	if c.BindJSON(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data in login form"})
		return
	}

	user, err := api.userService.Register(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch user.Type {
	case 1:
		artist := model.Artist{}
		artist.ID = user.ID
		artist.User = user
		artist, err = api.artistService.Create(artist)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, &artist)
	case 2:
		venue := model.Venue{}
		venue.ID = user.ID
		venue.User = user
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
