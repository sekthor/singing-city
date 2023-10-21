package repo

import (
	"github.com/sekthor/songbird-backend/pkg/model"
	"gorm.io/gorm"
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
	result := r.db.First(&application, id)
	return application, result.Error
}

func (r *ApplicationRepo) FetchByVenueId(venueId int) (model.Application, error) {
	var application model.Application
	result := r.db.Where("venue_id = ?", venueId).Find(&application)
	return application, result.Error
}

func (r *ApplicationRepo) FetchAll() []model.Application {
	var applications []model.Application
	_ = r.db.Find(&applications)
	return applications
}

func (r *ApplicationRepo) DeleteById(id int) error {
	return r.db.Delete(&model.Application{}, id).Error
}
