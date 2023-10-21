package model

import (
	"errors"
	"regexp"

	"gorm.io/gorm"
)

var (
	ErrorMissingEmail    = errors.New("missing email address")
	ErrorMissingUsername = errors.New("missing username")
	ErrorInvalidEmail    = errors.New("invalid email address format")
	ErrorInvalidUsername = errors.New("invalid username format: min 4 characters, can only contain letters and numbers")
	ErrorMissingPassword = errors.New("missing password")
	ErrorPasswordLength  = errors.New("password is too short")
)

type User struct {
	gorm.Model
	Username string `json:"username"`
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

	if u.Username == "" {
		return ErrorMissingUsername
	}

	if ok, err := regexp.MatchString("^[a-zA-Z0-9]{4,}", u.Username); !ok || err != nil {
		return ErrorInvalidUsername
	}

	if u.Password == "" {
		return ErrorMissingPassword
	}

	if len(u.Password) < 8 {
		return ErrorPasswordLength
	}

	return nil
}
