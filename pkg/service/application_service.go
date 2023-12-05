package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/sekthor/songbird-backend/pkg/model"
	"github.com/sekthor/songbird-backend/pkg/repo"
	"gorm.io/gorm"
)

type ApplicationService struct {
	repo       repo.ApplicationRepo
	artistRepo repo.ArtistRepo
	venueRepo  repo.VenueRepo
	notify     *NotificationService
}

func NewApplicationService(db *gorm.DB, notify *NotificationService) ApplicationService {
	return ApplicationService{
		repo:       repo.NewApplicationRepo(db),
		artistRepo: repo.NewArtistRepo(db),
		venueRepo:  repo.NewVenueRepo(db),
		notify:     notify,
	}
}

func (s *ApplicationService) GetById(id int) (model.Application, error) {
	return s.repo.FetchById(id)
}

func (s *ApplicationService) GetAllPending() []model.Application {
	return s.repo.FetchAllPending()
}

func (s *ApplicationService) Create(application model.Application) (model.Application, error) {
	return s.repo.Create(application)
}

func (s *ApplicationService) DeleteById(id int, userId int) error {
	app, err := s.repo.FetchById(id)
	if err != nil {
		return err
	}

	if app.ArtistID != uint(userId) && app.Timeslot.VenueID != uint(userId) && userId != 1 {
		return ErrorUnauthorized
	}

	if err = s.repo.DeleteById(id); err != nil {
		return err
	}

	// if deletion was initiated by venue, notify artist
	if uint(userId) == app.Timeslot.VenueID {
		venue, _ := s.venueRepo.FetchById(int(app.Timeslot.VenueID))
		artist, _ := s.artistRepo.FetchById(int(app.ArtistID))
        loc, _ := time.LoadLocation("Europe/Zurich")
        localTime := app.Timeslot.Time.In(loc)
		params := MessageParams{
			Username: artist.Name,
			Time:     localTime.Format("15:04"),
			Date:     localTime.Format("02.01.2006"),
			Venue:    venue.Name,
		}
		s.notify.SendRejectedMessage(artist.Contact, params)
	}

	return nil
}

func (s *ApplicationService) Apply(artistID int, timeslotID int) error {
	// validate artist exists
	artist, err := s.artistRepo.FetchById(artistID)

	if err != nil {
		return ErrorArtistNotExist
	}

	// validate timeslot exists
	slot, err := s.venueRepo.FetchTimeslotById(timeslotID)
	if err != nil {
		return ErrorSlotNotExist
	}

	// create application
	application := model.Application{
		ArtistID:   artist.ID,
		TimeslotID: slot.ID,
		Confirmed:  false,
	}

	apps, err := s.repo.FetchByArtistId(artistID)
	if err != nil {
		return err
	}
	for _, a := range apps {
		if a.TimeslotID == application.TimeslotID {
			return ErrorAlreadyApplied
		}
	}

	venue, err := s.venueRepo.FetchByIdWithUser(int(slot.VenueID))
	if err != nil {
		return err
	}

	application, err = s.repo.Create(application)

	if err != nil {
		return err
	}
    loc, _ := time.LoadLocation("Europe/Zurich")
    localTime := slot.Time.In(loc)
	params := MessageParams{
		Username: venue.User.Username,
		Artist:   artist.Name,
		Time:     localTime.Format("15:04"),
		Date:     localTime.Format("02.01.2006"),
	}
	s.notify.SendApplicationMessage(venue.Contact, params)

	return nil
}

func (s *ApplicationService) GetApplicationsByVenue(venueId int, status string) ([]model.Application, error) {

	var confirmed bool

	switch status {
	case "confirmed":
		confirmed = true
	case "open":
		confirmed = false
	case "":
		// return all ts of venue
		return s.repo.FetchByVenueId(venueId)
	default:
		return nil, ErrorInvalidStatus
	}
	// return all ts of venue with given status
	return s.repo.FetchByVenueIdAndStatus(venueId, confirmed)
}

func (s *ApplicationService) GetApplicationsByArtist(artistId int, status string) ([]model.Application, error) {

	var confirmed bool

	switch status {
	case "confirmed":
		confirmed = true
	case "open":
		confirmed = false
	case "":
		return s.repo.FetchByArtistId(artistId)
	default:
		return nil, ErrorInvalidStatus
	}
	// return all ts of venue with given status
	return s.repo.FetchByArtistIdAndStatus(artistId, confirmed)
}

func (s *ApplicationService) AcceptApplication(applicationId int, userId int) error {

	// find the application by id
	application, err := s.repo.FetchById(applicationId)
	if err != nil {
		return errors.New("could not find application with id")
	}

	// make sure userid matches the venueid -> user must be resource owner
	if application.Timeslot.VenueID != uint(userId) {
		return ErrorUnauthorized
	}

	// update the confirmation status
	application.Confirmed = true
	application, err = s.repo.Save(application)
	if err != nil {
		return errors.New("could not update application with id")
	}

	// set the artist to the timeslot
	application.Timeslot.ArtistID = application.ArtistID
	_, err = s.venueRepo.SaveTimeslot(application.Timeslot)
	if err != nil {
		return errors.New("could not update artist of timeslot")
	}

	// soft delete all other applications for this timeslot
	err = s.repo.DeleteByTimeslotIdExcept(int(application.TimeslotID), applicationId)
	if err != nil {
		return errors.New("could not delete remaining applications for same timeslot")
	}

	venue, err := s.venueRepo.FetchById(int(application.Timeslot.VenueID))
	if err != nil {
		return errors.New("could find venue")
	}

	artist, err := s.artistRepo.FetchById(int(application.ArtistID))
	if err != nil {
		return errors.New("could find artist")
	}

    loc, _ := time.LoadLocation("Europe/Zurich")
    localTime := application.Timeslot.Time.In(loc)

	params := MessageParams{
		Username: artist.Name,
		Contact:  venue.Contact,
		Venue:    venue.Name,
		Wage:     strconv.Itoa(application.Timeslot.Pay),
		Address:  fmt.Sprintf("%s, %d %s", venue.Address, venue.ZipCode, venue.City),
		Time:     localTime.Format("15:04"),
		Date:     localTime.Format("02.01.2006"),
	}

	s.notify.SendConfirmedMessage(artist.Contact, params)

	return nil
}
