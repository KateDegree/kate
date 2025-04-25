package usecase

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/domain/service"
	"back/pkg"
)

type loginUsecase struct {
	userRepository        repository.UserRepository
	accessTokenRepository repository.AccessTokenRepository
}

func NewLoginUsecase(userRepository repository.UserRepository, accessTokenRepository repository.AccessTokenRepository) *loginUsecase {
	return &loginUsecase{
		userRepository:        userRepository,
		accessTokenRepository: accessTokenRepository,
	}
}

type loginUsecaseResponse struct {
	AccessToken string `json:"access_token"`
}

func (u *loginUsecase) Execute(ue entity.UserEntity) (*loginUsecaseResponse, *pkg.Error) {
	userEntity, err := u.userRepository.FindByEmail(ue.Email)
	if err != nil {
		return nil, &pkg.Error{
			Message: "ログイン処理中に問題が発生しました。時間をおいて再試行してください。",
			Code:    500,
			Cause:   err,
		}
	}
	if userEntity == nil {
		return nil, &pkg.Error{
			Message: "メールアドレスまたはパスワードが正しくありません。",
			Code:    400,
			Cause:   err,
		}
	}

	err = service.NewAuthService().ValidatePassword(userEntity.Password, ue.Password)
	if err != nil {
		return nil, &pkg.Error{
			Message: "メールアドレスまたはパスワードが正しくありません。",
			Code:    400,
			Cause:   err,
		}
	}

	accessTokenEntity, err := u.accessTokenRepository.Create(userEntity.ID)
	if err != nil {
		return nil, &pkg.Error{
			Message: "ログイン処理中に問題が発生しました。時間をおいて再試行してください。",
			Code:    500,
			Cause:   err,
		}
	}

	return &loginUsecaseResponse{
		AccessToken: accessTokenEntity.Token,
	}, nil
}
