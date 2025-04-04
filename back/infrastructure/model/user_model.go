package model

import (
	"gorm.io/gorm"
	"time"
)

func (UserModel) TableName() string {
	return "users" // テーブル名を指定
}

// gorm.Modelを使用しないこと
// jsonタグを付けること
type UserModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Name     string `gorm:"size:255;not null" json:"name"`
	Email    string `gorm:"size:255;unique;not null" json:"email"`
	Password string `gorm:"size:255;not null" json:"password"`

	// リレーションは使う際にコメントアウトを解除する
	Spaces []SpaceModel `gorm:"many2many:space_users;foreignKey:ID;joinForeignKey:user_id;joinReferences:space_id" json:"spaces"`
	// AccessTokens []AccessToken `gorm:"foreignKey:UserID" json:"access_tokens"`
}
