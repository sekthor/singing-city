package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sekthor/songbird-backend/pkg/model"
	"github.com/sekthor/songbird-backend/pkg/service"
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

func (api *api) GetApplicationsOfUser(c *gin.Context) {

	var applications []model.Application

	status := c.Query("status")
	usertype := c.Param("usertype")
	id, err := strconv.Atoi(c.Param("userid"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	switch usertype {
	case "artist":
		applications, err = api.applicationService.GetApplicationsByArtist(id, status)
	case "venue":
		applications, err = api.applicationService.GetApplicationsByVenue(id, status)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user type"})
		return
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "could not find applications"})
		return
	}

	c.JSON(http.StatusOK, &applications)
}

func (api *api) AcceptApplication(c *gin.Context) {
	userid, err := api.getUserIdFromContext(c)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	applicationid, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	err = api.applicationService.AcceptApplication(applicationid, userid)

	if err != nil {
		if errors.Is(err, service.ErrorUnauthorized) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "could not accept application"})
		return
	}

	c.Status(http.StatusAccepted)
}

func (api *api) DeleteApplication(c *gin.Context) {

	userid, err := api.getUserIdFromContext(c)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	applicationid, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	err = api.applicationService.DeleteById(applicationid, userid)

	if err != nil {
		if errors.Is(err, service.ErrorUnauthorized) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "could not accept application"})
		return
	}

	c.Status(http.StatusOK)
}
