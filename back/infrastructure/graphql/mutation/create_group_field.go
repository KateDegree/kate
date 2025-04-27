package mutation

import (
	"back/domain/entity"
	"back/infrastructure/repository"
	"back/interface/validator"
	"back/usecase"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type CreateGroupResponse struct {
	Success  bool     `json:"success"`
	Messages []string `json:"messages"`
}

func CreateGroupField(orm *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "CreateGroupResponse",
			Fields: graphql.Fields{
				"success": &graphql.Field{
					Type: graphql.Boolean,
				},
				"messages": &graphql.Field{
					Type: graphql.NewList(graphql.String),
				},
			},
		}),
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			name := p.Args["name"].(string)
			errorMessages := validator.CreateGroupValidator(name)
			if len(errorMessages) > 0 {
				return CreateGroupResponse{
					Success:  false,
					Messages: errorMessages,
				}, nil
			}

			createGroupUsecase := usecase.NewCreateGroupUsecase(
				repository.NewGroupRepository(orm),
			)
			_, err := createGroupUsecase.Execute(entity.GroupEntity{
				Name: name,
			}, 1) // TODO: 認証トークンからログインユーザーのIDを使用したい
			if err != nil {
				return CreateGroupResponse{
					Success:  false,
					Messages: []string{err.Message},
				}, nil
			}

			return CreateGroupResponse{
				Success:  true,
				Messages: []string{"グループを登録しました。"},
			}, nil
		},
	}
}
