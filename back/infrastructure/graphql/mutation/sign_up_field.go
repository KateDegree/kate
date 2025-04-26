package mutation

import (
	"back/domain/entity"
	"back/infrastructure/repository"
	"back/interface/validator"
	"back/usecase"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type SignUpResponse struct {
	AccessToken string   `json:"accessToken"`
	Success     bool     `json:"success"`
	Messages    []string `json:"messages"`
}

func SignUpField(orm *gorm.DB) *graphql.Field {
	signUpResponseType := graphql.NewObject(graphql.ObjectConfig{
		Name: "SignUpResponse",
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
		Type: signUpResponseType,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"accountCode": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			name := p.Args["name"].(string)
			accountCode := p.Args["accountCode"].(string)
			password := p.Args["password"].(string)

			errorMessages := validator.SignUpValidator(name, accountCode, password)
			if len(errorMessages) > 0 {
				return SignUpResponse{
					Success:  false,
					Messages: errorMessages,
				}, nil
			}

			signUpUsecase := usecase.NewSignUpUsecase(
				repository.NewUserRepository(orm),
				repository.NewAccessTokenRepository(orm),
			)
			signUpUsecaseResponse, err := signUpUsecase.Execute(entity.UserEntity{
				Name:        name,
				AccountCode: accountCode,
				Password:    password,
			})

			if err != nil {
				return SignUpResponse{
					AccessToken: "",
					Success:     false,
					Messages:    []string{err.Message},
				}, nil
			}

			return SignUpResponse{
				AccessToken: signUpUsecaseResponse.AccessToken,
				Success:     true,
				Messages:    []string{"ログイン成功"},
			}, nil
		},
	}
}
