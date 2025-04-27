package model

import (
	"gorm.io/gorm"
	"time"
)

func (GroupModel) TableName() string {
	return "groups"
}

type GroupModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Name string `gorm:"size:255;not null" json:"name"`
}
