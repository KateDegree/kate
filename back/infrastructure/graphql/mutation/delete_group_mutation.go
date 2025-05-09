package mutation

import (
	"back/domain/entity"
	"back/infrastructure/graphql/request"
	"back/infrastructure/repository"
	"back/usecase"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type deleteGroupResponse struct {
	Success  bool     `json:"success"`
	Messages []string `json:"messages"`
}

func DeleteGroupMutation(orm *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "deleteGroup",
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
		},
		Resolve: func(rp graphql.ResolveParams) (interface{}, error) {
			deleteGroupRequest := request.NewDeleteGroupRequest(rp)
			if !deleteGroupRequest.IsValid {
				return deleteGroupResponse{
					Success:  false,
					Messages: deleteGroupRequest.Messages,
				}, nil
			}

			authUser := rp.Context.Value("authUser").(*entity.UserEntity)
			groupID := uint(deleteGroupRequest.Input.GroupID)

			deleteGroupUsecase := usecase.NewDeleteGroupUsecase(
				repository.NewGroupRepository(orm),
			)
			_, err := deleteGroupUsecase.Execute(groupID, authUser.ID)
			if err != nil {
				return deleteGroupResponse{
					Success:  false,
					Messages: []string{err.Message},
				}, nil
			}

			return deleteGroupResponse{
				Success:  true,
				Messages: []string{"グループを削除しました。"},
			}, nil
		},
	}
}
