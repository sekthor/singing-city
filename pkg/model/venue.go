package model

import (
	"time"

	"gorm.io/gorm"
)

type Venue struct {
	gorm.Model
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Address     string     `json:"address"`
	Slots       []Timeslot `json:"slots"`
	User        User       `json:"-"`
	UserID      uint       `json:"userID"`
}

type Timeslot struct {
	gorm.Model
	VenueID  uint      `json:"venueID"`
	Time     time.Time `json:"time"`
	Artist   Artist    `json:"artist"`
	ArtistID uint      `json:"artistID"`
}
