package model

import "gorm.io/gorm"

type Artist struct {
	gorm.Model
	Name        string       `json:"name"`
	User        User         `json:"-"`
	UserID      uint         `json:"userID"`
	Contact     string       `json:"contact"`
	Genere      string       `json:"genere"`
	Description string       `json:"description"`
	Socials     []SocialLink `json:"socials"`
}
