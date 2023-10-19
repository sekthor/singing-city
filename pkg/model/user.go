package model

import (
	"errors"
	"regexp"

	"gorm.io/gorm"
)

var (
	ErrorMissingEmail    = errors.New("missing email address")
	ErrorInvalidEmail    = errors.New("invalid email address format")
	ErrorMissingPassword = errors.New("missing password")
	ErrorPasswordLength  = errors.New("password is too short")
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (u *User) Validate() error {

	if u.Email == "" {
		return ErrorMissingPassword
	}

	if ok, err := regexp.MatchString("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,}$", u.Email); !ok || err != nil {
		return ErrorInvalidEmail
	}

	if u.Password == "" {
		return ErrorMissingPassword
	}

	if len(u.Password) < 8 {
		return ErrorPasswordLength
	}

	return nil
}
