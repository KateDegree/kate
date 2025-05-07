package request

import (
	"back/pkg"
	"github.com/graphql-go/graphql"
)

type SignUpInput struct {
	Name        string `validate:"required|max=255"`
	AccountCode string `validate:"required|max=255"`
	Password    string `validate:"required|max=255"`
}

type SignUpRequest struct {
	Input    SignUpInput
	Messages []string
	IsValid  bool
}

func NewSignUpRequest(rp graphql.ResolveParams) *SignUpRequest {
	input := SignUpInput{
		Name:        rp.Args["name"].(string),
		AccountCode: rp.Args["accountCode"].(string),
		Password:    rp.Args["password"].(string),
	}

	rules := map[string]map[string]string{
		"Name": {
			"required": "名前を入力してください",
			"max":      "名前は255文字以内で入力してください",
		},
		"AccountCode": {
			"required": "アカウントコードを入力してください",
			"max":      "アカウントコードは255文字以内で入力してください",
		},
		"Password": {
			"required": "パスワードを入力してください",
			"max":      "パスワードは255文字以内で入力してください",
		},
	}

	msgs, ok := pkg.Validate(input, rules)
	return &SignUpRequest{
		Input:    input,
		Messages: msgs,
		IsValid:  ok,
	}
}
