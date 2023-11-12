package service

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sekthor/songbird-backend/pkg/middleware"
	"github.com/sekthor/songbird-backend/pkg/model"
	"github.com/sekthor/songbird-backend/pkg/repo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	repo repo.UserRepo
}

func NewUserService(db *gorm.DB) UserService {
	return UserService{
		repo: repo.NewUserRepo(db),
	}
}

func (s *UserService) Login(user model.User) (string, error) {

	u, err := s.repo.FetchByEmail(user.Email)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))

	if err != nil {
		return "", ErrorInvalidCredentials
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  strconv.Itoa(int(u.ID)),
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"name": u.Username,
		"type": u.Type,
	})

	tokenString, err := token.SignedString(middleware.ServerSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *UserService) GetByEmail(email string) (model.User, error) {
	return s.repo.FetchByEmail(email)
}

func (s *UserService) GetById(id int) (model.User, error) {
	return s.repo.FetchById(id)
}

func (s *UserService) Register(user model.User) (model.User, error) {

	var err error

	if err = user.Validate(); err != nil {
		return user, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		return user, ErrorCouldNotHashPassword
	}

	user.Password = string(hash)

	return s.repo.Create(user)
}

func (s *UserService) DeleteById(id int) error {
	return s.repo.DeleteById(id)
}
