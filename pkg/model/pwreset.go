package model

import "time"

type PasswordReset struct {
	Time   time.Time
	UserID uint
	Code   string `gorm:"primaryKey"`
}
