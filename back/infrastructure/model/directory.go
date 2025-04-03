package model

import (
	"gorm.io/gorm"
)

type Directory struct {
	gorm.Model
	SpaceID  uint   `gorm:"not null"`
	ParentID *uint  `gorm:"index"`
	Name     string `gorm:"type:varchar(255);not null"`

	// リレーションは使う際にコメントアウトを解除する
	// Space      Space       `gorm:"foreignKey:SpaceID"`
	// Parent     Directory   `gorm:"foreignKey:DirectoryID"`
	// Directories []Directory `gorm:"foreignkey:DirectoryID"`
	// Files      []File      `gorm:"foreignkey:FileID"`
}
