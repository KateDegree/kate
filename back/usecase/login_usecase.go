package usecase

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"back/domain/repository"
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

func (u *loginUsecase) Login(email, password string) (string, error) {
	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return "", fmt.Errorf("failed to find user: %w", err)
	}
	if user == nil {
		return "", fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid password: %w", err)
	}

	accessTokenEntity, err := u.accessTokenRepository.Create(user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to create access token: %w", err)
	}

	return accessTokenEntity.Token, nil
}