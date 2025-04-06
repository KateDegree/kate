package entity

import "time"

type AccessTokenEntity struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Token     string `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
