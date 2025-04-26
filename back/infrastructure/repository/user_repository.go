package repository

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/infrastructure/model"
	"fmt"
	"gorm.io/gorm"
)

type userRepository struct {
	orm *gorm.DB
}

func NewUserRepository(orm *gorm.DB) repository.UserRepository {
	return &userRepository{orm: orm}
}

func (r *userRepository) FindByAccountCode(accountCode string) (*entity.UserEntity, error) {
	var userModel model.UserModel
	result := r.orm.Where("account_code = ?", accountCode).First(&userModel)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find user: %w", result.Error)
	}

	userEntity := &entity.UserEntity{
		ID:          userModel.ID,
		AccountCode: userModel.AccountCode,
		Password:    userModel.Password,
	}

	return userEntity, nil
}

func (r *userRepository) Create(user *entity.UserEntity) (*entity.UserEntity, error) {
	userModel := &model.UserModel{
		Name:        user.Name,
		AccountCode: user.AccountCode,
		Password:    user.Password,
	}

	if err := r.orm.Create(userModel).Error; err != nil {
		return nil, err
	}

	return &entity.UserEntity{
		ID:          userModel.ID,
		Name:        userModel.Name,
		AccountCode: userModel.AccountCode,
		Password:    userModel.Password,
	}, nil
}
