package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AI1411/gen/internal/infra/database"
	"github.com/AI1411/gen/internal/model"
)

func main() {
	// データベース接続
	db, err := database.NewDatabaseConnection()
	if err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}

	// テストユーザー作成
	testUser := model.User{
		Name:     "テストユーザー",
		Email:    "test@example.com",
		Password: "password123",
	}

	// ユーザーが存在しない場合のみ作成
	var count int64
	db.Model(&model.User{}).Where("email = ?", testUser.Email).Count(&count)
	if count == 0 {
		result := db.Create(&testUser)
		if result.Error != nil {
			log.Printf("テストユーザー作成エラー: %v", result.Error)
		} else {
			log.Println("テストユーザーを作成しました")
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go言語、GORM、PostgreSQLを使ったAPIサーバーです")
	})

	// ユーザー一覧を表示するハンドラー
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		var users []model.User
		result := db.Find(&users)
		if result.Error != nil {
			http.Error(w, "ユーザーの取得に失敗しました", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "ユーザー一覧（%d人）:\n", len(users))
		for _, user := range users {
			fmt.Fprintf(w, "ID: %d, 名前: %s, メール: %s\n", user.ID, user.Name, user.Email)
		}
	})

	// サーバー起動
	log.Println("サーバーを起動しています。ポート: 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("サーバー起動エラー: %v", err)
	}
}
