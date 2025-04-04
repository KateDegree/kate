package main

import (
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"net/http"
	"back/infrastructure"
	"back/infrastructure/graphqlobject"
)

func main() {
	e := echo.New()
	e.POST("/graphql", func(c echo.Context) error {
		var orm = infrastructure.Gorm()

		var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
			Query: graphql.NewObject(graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"users":  graphqlobject.UsersObject(orm),
					"spaces": graphqlobject.SpacesObject(orm),
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
