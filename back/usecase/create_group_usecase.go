package usecase

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/pkg"
)

type createGroupUsecase struct {
	groupRepository repository.GroupRepository
}

func NewCreateGroupUsecase(
	groupRepository repository.GroupRepository,
) *createGroupUsecase {
	return &createGroupUsecase{
		groupRepository: groupRepository,
	}
}

func (u *createGroupUsecase) Execute(ge entity.GroupEntity, userID uint) (*entity.GroupEntity, *pkg.Error) {
	groupEntity, err := u.groupRepository.Create(&ge, userID)
	if err != nil {
		return nil, &pkg.Error{
			Code:    400,
			Message: "グループの登録に失敗しました。",
		}
	}

	return groupEntity, nil
}
