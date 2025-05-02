package usecase

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/pkg"
)

type createGroupUsecase struct {
	groupRepository repository.GroupRepository
	pointRepository repository.PointRepository
}

func NewCreateGroupUsecase(
	groupRepository repository.GroupRepository,
	pointRepository repository.PointRepository,
) *createGroupUsecase {
	return &createGroupUsecase{
		groupRepository: groupRepository,
		pointRepository: pointRepository,
	}
}

// TODO: トランザクションを実装するべき
func (u *createGroupUsecase) Execute(ge entity.GroupEntity, userID uint) (*entity.GroupEntity, *pkg.Error) {
	groupEntity, err := u.groupRepository.Create(&ge, userID)
	if err != nil {
		return nil, &pkg.Error{
			Code:    400,
			Message: "グループの登録に失敗しました。",
		}
	}

	pe := entity.PointEntity{
		UserID:  userID,
		GroupID: groupEntity.ID,
	}
	_, err = u.pointRepository.Create(&pe)
	if err != nil {
		return nil, &pkg.Error{
			Code:    400,
			Message: "ポイントの登録に失敗しました。",
		}
	}

	return groupEntity, nil
}
