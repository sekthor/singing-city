package model

import (
	"time"
	"gorm.io/gorm"
)

type Venue struct {
    gorm.Model
    Name string
    Slots []Timeslot
}

type Timeslot struct {
    gorm.Model
    time.Time
}
