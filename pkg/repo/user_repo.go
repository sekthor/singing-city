package repo

import (
	"errors"

	"github.com/sekthor/songbird-backend/pkg/model"
	"gorm.io/gorm"
)

var (
	ErrorEmailExists = errors.New("email already exists")
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(user model.User) (model.User, error) {
	result := r.db.Create(&user)

	// TODO: actually check the nature of the error
	// for now we just assume this is why it fails
	if result.Error != nil {
		return user, ErrorEmailExists
	}

	return user, result.Error
}

func (r *UserRepo) FetchById(id int) (model.User, error) {
	var user model.User
	result := r.db.First(&user, id)
	return user, result.Error
}

func (r *UserRepo) FetchByEmail(email string) (model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	return user, result.Error
}

func (r *UserRepo) FetchAll() []model.User {
	var users []model.User
	_ = r.db.Find(&users)
	return users
}

func (r *UserRepo) DeleteById(id int) error {
	return r.db.Delete(&model.User{}, id).Error
}
