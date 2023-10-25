package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sekthor/songbird-backend/pkg/model"
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

func (api *api) AddTimeslot(c *gin.Context) {
	var slot model.Timeslot
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	if c.BindJSON(&slot) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid timeslot"})
		return
	}

	err = api.venueService.AddTimeslot(id, slot)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusAccepted)
}

func (api *api) DeleteTimeslot(c *gin.Context) {
	tsid, err := strconv.Atoi(c.Param("tsid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	err = api.venueService.DeleteTimeslot(tsid)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "could not delete timeslot"})
		return
	}

	c.Status(http.StatusOK)
}

func (api *api) GetTimeslots(c *gin.Context) {
	userId, err := api.getUserIdFromContext(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	timeslots, err := api.venueService.GetTimeslotsByUserId(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not get timestamps for user"})
		return
	}

	c.JSON(http.StatusOK, &timeslots)
}
