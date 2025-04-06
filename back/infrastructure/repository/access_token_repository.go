package repository

import (
	"back/infrastructure/model"
	"back/domain/entity"
	"back/domain/repository"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"time"
	"os"
)

type accessTokenRepository struct {
	orm *gorm.DB
}

func NewAccessTokenRepository(orm *gorm.DB) repository.AccessTokenRepository {
	return &accessTokenRepository{orm: orm}
}

func (r *accessTokenRepository) Create(userId uint) (*entity.AccessTokenEntity, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	accessToken := &model.AccessTokenModel{
		UserID:    userId,
		Token:     signedToken,
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
