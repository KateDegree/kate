package model

import (
	"gorm.io/gorm"
	"time"
)

func (PointModel) TableName() string {
	return "points"
}

type PointModel struct {
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	UserID  uint `gorm:"primaryKey;not null" json:"user_id"`
	GroupID uint `gorm:"primaryKey;not null" json:"group_id`
	Amount  int  `gorm:"not null" json:"amount"`

	User  UserModel  `gorm:"foreignKey:UserID" json:"user"`
	Group GroupModel `gorm:"foreignKey:GroupID" json:"group"`
}
