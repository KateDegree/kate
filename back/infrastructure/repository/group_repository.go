package repository

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/infrastructure/model"
	"gorm.io/gorm"
)

type groupRepository struct {
	orm *gorm.DB
}

func NewGroupRepository(orm *gorm.DB) repository.GroupRepository {
	return &groupRepository{orm: orm}
}

// TODO: userIDは引数に必要か？ -> geに含める
func (r *groupRepository) Create(ge *entity.GroupEntity, userID uint) (*entity.GroupEntity, error) {
	var userModel model.UserModel

	if err := r.orm.First(&userModel, userID).Error; err != nil {
		return nil, err
	}

	groupModel := model.GroupModel{
		Name: ge.Name,
	}

	if err := r.orm.Create(&groupModel).Error; err != nil {
		return nil, err
	}

	if err := r.orm.Model(&userModel).Association("Groups").Append(&groupModel); err != nil {
		return nil, err
	}

	return &entity.GroupEntity{
		ID:   groupModel.ID,
		Name: groupModel.Name,
	}, nil
}

func (r *groupRepository) FindByUserID(userID uint) ([]*entity.GroupEntity, error) {
	var user model.UserModel

	if err := r.orm.Preload("Groups").First(&user, userID).Error; err != nil {
		return nil, err
	}

	groupEntities := make([]*entity.GroupEntity, 0, len(user.Groups))
	for _, gm := range user.Groups {
		groupEntities = append(groupEntities, &entity.GroupEntity{
			ID:        gm.ID,
			Name:      gm.Name,
			CreatedAt: gm.CreatedAt,
			UpdatedAt: gm.UpdatedAt,
		})
	}

	return groupEntities, nil
}

func (r *groupRepository) Update(ge *entity.GroupEntity, userID uint) (*entity.GroupEntity, error) {
	var user model.UserModel
	if err := r.orm.Preload("Groups").First(&user, userID).Error; err != nil {
		return nil, err
	}

	var groupModel *model.GroupModel
	for _, group := range user.Groups {
		if group.ID == ge.ID {
			groupModel = &group
			break
		}
	}

	if groupModel == nil {
		return nil, gorm.ErrRecordNotFound
	}

	groupModel.Name = ge.Name
	if err := r.orm.Save(groupModel).Error; err != nil {
		return nil, err
	}

	return &entity.GroupEntity{
		ID:        groupModel.ID,
		Name:      groupModel.Name,
		CreatedAt: groupModel.CreatedAt,
		UpdatedAt: groupModel.UpdatedAt,
	}, nil
}
