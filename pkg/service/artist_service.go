package service

import (
	"github.com/sekthor/singing-city/pkg/model"
	"github.com/sekthor/singing-city/pkg/repo"
	"gorm.io/gorm"
)

type ArtistService struct {
	repo repo.ArtistRepo
}

func NewArtistService(db *gorm.DB) ArtistService {
	return ArtistService{
		repo: repo.NewArtistRepo(db),
	}
}

func (s *ArtistService) GetById(id int) (model.Artist, error) {
	return s.repo.FetchById(id)
}

func (s *ArtistService) GetAll() []model.Artist {
	return s.repo.FetchAll()
}

func (s *ArtistService) Create(artist model.Artist) (model.Artist, error) {
	return s.repo.Create(artist)
}

func (s *ArtistService) Update(id int, artist model.Artist) (model.Artist, error) {
	artist.ID = uint(id)
	return s.repo.Save(artist)
}

func (s *ArtistService) DeleteById(id int) error {
	return s.repo.DeleteById(id)
}
