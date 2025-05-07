package usecase

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/pkg"
)

type updateGroupUsecase struct {
	groupRepository repository.GroupRepository
}

func NewUpdateGroupUsecase(groupRepository repository.GroupRepository) *updateGroupUsecase {
	return &updateGroupUsecase{groupRepository: groupRepository}
}

func (u *updateGroupUsecase) Execute(ge *entity.GroupEntity, userID uint) (*entity.GroupEntity, *pkg.Error) {
	group, err := u.groupRepository.Update(ge, userID)
	if err != nil {
		return nil, &pkg.Error{
			Code:    400,
			Message: "グループの更新に失敗しました",
		}
	}

	return group, nil
}
