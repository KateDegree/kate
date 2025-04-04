package model

import (
	"gorm.io/gorm"
	"time"
)

type AccessTokenModel struct {
	gorm.Model
	UserID    uint      `gorm:"not null"`
	Token     string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
}
