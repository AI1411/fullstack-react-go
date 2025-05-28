package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"github.com/AI1411/fullstack-react-go/internal/handler"
	"github.com/AI1411/fullstack-react-go/internal/infra/datastore"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
	"github.com/AI1411/fullstack-react-go/internal/infra/logger"
	"github.com/AI1411/fullstack-react-go/internal/middleware"
	"github.com/AI1411/fullstack-react-go/internal/usecase"
)

// Module groups provide a way to organize dependencies
type Module struct {
	fx.Out

	Logger          *logger.Logger
	DBClient        db.Client
	GinEngine       *gin.Engine
	DisasterRepo    datastore.DisasterRepository
	DisasterUsecase usecase.DisasterUseCase
	DisasterHandler handler.Disaster
}

// ProvideLogger creates a new logger instance
func ProvideLogger() *logger.Logger {
	return logger.New(logger.DefaultConfig())
}

// ProvideDBClient creates a new database client
func ProvideDBClient(lc fx.Lifecycle, l *logger.Logger) (db.Client, error) {
	dbClient, err := db.NewSqlHandler(db.DefaultDatabaseConfig(), l)
	if err != nil {
		l.Error("failed to connect to database", "error", err)
		return nil, err
	}

	// Register lifecycle hooks for the database client
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			l.Info("Closing database connection")
			return nil // Add proper cleanup if needed
		},
	})
	return dbClient, nil
}

// ProvideGinEngine creates and configures a new Gin engine
func ProvideGinEngine(l *logger.Logger) *gin.Engine {
	r := gin.Default()

	// ミドルウェアの設定
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.NewLogging(l))
	r.Use(middleware.CORSMiddleware())

	return r
}

// ProvideDisasterRepository creates a new disaster repository
func ProvideDisasterRepository(dbClient db.Client) datastore.DisasterRepository {
	ctx := context.Background()
	return datastore.NewDisasterRepository(ctx, dbClient)
}

// ProvideDisasterUseCase creates a new disaster use case
func ProvideDisasterUseCase(repo datastore.DisasterRepository) usecase.DisasterUseCase {
	return usecase.NewDisasterUseCase(repo)
}

// ProvideDisasterHandler creates a new disaster handler
func ProvideDisasterHandler(usecase usecase.DisasterUseCase) handler.Disaster {
	return handler.NewDisasterHandler(usecase)
}

// RegisterRoutes registers all HTTP routes
func RegisterRoutes(
	lc fx.Lifecycle,
	r *gin.Engine,
	l *logger.Logger,
	dbClient db.Client,
	disasterHandler handler.Disaster,
) {
	// Context for health check
	ctx := context.Background()

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

	// Swagger JSON エンドポイント
	r.GET("/docs", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.File("./docs/swagger.json")
	})

	// Register lifecycle hooks for the HTTP server
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Start HTTP server in a goroutine so it doesn't block
			go func() {
				l.Info("Starting server on :8080")
				if err := r.Run(":8080"); err != nil {
					l.Error("Failed to start server", "error", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			l.Info("Shutting down server")
			return nil // Add proper cleanup if needed
		},
	})
}

func main() {
	app := fx.New(
		// Provide all the constructors needed by the application
		fx.Provide(
			ProvideLogger,
			ProvideDBClient,
			ProvideGinEngine,
			ProvideDisasterRepository,
			ProvideDisasterUseCase,
			ProvideDisasterHandler,
		),
		// Register the lifecycle hooks
		fx.Invoke(RegisterRoutes),
	)

	// Run the application
	app.Run()
}
