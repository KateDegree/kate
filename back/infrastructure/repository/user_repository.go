package repository

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/infrastructure/model"
	"gorm.io/gorm"
	"fmt"
)

type userRepository struct {
	orm *gorm.DB
}

func NewUserRepository(orm *gorm.DB) repository.UserRepository {
	return &userRepository{orm: orm}
}

func (r *userRepository) FindByEmail(email string) (*entity.UserEntity, error) {
	var userModel model.UserModel
	result := r.orm.Where("email = ?", email).First(&userModel)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find user: %w", result.Error)
	}

	userEntity := &entity.UserEntity{
		ID:    userModel.ID,
		Email: userModel.Email,
		Password: userModel.Password,
	}

	return userEntity, nil
}
