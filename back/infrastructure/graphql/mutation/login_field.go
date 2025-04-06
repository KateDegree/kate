package mutation

import (
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
	"back/infrastructure/repository"
	"back/usecase"
)

// LoginResponse はログイン処理の結果を表す構造体です。
type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	Success     bool   `json:"success"`
	Messages    string `json:"messages"`
}

// TODO: interfaceを使って入出力を管理する
func LoginField(orm *gorm.DB) *graphql.Field {
	// LoginResponse 用の GraphQL オブジェクトを定義
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
				Type: graphql.String,
			},
		},
	})

	return &graphql.Field{
		Type: loginResponseType,  // 戻り値の型を LoginResponse に設定
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.String,  // 引数としてメールアドレスを受け取る
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.String,  // 引数としてパスワードを受け取る
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 引数を取得
			email := p.Args["email"].(string)
			password := p.Args["password"].(string)

			// ユースケースのインスタンス化
			loginUsecase := usecase.NewLoginUsecase(
				repository.NewUserRepository(orm),
				repository.NewAccessTokenRepository(orm),
			)

			// ログイン処理を実行
			loginUsecaseResponse, err := loginUsecase.Execute(email, password)
			if err != nil {
				// エラーがあれば、エラーメッセージと失敗フラグを返す
				return LoginResponse{
					AccessToken: "",
					Success:     false,
					Messages:    err.Error(),
				}, nil
			}

			// ログイン成功時のレスポンスを返す
			return LoginResponse{
				AccessToken: loginUsecaseResponse.AccessToken,
				Success:     true,
				Messages:    "ログイン成功",
			}, nil
		},
	}
}
