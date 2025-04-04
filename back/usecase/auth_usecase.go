package usecase

type authUsecase struct{}
func NewAuthUsecase() *auth {
	return &authUsecase{}
}

func (u *authUsecase) Login(email, password string) {
	// 処理
}

func (u *authUsecase) SignUp(username, email, password string) (string, error) {
	// 処理
}
