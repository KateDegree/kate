package main

import (
	"back/infrastructure"
	"back/infrastructure/model"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	// TODO: 環境変数の読み込み
	LoadEnv()

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
