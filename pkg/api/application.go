package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *api) Apply(c *gin.Context) {

	var err error
	var userId, timeslotId int

	if userId, err = strconv.Atoi(c.Param("userid")); err != nil && userId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	if timeslotId, err = strconv.Atoi(c.Param("tsid")); err != nil && timeslotId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid timeslotId id"})
		return
	}

	err = api.applicationService.Apply(userId, timeslotId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusAccepted)
}

func (api *api) GetApplicationsOfVenue(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}
	status := c.Query("status")

	applications, err := api.applicationService.GetApplicationsByVenue(id, status)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "could not find applications"})
		return
	}

	c.JSON(http.StatusOK, &applications)
}

func (api *api) GetApplicationsOfArtist(c *gin.Context) {

}
