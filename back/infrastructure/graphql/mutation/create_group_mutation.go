package mutation

import (
	"back/domain/entity"
	"back/infrastructure/graphql/request"
	"back/infrastructure/repository"
	"back/usecase"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type CreateGroupResponse struct {
	Success  bool     `json:"success"`
	Messages []string `json:"messages"`
}

func CreateGroupMutation(orm *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "CreateGroup",
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
		Resolve: func(rp graphql.ResolveParams) (interface{}, error) {
			createGroupRequest := request.NewCreateGroupRequest(rp)
			if !createGroupRequest.IsValid {
				return CreateGroupResponse{
					Success:  false,
					Messages: createGroupRequest.Messages,
				}, nil
			}

			authUser := rp.Context.Value("authUser").(*entity.UserEntity)
			name := createGroupRequest.Input.Name

			createGroupUsecase := usecase.NewCreateGroupUsecase(
				repository.NewGroupRepository(orm),
				repository.NewPointRepository(orm),
				repository.NewTransactionRepository(orm),
			)
			_, err := createGroupUsecase.Execute(entity.GroupEntity{
				Name: name,
			}, authUser.ID)
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
