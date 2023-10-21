package service

import (
	"github.com/sekthor/songbird-backend/pkg/model"
	"github.com/sekthor/songbird-backend/pkg/repo"
	"gorm.io/gorm"
)

type ApplicationService struct {
	repo repo.ApplicationRepo
}

func NewApplicationService(db *gorm.DB) ApplicationService {
	return ApplicationService{
		repo: repo.NewApplicationRepo(db),
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
