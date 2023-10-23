package service

import (
	"errors"

	"github.com/sekthor/songbird-backend/pkg/model"
	"github.com/sekthor/songbird-backend/pkg/repo"
	"gorm.io/gorm"
)

var (
	ErrorArtistNotExist = errors.New("artist does not exist")
	ErrorSlotNotExist   = errors.New("timeslot does not exist")
	ErrorInvalidStatus  = errors.New("invalid status")
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

func (s *ApplicationService) DeleteById(id int) error {
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
