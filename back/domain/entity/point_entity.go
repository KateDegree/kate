package entity

import (
	"time"
)

type PointEntity struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	UserID  uint `json:"user_id"`
	GroupID uint `json:"group_id"`
	Amount  int  `json:"amount"`
}
