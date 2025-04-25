package service

import (
	"golang.org/x/crypto/bcrypt"
)

type authService struct{}

func NewAuthService() *authService {
	return &authService{}
}

func (s *authService) ValidatePassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
