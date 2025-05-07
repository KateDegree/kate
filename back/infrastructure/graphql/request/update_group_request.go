package request

import (
	"back/pkg"
	"github.com/graphql-go/graphql"
)

type UpdateGroupInput struct {
	GroupID int    `validate:"required|gt=0"`
	Name    string `validate:"required|max=255"`
}

type UpdateGroupRequest struct {
	Input    UpdateGroupInput
	Messages []string
	IsValid  bool
}

func NewUpdateGroupRequest(rp graphql.ResolveParams) *UpdateGroupRequest {
	input := UpdateGroupInput{
		GroupID: rp.Args["groupId"].(int),
		Name:    rp.Args["name"].(string),
	}

	rules := map[string]map[string]string{
		"GroupID": {
			"required": "グループIDを入力してください。",
			"gt":       "グループIDは0より大きい必要があります",
		},
		"Name": {
			"required": "グループ名を入力してください。",
			"max":      "グループ名は255文字以内で入力してください。",
		},
	}

	msgs, ok := pkg.Validate(input, rules)
	return &UpdateGroupRequest{
		Input:    input,
		Messages: msgs,
		IsValid:  ok,
	}
}
