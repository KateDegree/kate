package validator

import (
	"back/pkg"
)

type Request struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func LoginValidator(email, password string) []string {
	request := Request{
		Email:    email,
		Password: password,
	}

	messages := map[string]map[string]string{
		"Email": {
			"required": "メールアドレスは必須です。",
			"email":    "有効なメールアドレスを入力してください。",
		},
		"Password": {
			"required": "パスワードは必須です。",
		},
	}

	return pkg.Validator(request, messages)
}
