package service

import (
	"errors"

	"github.com/sekthor/songbird-backend/pkg/model"
	"github.com/sekthor/songbird-backend/pkg/repo"
	"gorm.io/gorm"
)

type ApplicationService struct {
	repo       repo.ApplicationRepo
	artistRepo repo.ArtistRepo
	venueRepo  repo.VenueRepo
}

func NewApplicationService(db *gorm.DB) ApplicationService {
	return ApplicationService{
		repo:       repo.NewApplicationRepo(db),
		artistRepo: repo.NewArtistRepo(db),
		venueRepo:  repo.NewVenueRepo(db),
	}
}

func (s *ApplicationService) GetById(id int) (model.Application, error) {
	return s.repo.FetchById(id)
}

func (s *ApplicationService) GetAll() []model.Application {
	return s.repo.FetchAll()
}

func (s *ApplicationService) Create(application model.Application) (model.Application, error) {
	return s.repo.Create(application)
}

func (s *ApplicationService) DeleteById(id int, userId int) error {
	app, err := s.repo.FetchById(id)
	if err != nil {
		return err
	}

	if app.ArtistID != uint(userId) && app.Timeslot.VenueID != uint(userId) {
		return ErrorUnauthorized
	}

	return s.repo.DeleteById(id)
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

	_, err = s.repo.Create(application)

	return err
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

	return nil
}
