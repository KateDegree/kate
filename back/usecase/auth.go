package usecase

type auth struct{}
func NewAuth() *auth {
	return &auth{}
}

func (a *auth) Login(email, password string) {
	// 処理
}

func (a *auth) SignUp(username, email, password string) (string, error) {
	// 処理
}
