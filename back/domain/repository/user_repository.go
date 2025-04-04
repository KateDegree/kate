package repository

type UserRepository interface {
	Login(email string, password string) (string, error)
	SignUp(name string, email string, password string) error
}
