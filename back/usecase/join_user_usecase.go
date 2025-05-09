package usecase

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/usecase/internal"
)

type joinUserUsecase struct {
	groupRepository repository.GroupRepository
}

func NewJoinUserUsecase(groupRepository repository.GroupRepository) *joinUserUsecase {
	return &joinUserUsecase{groupRepository: groupRepository}
}

func (u *joinUserUsecase) Execute(groupID uint, accountCode string, userID uint) (*entity.GroupEntity, *internal.UsecaseError) {
	group, err := u.groupRepository.JoinUser(groupID, accountCode, userID)
	if err != nil {
		return nil, &internal.UsecaseError{
			Code:    400,
			Message: "ユーザーの参加に失敗しました",
		}
	}

	return group, nil
}
