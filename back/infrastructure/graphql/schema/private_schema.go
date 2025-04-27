package schema

import (
	"back/infrastructure/graphql/mutation"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

func PrivateSchema(orm *gorm.DB) (graphql.Schema, error) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"_dummy": &graphql.Field{Type: graphql.String},
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"createGroup": mutation.CreateGroupField(orm),
			},
		}),
	})
	if err != nil {
		return graphql.Schema{}, err
	}
	return schema, nil
}
