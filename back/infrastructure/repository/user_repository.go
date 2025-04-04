package repository

import (
	"gorm.io/gorm"
	"back/domain/repository"
)

type userRepository struct{
	orm *gorm.DB
}
func NewUserRepository(orm *gorm.DB) repository.UserRepository {
	return &userRepository{orm: orm}
}

func (r *userRepository) Login(email string, password string) (string, error) {
	// 処理
}

func (r *userRepository) SignUp(name string, email string, password string) error {
	// 処理
}
