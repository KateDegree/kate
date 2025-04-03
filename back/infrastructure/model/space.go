package model

import (
	"gorm.io/gorm"
	"time"
)

type Space struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Name string `gorm:"type:varchar(255);not null" json:"name"`
	Type string `gorm:"type:varchar(10);not null" json:"type"`

	// リレーションは使う際にコメントアウトを解除する
	// Users       []User      `gorm:"many2many:space_users;"`
	// Directories []Directory `gorm:"foreignkey:SpaceID"`
	// Files       []File      `gorm:"foreignkey:SpaceID"`
}
