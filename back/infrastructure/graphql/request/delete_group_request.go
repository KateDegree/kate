package request

import (
	"back/infrastructure/graphql/request/internal"
	"github.com/graphql-go/graphql"
)

type DeleteGroupInput struct {
	GroupID int
}

type DeleteGroupRequest struct {
	Input    DeleteGroupInput
	Messages []string
	IsValid  bool
}

func NewDeleteGroupRequest(rp graphql.ResolveParams) *DeleteGroupRequest {
	input := DeleteGroupInput{
		GroupID: rp.Args["groupId"].(int),
	}

	rules := map[string]map[string]string{
		"GroupID": {
			"required": "グループIDは必須です。",
		},
	}

	msgs, ok := internal.Validate(input, rules)
	return &DeleteGroupRequest{
		Input:    input,
		Messages: msgs,
		IsValid:  ok,
	}
}
