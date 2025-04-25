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

	plainPassword := "Kate0418"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	orm := infrastructure.Gorm()

	user := model.UserModel{
		Name:     "nakao",
		Email:    "nakao@gmail.com",
		Password: string(hashedPassword),
	}

	space := model.SpaceModel{
		Name: "テストスペース",
		Type: "private",
	}

	if err := orm.Create(&space).Error; err != nil {
		log.Fatalf("Error creating space: %v", err)
	}

	user.Spaces = append(user.Spaces, space)

	if err := orm.Create(&user).Error; err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	log.Println("User and space created successfully")
}
