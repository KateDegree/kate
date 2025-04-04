package graphqlobject

import (
	"back/infrastructure/model"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

func SpacesObject(orm *gorm.DB) *graphql.Field {
	var spacesType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Spaces",
		Fields: graphql.Fields{
			"id":         &graphql.Field{Type: graphql.String},
			"name":       &graphql.Field{Type: graphql.String},
			"type":       &graphql.Field{Type: graphql.String},
			"created_at": &graphql.Field{Type: graphql.DateTime},
			"updated_at": &graphql.Field{Type: graphql.DateTime},
			"deleted_at": &graphql.Field{Type: graphql.DateTime},
		},
	})

	return &graphql.Field{
		Type: graphql.NewList(spacesType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var spaces []model.SpaceModel
			if err := orm.Find(&spaces).Error; err != nil {
				return nil, err
			}
			return spaces, nil
		},
	}
}
