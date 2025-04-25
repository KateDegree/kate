package model

import (
	"gorm.io/gorm"
	"time"
)

func (UserModel) TableName() string {
	return "users"
}

type UserModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Name        string `gorm:"size:255;not null" json:"name"`
	AccountCode string `gorm:"size:255;unique;not null" json:"account_code"`
	Password    string `gorm:"size:255;not null" json:"password"`

	// Group       []GroupModel       `gorm:"many2many:group_users;foreignKey:ID;joinForeignKey:user_id;joinReferences:group_id" json:"groups"`
	AccessTokens []AccessTokenModel `gorm:"foreignKey:UserID" json:"access_tokens"`
}
