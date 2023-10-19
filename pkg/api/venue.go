package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *api) GetAllVenues(c *gin.Context) {
	venues := api.venueService.GetAll()
	c.JSON(http.StatusOK, &venues)
}

func (api *api) GetVenueByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	venue, err := api.venueService.GetById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "venue not found"})
		return
	}

	c.JSON(http.StatusOK, &venue)
}
