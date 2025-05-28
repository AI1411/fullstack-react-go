package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AI1411/fullstack-react-go/internal/handler"
	"github.com/AI1411/fullstack-react-go/internal/infra/datastore"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
	"github.com/AI1411/fullstack-react-go/internal/infra/logger"
	"github.com/AI1411/fullstack-react-go/internal/middleware"
	"github.com/AI1411/fullstack-react-go/internal/usecase"
)

func main() {
	ctx := context.Background()
	l := logger.New(logger.DefaultConfig())
	// データベース接続
	dbClient, err := db.NewSqlHandler(db.DefaultDatabaseConfig(), l)
	if err != nil {
		l.Error("failed to connect to database", "error", err)
	}

	// Ginの初期化
	r := gin.Default()

	// ミドルウェアの設定
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.NewLogging(l))
	r.Use(middleware.CORSMiddleware())

	// repository
	disasterRepo := datastore.NewDisasterRepository(ctx, dbClient)

	// usecase
	disasterUsecase := usecase.NewDisasterUseCase(disasterRepo)

	// handler
	disasterHandler := handler.NewDisasterHandler(disasterUsecase)

	// ヘルスチェックエンドポイント
	r.GET("/health", func(c *gin.Context) {
		if err := dbClient.Ping(ctx); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "unhealthy",
				"error":  "database ping failed",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   "healthy",
			"database": "connected",
		})
	})

	r.GET("/disasters", disasterHandler.ListDisasters)

	// サーバー起動
	l.Info("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		l.Error("Failed to start server", "error", err)
	}
}
