package main

import (
	"back/infrastructure"
	"back/infrastructure/graphql/mutation"
	// "back/infrastructure/graphql/query"
	"back/pkg"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	// .envファイルから環境変数を読み込む
	pkg.LoadEnv()

	// GORMによるDB接続インスタンスの生成
	var orm = infrastructure.Gorm()

	// GraphQLスキーマの構築
	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Query",
			Fields: graphql.Fields{
				// "users":  query.UsersField(orm),
				// "spaces": query.SpacesField(orm),
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

	e := echo.New()
	e.POST("/graphql", func(c echo.Context) error {
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
