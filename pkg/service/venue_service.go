package service

import (
	"github.com/sekthor/singing-city/pkg/model"
	"github.com/sekthor/singing-city/pkg/repo"
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

func (s *VenueService) GetAllWithoutTimeslot() []model.Venue {
	return s.repo.FetchAll()
}

func (s *VenueService) Create(venue model.Venue) (model.Venue, error) {
	return s.repo.Create(venue)
}

func (s *VenueService) Update(id int, venue model.Venue) (model.Venue, error) {

	// make sure venue exists
	_, err := s.repo.FetchById(id)
	if err != nil {
		return venue, ErrorVenueNotExist
	}

	venue.ID = uint(id)

	venue, err = s.repo.Save(venue)
	if err != nil {
		return venue, ErrorUpdateVenueFailed
	}

	return venue, nil
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

func (s *VenueService) GetTimeslotsByUserId(userId int) ([]model.Timeslot, error) {
	return s.repo.FetchTimeslotByUserId(userId)
}

func (s *VenueService) GetAllConfirmedTimeslots() ([]model.Timeslot, error) {
	return s.repo.FetchAllConfirmedTimeslots()
}
