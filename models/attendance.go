package models

import (
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model `json:"-"`
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	LocationID int       `json:"location_id"`
	CheckInOut string    `json:"check_in_out"`
	User       User      `gorm:"foreignkey:UserID"`
	Location   Location  `gorm:"foreignkey:LocationID"`
	CreatedAt  time.Time `json:"created_at"`
}
