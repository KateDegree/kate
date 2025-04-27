package main

import (
	"back/infrastructure"
	"back/infrastructure/graphql/schema"
	"back/pkg"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	pkg.LoadEnv()
	orm := infrastructure.Gorm()

	privateSchema, err := schema.PrivateSchema(orm)
	if err != nil {
		panic(err)
	}
	publicSchema, err := schema.PublicSchema(orm)
	if err != nil {
		panic(err)
	}

	// GraphQLのリクエストを処理する共通ハンドラー
	handleGraphQL := func(s graphql.Schema) echo.HandlerFunc {
		return func(c echo.Context) error {
			var params struct {
				Query string `json:"query"`
			}
			if err := c.Bind(&params); err != nil {
				return err
			}
			result := graphql.Do(graphql.Params{
				Schema:        s,
				RequestString: params.Query,
			})
			return c.JSON(http.StatusOK, result)
		}
	}

	e := echo.New()
	e.POST("/graphql/public", handleGraphQL(publicSchema))
	e.POST("/graphql/private", handleGraphQL(privateSchema))

	e.Logger.Fatal(e.Start(":8080"))
}
