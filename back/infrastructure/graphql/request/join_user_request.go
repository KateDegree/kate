package request

import (
	"back/infrastructure/graphql/request/internal"
	"github.com/graphql-go/graphql"
)

type joinUserInput struct {
	GroupID int `validate:"required"`
	UserID  int `validate:"required"`
}

type JoinUserRequest struct {
	Input    joinUserInput
	Messages []string
	IsValid  bool
}

func NewJoinUserRequest(rp graphql.ResolveParams) *JoinUserRequest {
	input := &joinUserInput{
		GroupID: rp.Args["groupId"].(int),
		UserID:  rp.Args["userId"].(int),
	}

	rules := map[string]map[string]string{
		"GroupID": {
			"required": "グループIDは必須です",
		},
		"UserID": {
			"required": "ユーザーIDは必須です",
		},
	}

	msgs, ok := internal.Validate(input, rules)
	return &JoinUserRequest{
		Input:    *input,
		Messages: msgs,
		IsValid:  ok,
	}
}
