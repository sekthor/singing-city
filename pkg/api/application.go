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

	if timeslotId, err = strconv.Atoi(c.Param("userid")); err != nil && timeslotId <= 0 {
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
