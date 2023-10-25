package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *api) GetArtistById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	venue, err := api.artistService.GetById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "artist not found"})
		return
	}

	c.JSON(http.StatusOK, &venue)

}
