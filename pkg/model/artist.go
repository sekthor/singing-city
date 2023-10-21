package model

import "gorm.io/gorm"

type Artist struct {
	gorm.Model
	Name   string
	User   User `json:"-"`
	UserID uint
}
