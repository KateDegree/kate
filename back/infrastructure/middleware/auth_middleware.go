package middleware

import (
	"back/domain/service"
	"back/infrastructure/repository"
	"bytes"
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"io"
	"net/http"
	"strings"
)

func AuthMiddleware(orm *gorm.DB, next echo.HandlerFunc) echo.HandlerFunc {
	var SKIP_AUTH_MUTATIONS = []string{"login", "signUp"}

	return func(c echo.Context) error {
		var bodyBytes []byte
		if c.Request().Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request().Body)
			c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		var requestData map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &requestData); err == nil {
			if query, ok := requestData["query"].(string); ok && strings.Contains(query, "mutation") {
				for _, mutation := range SKIP_AUTH_MUTATIONS {
					if strings.Contains(query, mutation) {
						return next(c)
					}
				}
			}
		}

		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			return echo.NewHTTPError(http.StatusUnauthorized, "認証が必要です")
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		authService := service.NewAuthService()
		if !authService.ValidateToken(tokenString, repository.NewAccessTokenRepository(orm)) {
			return echo.NewHTTPError(http.StatusUnauthorized, "認証が必要です")
		}

		authUser, err := repository.NewUserRepository(orm).FindByToken(tokenString)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "ユーザーの取得に失敗しました")
		}
		if authUser == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "ユーザーが見つかりません")
		}

		ctx := context.WithValue(c.Request().Context(), "authUser", authUser)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
