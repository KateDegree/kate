package validator

import (
	"back/pkg"
)

type GroupRequest struct {
	Name string `validate:"required"`
}

func CreateGroupValidator(name string) []string {
	request := GroupRequest{
		Name: name,
	}

	messages := map[string]map[string]string{
		"Name": {
			"required": "グループ名を入力してください。",
		},
	}

	return pkg.Validator(request, messages)
}
