package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/sekthor/songbird-backend/pkg/model"
)

func (api *api) GetAdminInfo(c *gin.Context) {
	userId, err := api.getUserIdFromContext(c)

	if err != nil || userId != 1 {
		log.Debug().Err(err).Msgf("user is not allowed to access admin info")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var adminInfo struct {
		Confirmed []model.Timeslot    `json:"confirmed"`
		Pending   []model.Application `json:"pending"`
		Venues    []model.Venue       `json:"venues"`
		Artists   []model.Artist      `json:"artists"`
		Invites   []model.Invite      `json:"invites"`
	}

	ts, err := api.venueService.GetAllConfirmedTimeslots()
	if err != nil {
		log.Debug().Err(err).Msgf("could not get confirmed timeslots")
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not get admin info"})
		return
	}

	venues := api.venueService.GetAllWithoutTimeslot()
	artists := api.artistService.GetAll()
	applications := api.applicationService.GetAllPending()
	invites := api.userService.GetAllInvites()

	adminInfo.Confirmed = ts
	adminInfo.Venues = venues
	adminInfo.Artists = artists
	adminInfo.Pending = applications
	adminInfo.Invites = invites

	c.JSON(http.StatusOK, &adminInfo)
}
