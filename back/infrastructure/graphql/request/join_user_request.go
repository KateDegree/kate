package request

import (
	"back/infrastructure/graphql/request/internal"
	"github.com/graphql-go/graphql"
)

type joinUserInput struct {
	GroupID     int    `validate:"required"`
	AccountCode string `validate:"required"`
}

type JoinUserRequest struct {
	Input    joinUserInput
	Messages []string
	IsValid  bool
}

func NewJoinUserRequest(rp graphql.ResolveParams) *JoinUserRequest {
	input := &joinUserInput{
		GroupID:     rp.Args["groupId"].(int),
		AccountCode: rp.Args["accountCode"].(string),
	}

	rules := map[string]map[string]string{
		"GroupID": {
			"required": "グループIDは必須です",
		},
		"AccountCode": {
			"required": "アカウントコードは必須です",
		},
	}

	msgs, ok := internal.Validate(input, rules)
	return &JoinUserRequest{
		Input:    *input,
		Messages: msgs,
		IsValid:  ok,
	}
}
