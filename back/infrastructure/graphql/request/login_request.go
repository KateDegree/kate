package request

import (
	"back/pkg"
	"github.com/graphql-go/graphql"
)

type LoginInput struct {
	AccountCode string `validate:"required"`
	Password    string `validate:"required"`
}

type LoginRequest struct {
	Input    LoginInput
	Messages []string
	IsValid  bool
}

func NewLoginRequest(rp graphql.ResolveParams) *LoginRequest {
	input := &LoginInput{
		AccountCode: rp.Args["accountCode"].(string),
		Password:    rp.Args["password"].(string),
	}

	rules := map[string]map[string]string{
		"AccountCode": {
			"required": "アカウントコードは必須です",
		},
		"Password": {
			"required": "パスワードは必須です",
		},
	}

	msgs, ok := pkg.Validate(input, rules)
	return &LoginRequest{
		Input:    *input,
		Messages: msgs,
		IsValid:  ok,
	}
}
