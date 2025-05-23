package repository

import (
	"back/domain/entity"
)

type GroupRepository interface {
	Create(ge *entity.GroupEntity, authID uint) (*entity.GroupEntity, error)
	FindByUserID(userID uint) ([]*entity.GroupEntity, error)
	Update(ge *entity.GroupEntity, authID uint) (*entity.GroupEntity, error)
	Delete(groupID uint, authID uint) (*entity.GroupEntity, error)
	LinkUser(groupID, userID, authID uint) (*entity.GroupEntity, error)
}
