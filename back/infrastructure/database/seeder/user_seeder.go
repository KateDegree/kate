package main

import (
	"back/infrastructure"
	"back/infrastructure/model"
	"back/pkg"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	// .envファイルから環境変数を読み込む
	pkg.LoadEnv()

	plainPassword := "keito0418"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	orm := infrastructure.Gorm()

	user := model.UserModel{
		Name:        "中尾 渓斗",
		AccountCode: "kate.degree",
		Password:    string(hashedPassword),
	}

	if err := orm.Create(&user).Error; err != nil {
		log.Fatalf("Error creating user: %v", err)
	}
}
