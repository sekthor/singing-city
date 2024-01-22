package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/sekthor/singing-city/pkg/model"
	"github.com/sekthor/singing-city/pkg/service"
)

func (api *api) Apply(c *gin.Context) {

	var err error
	var userId, timeslotId int

	if userId, err = strconv.Atoi(c.Param("userid")); err != nil && userId <= 0 {
		log.Debug().Err(err).Msgf("invalid user id: '%s'", c.Param("userId"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	if timeslotId, err = strconv.Atoi(c.Param("tsid")); err != nil && timeslotId <= 0 {
		log.Debug().Err(err).Msgf("invalid timeslot id: '%s'", c.Param("tsid"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid timeslotId id"})
		return
	}

	err = api.applicationService.Apply(userId, timeslotId)

	if err != nil {
		log.Debug().Err(err).Msgf("user '%d' could not apply for timeslot '%d'", userId, timeslotId)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Debug().Msgf("user '%d' applied for timeslot '%d'", userId, timeslotId)
	c.Status(http.StatusAccepted)
}

func (api *api) GetApplicationsOfUser(c *gin.Context) {

	var applications []model.Application

	status := c.Query("status")
	usertype := c.Param("usertype")
	id, err := strconv.Atoi(c.Param("userid"))

	if err != nil {
		log.Debug().Err(err).Msgf("invalid user id: '%s'", c.Param("userId"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	switch usertype {
	case "artist":
		log.Debug().Msgf("fetching artist profile for user with id: '%d'", id)
		applications, err = api.applicationService.GetApplicationsByArtist(id, status)
	case "venue":
		log.Debug().Msgf("fetching venue profile for user with id: '%d'", id)
		applications, err = api.applicationService.GetApplicationsByVenue(id, status)
	default:
		log.Debug().Err(err).Msgf("invalid user type '%s'", usertype)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user type"})
		return
	}

	if err != nil {
		log.Debug().Err(err).Msgf("could not fetch %s profile for user %d", usertype, id)
		c.JSON(http.StatusNotFound, gin.H{"error": "could not find applications"})
		return
	}

	c.JSON(http.StatusOK, &applications)
}

func (api *api) AcceptApplication(c *gin.Context) {
	userid, err := api.getUserIdFromContext(c)
	if err != nil {
		log.Debug().Err(err).Msgf("no userid found in context")
		c.Status(http.StatusUnauthorized)
		return
	}

	applicationid, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Debug().Err(err).Msgf("invalid application id '%s'", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	err = api.applicationService.AcceptApplication(applicationid, userid)

	if err != nil {
		if errors.Is(err, service.ErrorUnauthorized) {
			log.Debug().Err(err).Msgf("user '%d' is not authorized to accept application '%d'", userid, applicationid)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		log.Debug().Err(err).Msgf("user '%d' could not accept application '%d'", userid, applicationid)
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not accept application"})
		return
	}

	c.Status(http.StatusAccepted)
}

func (api *api) DeleteApplication(c *gin.Context) {

	userid, err := api.getUserIdFromContext(c)
	if err != nil {
		log.Debug().Err(err).Msgf("no userid found in context")
		c.Status(http.StatusUnauthorized)
		return
	}

	applicationid, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Debug().Err(err).Msgf("invalid application id '%s'", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	err = api.applicationService.DeleteById(applicationid, userid)

	if err != nil {
		if errors.Is(err, service.ErrorUnauthorized) {
			log.Debug().Err(err).Msgf("user '%d' is not authorized to delete application '%d'", userid, applicationid)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		log.Debug().Err(err).Msgf("user '%d' could not delte application '%d'", userid, applicationid)
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not accept application"})
		return
	}

	c.Status(http.StatusOK)
}
