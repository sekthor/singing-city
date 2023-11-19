package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (api *api) CreateInvite(c *gin.Context) {
	id, err := api.getUserIdFromContext(c)

	if err != nil || id != 1 {
		log.Debug().Err(err).Msgf("unauthorized")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	i, err := api.userService.CreateInvite()

	if err != nil {
		log.Debug().Err(err).Msgf("cloud not create invite")
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not create invite"})
		return
	}

	c.JSON(http.StatusOK, &i)
}

func (api *api) GetAllInvites(c *gin.Context) {
	id, err := api.getUserIdFromContext(c)

	if err != nil || id != 1 {
		log.Debug().Err(err).Msgf("unauthorized")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	i := api.userService.GetAllInvites()

	c.JSON(http.StatusOK, &i)
}
