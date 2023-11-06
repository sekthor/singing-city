package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sekthor/songbird-backend/pkg/model"
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

func (api *api) UpdateArtist(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userid"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	var artist model.Artist
	if c.BindJSON(&artist) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data format"})
		return
	}

	artist, err = api.artistService.Update(id, artist)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not update artist"})
		return
	}

	c.JSON(http.StatusAccepted, &artist)
}
