package repository

import (
	"back/domain/entity"
)

type GroupRepository interface {
	Create(ge *entity.GroupEntity, userID uint) (*entity.GroupEntity, error)
	FindByUserID(userID uint) ([]*entity.GroupEntity, error)
	Update(ge *entity.GroupEntity, userID uint) (*entity.GroupEntity, error)
	Delete(groupID uint, userID uint) (*entity.GroupEntity, error)
	JoinUser(groupID uint, accountCode string, userID uint) (*entity.GroupEntity, error)
}
