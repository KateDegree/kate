package repository

import "back/domain/entity"

type AccessTokenRepository interface {
	Create(userId string) (*entity.AccessTokenEntity, error)
}
