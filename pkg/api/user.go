package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sekthor/songbird-backend/pkg/model"
)

func (api *api) Signup(c *gin.Context) {

	var user model.User

	if c.BindJSON(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data in login form"})
		return
	}

	user, err := api.userService.Create(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
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
	c.SetCookie("Authorization", token, 3600*24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
