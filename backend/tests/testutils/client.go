package testutils

import (
	"context"
	"testing"

	"github.com/AI1411/fullstack-react-go/internal/infra/db"
	applogger "github.com/AI1411/fullstack-react-go/internal/infra/logger"
)

// SetupTestDB テスト用のDBクライアントをセットアップする
func SetupTestDB(t *testing.T) db.Client {
	client, err := db.NewSQLHandler(&db.DatabaseConfig{
		Host:     "db-test",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DBName:   "gen_test",
		SSLMode:  "disable",
		Timezone: "Asia/Tokyo",
	}, applogger.New(applogger.DefaultConfig()))
	if err != nil {
		t.Errorf("failed to connect to test database: %v", err)
	}

	return client
}

// TruncateAllTables テスト用のDBの全テーブルをトランケートする
func TruncateAllTables(t *testing.T, client db.Client) {
	// トランザクションを開始
	tx := client.Conn(context.Background()).Begin()
	if tx.Error != nil {
		t.Fatalf("failed to begin transaction: %v", tx.Error)
	}

	// 全テーブルをトランケート
	if err := tx.Exec("TRUNCATE TABLE prefectures, municipalities, damage_levels RESTART IDENTITY CASCADE").Error; err != nil {
		tx.Rollback()
		t.Fatalf("failed to truncate tables: %v", err)
	}

	// トランザクションをコミット
	if err := tx.Commit().Error; err != nil {
		t.Fatalf("failed to commit transaction: %v", err)
	}
}
