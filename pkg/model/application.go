package model

import "gorm.io/gorm"

type Application struct {
	gorm.Model

	// the artist applying for the timeslot
	ArtistID uint   `json:"artistID"`
	Artist   Artist `json:"artist"`

	TimeslotID uint     `json:"timeslotID"`
	Timeslot   Timeslot `json:"timeslot"`

	Confirmed bool `json:"confirmed"`
}
