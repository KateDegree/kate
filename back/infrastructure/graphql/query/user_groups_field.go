package query

import (
	"back/domain/entity"
	"back/infrastructure/repository"
	"back/usecase"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserGroupsField(orm *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
			Name: "UserGroups",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"createdAt": &graphql.Field{
					Type: graphql.String,
				},
				"updatedAt": &graphql.Field{
					Type: graphql.String,
				},
			},
		})),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			authUser, ok := p.Context.Value("authUser").(*entity.UserEntity)
			if !ok || authUser == nil {
				return nil, echo.NewHTTPError(401, "認証エラー: ユーザーが見つかりません")
			}

			userGroupsUsecase := usecase.NewGetUserGroupsUsecase(
				repository.NewGroupRepository(orm),
			)

			groups, err := userGroupsUsecase.Execute(authUser.ID)
			if err != nil {
				return nil, echo.NewHTTPError(500, err.Message)
			}

			return groups, nil
		},
	}
}
