package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/sekthor/songbird-backend/pkg/model"
)

func (api *api) GetArtistById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Debug().Err(err).Msgf("invalid user id: '%s'", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	venue, err := api.artistService.GetById(id)

	if err != nil {
		log.Debug().Err(err).Msgf("could not find user '%d'", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "artist not found"})
		return
	}

	c.JSON(http.StatusOK, &venue)

}

func (api *api) UpdateArtist(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userid"))

	if err != nil {
		log.Debug().Err(err).Msgf("invalid user id: '%s'", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	var artist model.Artist
	if c.BindJSON(&artist) != nil {
		log.Debug().Err(err).Msgf("could not unmarshall artist")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data format"})
		return
	}

	artist, err = api.artistService.Update(id, artist)

	if err != nil {
		log.Debug().Err(err).Msgf("cloud not update user '%d'", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not update artist"})
		return
	}

	c.JSON(http.StatusAccepted, &artist)
}
