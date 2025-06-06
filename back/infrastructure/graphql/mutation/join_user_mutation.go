package mutation

import (
	"back/domain/entity"
	"back/infrastructure/graphql/request"
	"back/infrastructure/repository"
	"back/usecase"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type joinUserResponse struct {
	Success  bool     `json:"success"`
	Messages []string `json:"messages"`
}

func JoinUserMutation(orm *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "joinUser",
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
			"userId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(rp graphql.ResolveParams) (interface{}, error) {
			joinUserRequest := request.NewJoinUserRequest(rp)
			if !joinUserRequest.IsValid {
				return joinUserResponse{
					Success:  false,
					Messages: joinUserRequest.Messages,
				}, nil
			}

			authUser := rp.Context.Value("authUser").(*entity.UserEntity)
			groupID := uint(joinUserRequest.Input.GroupID)
			userID := uint(joinUserRequest.Input.UserID)

			joinUserUsecase := usecase.NewJoinUserUsecase(
				repository.NewGroupRepository(orm),
				repository.NewPointRepository(orm),
				repository.NewTransactionRepository(orm),
			)
			_, err := joinUserUsecase.Execute(groupID, userID, authUser.ID)
			if err != nil {
				return joinUserResponse{
					Success:  false,
					Messages: []string{err.Message},
				}, nil
			}

			return joinUserResponse{
				Success:  true,
				Messages: []string{"ユーザーを招待しました"},
			}, nil
		},
	}
}
