package usecase

import (
	"back/domain/repository"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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
	AccessToken string `json:"access_token,omitempty"`
	Message     string `json:"message"`
}

func (u *loginUsecase) Execute(email, password string) (loginUsecaseResponse, error) {
	userEntity, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return loginUsecaseResponse{
			Message: "ログイン処理中に問題が発生しました。時間をおいて再試行してください。",
		}, fmt.Errorf("failed to find user: %w", err)
	}
	if userEntity == nil {
		return loginUsecaseResponse{
			Message: "メールアドレスまたはパスワードが正しくありません。",
		}, fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userEntity.Password), []byte(password))
	if err != nil {
		return loginUsecaseResponse{
			Message: "メールアドレスまたはパスワードが正しくありません。",
		}, fmt.Errorf("invalid password: %w", err)
	}

	accessTokenEntity, err := u.accessTokenRepository.Create(userEntity.ID)
	if err != nil {
		return loginUsecaseResponse{
			Message: "ログイン処理中に問題が発生しました。時間をおいて再試行してください。",
		}, fmt.Errorf("failed to create access token: %w", err)
	}

	return loginUsecaseResponse{
		AccessToken: accessTokenEntity.Token,
		Message:     "ログインに成功しました。",
	}, nil
}
