package request

import (
	"back/pkg"
	"github.com/graphql-go/graphql"
)

type CreateGroupInput struct {
	Name string `validate:"required|max=255"`
}

type CreateGroupRequest struct {
	Input    CreateGroupInput
	Messages []string
	IsValid  bool
}

func NewCreateGroupRequest(rp graphql.ResolveParams) *CreateGroupRequest {
	input := CreateGroupInput{
		Name: rp.Args["name"].(string),
	}

	rules := map[string]map[string]string{
		"Name": {
			"required": "グループ名を入力してください。",
			"max":      "グループ名は255文字以内で入力してください。",
		},
	}

	msgs, ok := pkg.Validate(input, rules)
	return &CreateGroupRequest{
		Input:    input,
		Messages: msgs,
		IsValid:  ok,
	}
}
