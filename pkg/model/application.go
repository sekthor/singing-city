package model

import "gorm.io/gorm"

type Application struct {
	gorm.Model

	// the artist applying for the timeslot
	ArtistID uint
	Artist   Artist

	TimeslotID uint
	Timeslot   Timeslot

	Confirmed bool
}
