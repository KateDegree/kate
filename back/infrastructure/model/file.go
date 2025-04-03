package model

import (
	"gorm.io/gorm"
)

type FileType string

const (
	FileTypePage  FileType = "page"
	FileTypeTable FileType = "table"
	FileTypeTask  FileType = "task"
	FileTypeER    FileType = "er"
)

type File struct {
	gorm.Model
	SpaceID  uint     `gorm:"not null"`
	ParentID *uint    `gorm:"index"`
	Name     string   `gorm:"type:varchar(255);not null"`
	Type     FileType `gorm:"type:varchar(20);not null"`

	// リレーションは使う際にコメントアウトを解除する
	// Space     Space     `gorm:"foreignKey:SpaceID"`
	// Parent    Directory `gorm:"foreignKey:DirectoryID"`
}
