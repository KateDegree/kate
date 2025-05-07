package mutation

import (
	"back/domain/entity"
	"back/infrastructure/graphql/request"
	"back/infrastructure/repository"
	"back/usecase"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type updateGroupResponse struct {
	Success  bool     `json:"success"`
	Messages []string `json:"messages"`
}

func UpdateGroupMutation(orm *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "UpdateGroup",
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
			"groupId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(rp graphql.ResolveParams) (interface{}, error) {
			updateGroupRequest := request.NewUpdateGroupRequest(rp)
			if !updateGroupRequest.IsValid {
				return updateGroupResponse{
					Success:  false,
					Messages: updateGroupRequest.Messages,
				}, nil
			}

			authUser := rp.Context.Value("authUser").(*entity.UserEntity)
			groupID := uint(updateGroupRequest.Input.GroupID)
			name := updateGroupRequest.Input.Name

			updateGroupUsecase := usecase.NewUpdateGroupUsecase(
				repository.NewGroupRepository(orm),
			)

			ge := &entity.GroupEntity{
				ID:   groupID,
				Name: name,
			}
			_, err := updateGroupUsecase.Execute(ge, authUser.ID)
			if err != nil {
				return updateGroupResponse{
					Success:  false,
					Messages: []string{err.Message},
				}, nil
			}

			return updateGroupResponse{
				Success:  true,
				Messages: []string{"グループを更新しました"},
			}, nil
		},
	}
}
