package service

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"github.com/sekthor/singing-city/pkg/middleware"
	"github.com/sekthor/singing-city/pkg/model"
	"github.com/sekthor/singing-city/pkg/repo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	repo            repo.UserRepo
	notify          NotificationService
	FrontendBaseUrl string
}

func NewUserService(db *gorm.DB, notify *NotificationService, frontendbaseurl string) UserService {
	return UserService{
		repo:            repo.NewUserRepo(db),
		notify:          *notify,
		FrontendBaseUrl: frontendbaseurl,
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

func (s *UserService) Update(id int, user model.User) (model.User, error) {
	user.ID = uint(id)
	if err := user.ValidateIngorePassword(); err != nil {
		return user, err
	}
	return s.repo.SaveOmitPassword(user)
}

func (s *UserService) Register(user model.User, invite string) (model.User, error) {

	var err error

	if inv, err := s.repo.FetchInviteById(invite); err != nil {
		log.Debug().Msgf("cloud find invite '%s'", inv)
		return user, ErrorInviteNotFound
	}

	if err = user.Validate(); err != nil {
		log.Trace().Msgf("cloud not validate user '%d'", user.ID)
		return user, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		log.Trace().Msgf("cloud not generate bcrypt hash for user '%d'", user.ID)
		return user, ErrorCouldNotHashPassword
	}

	user.Password = string(hash)

	if user, err = s.repo.Create(user); err != nil {
		log.Trace().Msgf("cloud not create user '%d'", user.ID)
		return user, err
	}

	s.repo.DeleteInviteById(invite)

	if err = s.notify.SendRegisterMessage(user, s.FrontendBaseUrl); err != nil {
		log.Trace().Err(err).Msgf("could not send register email to user '%d'", user.ID)
	} else {
		log.Trace().Msgf("sent register email to user '%d'", user.ID)
	}

	return user, nil
}

func (s *UserService) EnsureAdminUser(password string) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return ErrorCouldNotHashPassword
	}

	var user model.User
	user.ID = 1
	user.Username = "admin"
	user.Email = "admin@songbird.ch"
	user.Password = string(hash)
	user.Type = 0

	_, err = s.repo.Save(user)

	return err
}

func (s *UserService) DeleteById(id int) error {
	return s.repo.DeleteById(id)
}

func (s *UserService) CreateInvite() (model.Invite, error) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	invite := model.Invite{
		Invite: string(b),
	}
	return s.repo.CreateInvite(invite)
}

func (s *UserService) GetAllInvites() []model.Invite {
	return s.repo.FetchAllInvites()
}

func (s *UserService) ForgotPassword(email string) error {

	user, err := s.repo.FetchByEmail(email)
	if err != nil {
		return err
	}

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 64)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	resetRequest := model.PasswordReset{
		UserID: user.ID,
		Time:   time.Now(),
		Code:   string(b),
	}

	resetRequest, err = s.repo.CreatePasswordResetRequest(resetRequest)

	if err != nil {
		return err
	}

	params := MessageParams{
		Username: user.Username,
		Code:     resetRequest.Code,
		BaseUrl:  s.FrontendBaseUrl,
	}

	err = s.notify.SendPasswordResetLink(user.Email, params)

	return err
}

func (s *UserService) ResetPassword(code string, password string) error {

	resetRequest, err := s.repo.FetchPasswordResetRequestByCode(code)
	if err != nil {
		return errors.New("invalid reset request code")
	}

	if resetRequest.Time.Before(time.Now().Add(time.Minute * -15)) {
		s.repo.DeletePasswordResetRequestsByUserID(resetRequest.UserID)
		return errors.New("reset request has expired")
	}

	if err := s.repo.DeletePasswordResetRequestsByUserID(resetRequest.UserID); err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		log.Trace().Msgf("cloud not generate bcrypt hash for user '%d'", resetRequest.UserID)
		return ErrorCouldNotHashPassword
	}

	return s.repo.SetNewPasswordForUser(resetRequest.UserID, string(hash))
}
