package repo

import (
	"github.com/sekthor/songbird-backend/pkg/model"
	"gorm.io/gorm"
)

type VenueRepo struct {
	db *gorm.DB
}

func NewVenueRepo(db *gorm.DB) VenueRepo {
	return VenueRepo{
		db: db,
	}
}

func (r *VenueRepo) Create(venue model.Venue) (model.Venue, error) {
	result := r.db.Create(&venue)
	return venue, result.Error
}

func (r *VenueRepo) FetchById(id int) (model.Venue, error) {
	var venue model.Venue
	result := r.db.Preload("Slots").First(&venue, id)
	return venue, result.Error
}

func (r *VenueRepo) FetchAll() []model.Venue {
	var venues []model.Venue
	_ = r.db.Preload("Slots").Find(&venues)
	return venues
}

func (r *VenueRepo) DeleteById(id int) error {
	return r.db.Delete(&model.Venue{}, id).Error
}

func (r *VenueRepo) CreateTimeSlot(slot model.Timeslot) (model.Timeslot, error) {
	var result *gorm.DB

	// if artistid is not set, it's value is 0 which will result in a sql foreign key error
	// as there is not artist with id 0 -> ignore the artist_id if equal to 0
	if slot.ArtistID == 0 {
		result = r.db.Omit("artist_id").Create(&slot)
	} else {
		result = r.db.Create(&slot)
	}
	return slot, result.Error
}

func (r *VenueRepo) DeleteTimeslot(tsid int) error {
	return r.db.Delete(&model.Timeslot{}, tsid).Error
}

func (r *VenueRepo) FetchTimeslotById(tsid int) (model.Timeslot, error) {
	var ts model.Timeslot
	result := r.db.First(&ts, tsid)
	return ts, result.Error
}
