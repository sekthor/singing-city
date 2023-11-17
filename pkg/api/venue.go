package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/sekthor/songbird-backend/pkg/model"
)

func (api *api) GetAllVenues(c *gin.Context) {
	venues := api.venueService.GetAll()
	c.JSON(http.StatusOK, &venues)
}

func (api *api) GetVenueByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Debug().Err(err).Msgf("invalid venue id: '%s'", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	venue, err := api.venueService.GetById(id)

	if err != nil {
		log.Debug().Err(err).Msgf("cloud not find venue '%d'", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "venue not found"})
		return
	}

	c.JSON(http.StatusOK, &venue)
}

func (api *api) UpdateVenue(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("userid"))

	if err != nil {
		log.Debug().Err(err).Msgf("invalid venue id: '%s'", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	var venue model.Venue
	if c.BindJSON(&venue) != nil {
		log.Debug().Err(err).Msgf("could not unmarshall venue")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data format"})
		return
	}

	venue, err = api.venueService.Update(id, venue)
	if err != nil {
		log.Debug().Err(err).Msgf("cloud not update venue '%d'", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (api *api) AddTimeslot(c *gin.Context) {
	var slot model.Timeslot
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		log.Debug().Err(err).Msgf("invalid venue id: '%s'", c.Param("userid"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	if c.BindJSON(&slot) != nil {
		log.Debug().Err(err).Msgf("could not unmarshall timeslot")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid timeslot"})
		return
	}

	err = api.venueService.AddTimeslot(id, slot)

	if err != nil {
		log.Debug().Err(err).Msgf("cloud not add new timeslot to venue '%d'", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusAccepted)
}

func (api *api) DeleteTimeslot(c *gin.Context) {
	tsid, err := strconv.Atoi(c.Param("tsid"))
	if err != nil {
		log.Debug().Err(err).Msgf("invalid timeslot id: '%s'", c.Param("tsid"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	err = api.venueService.DeleteTimeslot(tsid)

	if err != nil {
		log.Debug().Err(err).Msgf("cloud not delete timeslot '%d'", tsid)
		c.JSON(http.StatusNotFound, gin.H{"error": "could not delete timeslot"})
		return
	}

	c.Status(http.StatusOK)
}

func (api *api) DeleteTimeslotAsAdmin(c *gin.Context) {

	id, err := api.getUserIdFromContext(c)

	if err != nil || id != 1 {
		log.Debug().Err(err).Msgf("user is no authorized admin")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	tsid, err := strconv.Atoi(c.Param("tsid"))
	if err != nil {
		log.Debug().Err(err).Msgf("invalid timeslot id: '%s'", c.Param("tsid"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	err = api.venueService.DeleteTimeslot(tsid)

	if err != nil {
		log.Debug().Err(err).Msgf("cloud not delete timeslot '%d'", tsid)
		c.JSON(http.StatusNotFound, gin.H{"error": "could not delete timeslot"})
		return
	}

	c.Status(http.StatusOK)
}

func (api *api) GetTimeslots(c *gin.Context) {
	userId, err := api.getUserIdFromContext(c)

	if err != nil {
		log.Debug().Err(err).Msgf("no userid found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	timeslots, err := api.venueService.GetTimeslotsByUserId(userId)

	if err != nil {
		log.Debug().Err(err).Msgf("cloud not delete timeslots for user '%d'", userId)
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not get timeslots for user"})
		return
	}

	c.JSON(http.StatusOK, &timeslots)
}
