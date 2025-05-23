package usecase

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/usecase/internal"
)

type getUserGroupsUsecase struct {
	groupRepository repository.GroupRepository
}

func NewGetUserGroupsUsecase(
	groupRepository repository.GroupRepository,
) *getUserGroupsUsecase {
	return &getUserGroupsUsecase{
		groupRepository: groupRepository,
	}
}

func (u *getUserGroupsUsecase) Execute(userID uint) ([]entity.GroupEntity, *internal.UsecaseError) {
	groups, err := u.groupRepository.FindByUserID(userID)
	if err != nil {
		return nil, &internal.UsecaseError{
			Code:    500,
			Message: "グループの取得に失敗しました",
		}
	}

	var result []entity.GroupEntity
	for _, group := range groups {
		result = append(result, *group)
	}
	return result, nil
}
