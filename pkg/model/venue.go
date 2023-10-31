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
	ZipCode     int        `json:"zip"`
	City        string     `json:"city"`
	Slots       []Timeslot `json:"slots"`
	User        User       `json:"-"`
	UserID      uint       `json:"userID"`
	Contact     string     `json:"contact"`
}

type Timeslot struct {
	gorm.Model
	VenueID  uint      `json:"venueID"`
	Time     time.Time `json:"time"`
	Artist   Artist    `json:"artist"`
	ArtistID uint      `json:"artistID"`
	Pay      int       `json:"pay"`
	Private  bool      `json:"private"`
	Duration int       `json:"duration"`
}
