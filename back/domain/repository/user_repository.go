package repository

import "back/domain/entity"

type UserRepository interface {
	FindByAccountCode(account_code string) (*entity.UserEntity, error)
}
