package main

import (
	"back/infrastructure"
	"back/infrastructure/graphql/mutation"
	"back/infrastructure/graphql/query"
	"back/infrastructure/middleware"
	"back/pkg"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	pkg.LoadEnv()
	orm := infrastructure.Gorm()

	// スキーマ定義
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"userGroups": query.UserGroupsField(orm),
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"login":       mutation.LoginField(orm),
				"signUp":      mutation.SignUpField(orm),
				"createGroup": mutation.CreateGroupField(orm),
			},
		}),
	})
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return middleware.AuthMiddleware(orm, next)
	})

	e.POST(
		"/graphql",
		func(c echo.Context) error {
			var params struct {
				Query string `json:"query"`
			}
			if err := c.Bind(&params); err != nil {
				return err
			}
			result := graphql.Do(graphql.Params{
				Schema:        schema,
				RequestString: params.Query,
				Context:       c.Request().Context(),
			})
			return c.JSON(http.StatusOK, result)
		},
	)

	e.Logger.Fatal(e.Start(":8080"))
}
