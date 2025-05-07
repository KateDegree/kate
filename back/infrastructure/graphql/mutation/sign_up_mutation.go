package mutation

import (
	"back/domain/entity"
	"back/infrastructure/graphql/request"
	"back/infrastructure/repository"
	"back/usecase"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type SignUpResponse struct {
	AccessToken string   `json:"accessToken"`
	Success     bool     `json:"success"`
	Messages    []string `json:"messages"`
}

func SignUpMutation(orm *gorm.DB) *graphql.Field {
	signUpResponseType := graphql.NewObject(graphql.ObjectConfig{
		Name: "SignUp",
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
		Resolve: func(rp graphql.ResolveParams) (interface{}, error) {
			signUpRequest := request.NewSignUpRequest(rp)
			if !signUpRequest.IsValid {
				return SignUpResponse{
					Success:  false,
					Messages: signUpRequest.Messages,
				}, nil
			}

			name := signUpRequest.Input.Name
			accountCode := signUpRequest.Input.AccountCode
			password := signUpRequest.Input.Password

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
