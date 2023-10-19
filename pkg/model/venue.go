package model

import (
	"gorm.io/gorm"
	"time"
)

type Venue struct {
	gorm.Model
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Slots       []Timeslot `json:"slots"`
}

type Timeslot struct {
	gorm.Model
	VenueID uint
	Time    time.Time `json:"time"`
}
