package pkg

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validator(request any, messages map[string]map[string]string) []string {
	err := validate.Struct(request)
	if err == nil {
		return nil
	}

	ves, ok := err.(validator.ValidationErrors)
	if !ok {
		return []string{"エラーが発生しました。"}
	}

	var msgs []string
	for _, fe := range ves {
		if tags, ok := messages[fe.Field()]; ok {
			if msg, ok := tags[fe.Tag()]; ok {
				msgs = append(msgs, msg)
			}
		}
	}

	return msgs
}
