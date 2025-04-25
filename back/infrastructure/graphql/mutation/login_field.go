package mutation

import (
	"back/domain/entity"
	"back/infrastructure/repository"
	"back/interface/validator"
	"back/usecase"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type LoginResponse struct {
	AccessToken string   `json:"accessToken"`
	Success     bool     `json:"success"`
	Messages    []string `json:"messages"`
}

func LoginField(orm *gorm.DB) *graphql.Field {
	loginResponseType := graphql.NewObject(graphql.ObjectConfig{
		Name: "LoginResponse",
		Fields: graphql.Fields{
			"accessToken": &graphql.Field{
				Type: graphql.String,
			},
			"success": &graphql.Field{
				Type: graphql.Boolean,
			},
			"messages": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	})

	return &graphql.Field{
		Type: loginResponseType,
		Args: graphql.FieldConfigArgument{
			"accountCode": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			accountCode := p.Args["accountCode"].(string)
			password := p.Args["password"].(string)

			errorMessages := validator.LoginValidator(accountCode, password)
			if len(errorMessages) > 0 {
				return LoginResponse{
					AccessToken: "",
					Success:     false,
					Messages:    errorMessages,
				}, nil
			}

			loginUsecase := usecase.NewLoginUsecase(
				repository.NewUserRepository(orm),
				repository.NewAccessTokenRepository(orm),
			)
			loginUsecaseResponse, err := loginUsecase.Execute(
				entity.UserEntity{
					AccountCode: accountCode,
					Password:    password,
				})
			if err != nil {
				return LoginResponse{
					AccessToken: "",
					Success:     false,
					Messages:    []string{err.Message},
				}, nil
			}

			return LoginResponse{
				AccessToken: loginUsecaseResponse.AccessToken,
				Success:     true,
				Messages:    []string{"ログイン成功"},
			}, nil
		},
	}
}
