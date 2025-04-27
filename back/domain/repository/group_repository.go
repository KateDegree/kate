package repository

import (
	"back/domain/entity"
)

type GroupRepository interface {
	Create(ge *entity.GroupEntity, userID uint) (*entity.GroupEntity, error)
}
