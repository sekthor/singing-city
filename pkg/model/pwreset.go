package model

import "time"

type PasswordReset struct {
	Time    time.Time
	UserID  uint
	Request string `gorm:"primaryKey"`
}
