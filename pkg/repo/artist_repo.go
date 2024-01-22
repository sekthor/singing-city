package repo

import (
	"github.com/sekthor/singing-city/pkg/model"
	"gorm.io/gorm"
)

type ArtistRepo struct {
	db *gorm.DB
}

func NewArtistRepo(db *gorm.DB) ArtistRepo {
	return ArtistRepo{
		db: db,
	}
}

func (r *ArtistRepo) Create(artist model.Artist) (model.Artist, error) {
	result := r.db.Create(&artist)
	return artist, result.Error
}

func (r *ArtistRepo) Save(artist model.Artist) (model.Artist, error) {
	result := r.db.Where("artist_id = ?", artist.ID).Delete(&model.SocialLink{})
	result = r.db.Save(&artist)
	return artist, result.Error
}

func (r *ArtistRepo) FetchById(id int) (model.Artist, error) {
	var artist model.Artist
	result := r.db.
		Preload("Socials").
		First(&artist, id)
	return artist, result.Error
}

func (r *ArtistRepo) FetchAll() []model.Artist {
	var artists []model.Artist
	_ = r.db.Find(&artists)
	return artists
}

func (r *ArtistRepo) DeleteById(id int) error {
	return r.db.Delete(&model.Artist{}, id).Error
}
