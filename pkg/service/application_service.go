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
	/*
		// validate artist exists
		artist, err := s.artistRepo.FetchById(artistID)

		if err != nil {
			return ErrorArtistNotExist
		}

		// TODO: validate timeslot exists
		slot, err := s.venueRepo.FetchById(timeslotID)

		// TODO: create application
		application := model.Application{
			ArtistID:   artist.ID,
			TimeslotID: slot.ID,
		}

		// TODO: save application
		application.TimeslotID = 1
	*/
	return nil
}
