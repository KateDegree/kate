package repository

import "back/domain/entity"

type UserRepository interface {
	FindByEmail(email string) (*entity.UserEntity, error)
}
