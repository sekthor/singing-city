package repo

import (
	"github.com/sekthor/songbird-backend/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ApplicationRepo struct {
	db *gorm.DB
}

func NewApplicationRepo(db *gorm.DB) ApplicationRepo {
	return ApplicationRepo{
		db: db,
	}
}

func (r *ApplicationRepo) Create(application model.Application) (model.Application, error) {
	result := r.db.Create(&application)
	return application, result.Error
}

func (r *ApplicationRepo) FetchById(id int) (model.Application, error) {
	var application model.Application
	result := r.db.Preload("Timeslot").First(&application, id)
	return application, result.Error
}

func (r *ApplicationRepo) FetchByVenueId(venueId int) ([]model.Application, error) {
	var applications []model.Application
	result := r.db.
		Preload(clause.Associations).
		Where("timeslot_id IN (SELECT timeslots.ID FROM timeslots WHERE venue_id = ?)", venueId).
		Find(&applications)
	return applications, result.Error
}

func (r *ApplicationRepo) FetchByVenueIdAndStatus(venueId int, confirmed bool) ([]model.Application, error) {
	var applications []model.Application
	result := r.db.
		Preload(clause.Associations).
		Where("timeslot_id IN (SELECT timeslots.ID FROM timeslots WHERE venue_id = ?) AND confirmed = ?", venueId, confirmed).
		Find(&applications)

	return applications, result.Error
}

func (r *ApplicationRepo) FetchByArtistId(artistId int) ([]model.Application, error) {
	var applications []model.Application
	result := r.db.
		Preload(clause.Associations).
		Where("artist_id = ?", artistId).
		Find(&applications)

	return applications, result.Error
}

func (r *ApplicationRepo) FetchByArtistIdAndStatus(artistId int, confirmed bool) ([]model.Application, error) {
	var applications []model.Application
	result := r.db.
		Preload(clause.Associations).
		Where("artist_id = ? AND confirmed = ?", artistId, confirmed).
		Find(&applications)

	return applications, result.Error
}

func (r *ApplicationRepo) FetchAllPending() []model.Application {
	var applications []model.Application
	_ = r.db.
		Preload("Timeslot").
		Preload("Artist").
		Where("confirmed == false").
		Find(&applications)
	return applications
}

func (r *ApplicationRepo) DeleteById(id int) error {
	return r.db.Delete(&model.Application{}, id).Error
}

func (r *ApplicationRepo) Save(application model.Application) (model.Application, error) {
	result := r.db.Save(&application)
	return application, result.Error
}

func (r *ApplicationRepo) DeleteByTimeslotIdExcept(tsid int, exceptions ...int) error {
	var applications []model.Application
	result := r.db.
		Not(exceptions).
		Where("timeslot_id = ?", tsid).
		Delete(&applications)
	return result.Error
}
