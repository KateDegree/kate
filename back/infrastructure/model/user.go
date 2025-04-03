package model

import (
	"gorm.io/gorm"
	"time"
)

// gorm.Modelを使用しないこと
// jsonタグを付けること
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Name     string `gorm:"size:255;not null" json:"name"`
	Email    string `gorm:"size:255;unique;not null" json:"email"`
	Password string `gorm:"size:255;not null" json:"password"`

	// リレーションは使う際にコメントアウトを解除する
	Spaces []Space `gorm:"many2many:space_users;" json:"spaces"`
	// AccessTokens []AccessToken `gorm:"foreignKey:UserID" json:"access_tokens"`
}
