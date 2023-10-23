package service

import (
	"github.com/sekthor/songbird-backend/pkg/model"
	"github.com/sekthor/songbird-backend/pkg/repo"
	"gorm.io/gorm"
)

type VenueService struct {
	repo repo.VenueRepo
}

func NewVenueService(db *gorm.DB) VenueService {
	return VenueService{
		repo: repo.NewVenueRepo(db),
	}
}

func (s *VenueService) GetById(id int) (model.Venue, error) {
	return s.repo.FetchById(id)
}

func (s *VenueService) GetAll() []model.Venue {
	return s.repo.FetchAll()
}

func (s *VenueService) Create(venue model.Venue) (model.Venue, error) {
	return s.repo.Create(venue)
}

func (s *VenueService) DeleteById(id int) error {
	return s.repo.DeleteById(id)
}

func (s *VenueService) AddTimeslot(venueId int, slot model.Timeslot) error {
	slot.VenueID = uint(venueId)
	_, err := s.repo.CreateTimeSlot(slot)
	return err
}

func (s *VenueService) DeleteTimeslot(tsid int) error {
	return s.repo.DeleteTimeslot(tsid)
}

func (s *VenueService) GetTimeslot(tsid int) (model.Timeslot, error) {
	return s.repo.FetchTimeslotById(tsid)
}
