package validator

import (
	"back/pkg"
)

type Request struct {
	AccountCode string `validate:"required"`
	Password    string `validate:"required"`
}

func LoginValidator(accountCode, password string) []string {
	request := Request{
		AccountCode: accountCode,
		Password:    password,
	}

	messages := map[string]map[string]string{
		"AccountCode": {
			"required": "アカウントコードは必須です。",
		},
		"Password": {
			"required": "パスワードは必須です。",
		},
	}

	return pkg.Validator(request, messages)
}
