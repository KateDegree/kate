package validator

import (
	"back/pkg"
)

type SignUpRequest struct {
	Name        string `validate:"required"`
	AccountCode string `validate:"required"`
	Password    string `validate:"required"`
}

func SignUpValidator(name, accountCode, password string) []string {
	request := SignUpRequest{
		Name:        name,
		AccountCode: accountCode,
		Password:    password,
	}

	messages := map[string]map[string]string{
		"Name": {
			"required": "名前は必須です。",
		},
		"AccountCode": {
			"required": "アカウントコードは必須です。",
		},
		"Password": {
			"required": "パスワードは必須です。",
		},
	}

	// バリデーション実行
	return pkg.Validator(request, messages)
}
