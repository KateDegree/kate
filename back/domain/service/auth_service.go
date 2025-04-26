package service

import (
	"back/domain/repository"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type authService struct{}

func NewAuthService() *authService {
	return &authService{}
}

// ログイン時のパスワード検証
func (s *authService) ValidatePassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

// アカウントコードの規則検証
func (s *authService) IsValidAccountCode(code string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	return re.MatchString(code)
}

// アカウントコードの重複検証
func (s *authService) IsAccountCodeDuplicate(accountCode string, userRepository repository.UserRepository) bool {
	user, err := userRepository.FindByAccountCode(accountCode)
	if err != nil {
		return false
	}
	return user != nil
}
