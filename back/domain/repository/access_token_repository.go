package repository

import "back/domain/entity"

type AccessTokenRepository interface {
	Create(userId uint) (*entity.AccessTokenEntity, error)
}
