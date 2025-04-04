package repository

import (
	"back/infrastructure/model"
	"back/domain/entity"
	"back/domain/repository"

	"gorm.io/gorm"
	"time"
)

type accessTokenRepository struct {
	orm *gorm.DB
}

func NewAccessTokenRepository(orm *gorm.DB) repository.AccessTokenRepository {
	return &accessTokenRepository{orm: orm}
}

func (r *accessTokenRepository) Create(userId string) (*entity.AccessTokenEntity, error) {
	accessToken := &model.AccessTokenModel{
		UserID:    *userId,
		Token:     uuid.New().String(), // TODO: 適切なトークン生成方法を選択
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	if err := r.orm.Create(accessToken).Error; err != nil {
		return nil, err
	}

	return &entity.AccessTokenEntity{
		ID:        accessToken.ID,
		UserID:    accessToken.UserID,
		Token:     accessToken.Token,
		ExpiresAt: accessToken.ExpiresAt,
	}, nil
}
