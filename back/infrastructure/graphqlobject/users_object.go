package graphqlobject

import (
	"back/infrastructure/model" //TODO: domain/entityから取得するようにする
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

func UsersObject(orm *gorm.DB) *graphql.Field {
	var userType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Users",
		Fields: graphql.Fields{
			"id":         &graphql.Field{Type: graphql.String},
			"name":       &graphql.Field{Type: graphql.String},
			"email":      &graphql.Field{Type: graphql.String},
			"created_at": &graphql.Field{Type: graphql.DateTime},
			"updated_at": &graphql.Field{Type: graphql.DateTime},
			"deleted_at": &graphql.Field{Type: graphql.DateTime},
			"spaces": &graphql.Field{
				Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
					Name: "Space",
					Fields: graphql.Fields{
						"id": &graphql.Field{Type: graphql.String},
						"name": &graphql.Field{Type: graphql.String},
						"type": &graphql.Field{Type: graphql.String},
					},
				})),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(model.UserModel); ok {
						// ここで Preload を使って関連データをロードする
						if err := orm.Preload("Spaces").Find(&user).Error; err != nil {
							return nil, err
						}
						return user.Spaces, nil
					}
					return nil, nil
				},
			},
		},
	})

	return &graphql.Field{
		Type: graphql.NewList(userType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var users []model.UserModel
			if err := orm.Preload("Spaces").Find(&users).Error; err != nil {
				return nil, err
			}
			return users, nil
		},
	}
}
