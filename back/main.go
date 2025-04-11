package main

import (
	"back/infrastructure"
	"back/infrastructure/graphql/mutation"
	"back/infrastructure/graphql/query"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.POST("/graphql", func(c echo.Context) error {
		var orm = infrastructure.Gorm()

		var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
			Query: graphql.NewObject(graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"users":  query.UsersField(orm),
					"spaces": query.SpacesField(orm),
				},
			}),
			Mutation: graphql.NewObject(graphql.ObjectConfig{
				Name: "Mutation",
				Fields: graphql.Fields{
					"login": mutation.LoginField(orm),
				},
			}),
		})

		var params struct {
			Query string `json:"query"`
		}
		if err := c.Bind(&params); err != nil {
			return err
		}

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: params.Query,
		})
		return c.JSON(http.StatusOK, result)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
