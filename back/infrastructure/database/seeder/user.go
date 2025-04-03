package main

import (
	"back/infrastructure"
	"back/infrastructure/model"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	// プレーンパスワードをハッシュ化
	plainPassword := "Kate0418"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	// データベース接続
	db := infrastructure.Gorm()

	// 新しいユーザーの作成
	user := model.User{
		Name:     "nakao",
		Email:    "nakao@gmail.com",
		Password: string(hashedPassword),
	}

	// 新しいスペースの作成
	space := model.Space{
		Name: "テストスペース",
		Type: "private",
	}

	// データベースにスペースを保存
	if err := db.Create(&space).Error; err != nil {
		log.Fatalf("Error creating space: %v", err)
	}

	// ユーザーにスペースを関連付ける
	user.Spaces = append(user.Spaces, space)

	// ユーザーをデータベースに保存（同時に関連するスペースも保存）
	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	log.Println("User and space created successfully")
}
