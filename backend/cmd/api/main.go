package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"github.com/AI1411/fullstack-react-go/internal/env"
	"github.com/AI1411/fullstack-react-go/internal/handler"
	"github.com/AI1411/fullstack-react-go/internal/infra/datastore"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
	"github.com/AI1411/fullstack-react-go/internal/infra/logger"
	"github.com/AI1411/fullstack-react-go/internal/middleware"
	"github.com/AI1411/fullstack-react-go/internal/usecase"
)

// Swagger メタデータ
// @title           農業災害支援システム API
// @version         1.0
// @description     農業災害の報告と支援申請を管理するためのAPI
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

// Module groups provide a way to organize dependencies
type Module struct {
	fx.Out

	Logger                    *logger.Logger
	EnvValues                 *env.Values
	DBClient                  db.Client
	GinEngine                 *gin.Engine
	DisasterRepo              datastore.DisasterRepository
	PrefectureRepo            datastore.PrefectureRepository
	TimelineRepo              datastore.TimelineRepository
	SupportApplicationRepo    datastore.SupportApplicationRepository
	DamageLevelRepo           datastore.DamageLevelRepository
	FacilityEquipmentRepo     datastore.FacilityEquipmentRepository
	NotificationRepo          datastore.NotificationRepository
	OrganizationRepo          datastore.OrganizationRepository
	DisasterUsecase           usecase.DisasterUseCase
	PrefectureUsecase         usecase.PrefectureUseCase
	TimelineUsecase           usecase.TimelineUseCase
	SupportApplicationUsecase usecase.SupportApplicationUseCase
	DamageLevelUsecase        usecase.DamageLevelUseCase
	FacilityEquipmentUsecase  usecase.FacilityEquipmentUseCase
	NotificationUsecase       usecase.NotificationUseCase
	OrganizationUsecase       usecase.OrganizationUseCase
	DisasterHandler           handler.Disaster
	PrefectureHandler         handler.Prefecture
	TimelineHandler           handler.Timeline
	SupportApplicationHandler handler.SupportApplication
	DamageLevelHandler        handler.DamageLevel
	FacilityEquipmentHandler  handler.FacilityEquipment
	NotificationHandler       handler.Notification
	OrganizationHandler       handler.Organization
	UserRepo                  datastore.UserRepository
	UserUsecase               usecase.UserUseCase
	UserHandler               handler.User
	AuthHandler               handler.Auth
}

// ProvideLogger creates a new logger instance
func ProvideLogger() *logger.Logger {
	return logger.New(logger.DefaultConfig())
}

// ProvideEnvValues creates a new env values instance
func ProvideEnvValues() (*env.Values, error) {
	return env.NewValues()
}

// ProvideDBClient creates a new database client
func ProvideDBClient(lc fx.Lifecycle, l *logger.Logger) (db.Client, error) {
	e, err := env.NewValues()
	if err != nil {
		l.Error("failed to load environment variables", "error", err)
		return nil, err
	}
	dbClient, err := db.NewSQLHandler(&db.DatabaseConfig{
		Host:            e.DatabaseHost,
		Port:            e.DatabasePort,
		User:            e.DatabaseUsername,
		Password:        e.DatabasePassword,
		DBName:          e.DatabaseName,
		SSLMode:         "disable",
		Timezone:        "Asia/Tokyo",
		MaxIdleConns:    e.ConnectionMaxIdle,
		MaxOpenConns:    e.ConnectionMaxOpen,
		ConnMaxLifetime: e.ConnectionMaxLifetime,
	}, l)
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

// ProvidePrefectureRepository creates a new prefecture repository
func ProvidePrefectureRepository(dbClient db.Client) datastore.PrefectureRepository {
	ctx := context.Background()
	return datastore.NewPrefectureRepository(ctx, dbClient)
}

// ProvideTimelineRepository creates a new timeline repository
func ProvideTimelineRepository(dbClient db.Client) datastore.TimelineRepository {
	ctx := context.Background()
	return datastore.NewTimelineRepository(ctx, dbClient)
}

// ProvideSupportApplicationRepository creates a new support application repository
func ProvideSupportApplicationRepository(dbClient db.Client) datastore.SupportApplicationRepository {
	ctx := context.Background()
	return datastore.NewSupportApplicationRepository(ctx, dbClient)
}

// ProvideDisasterUseCase creates a new disaster use case
func ProvideDisasterUseCase(repo datastore.DisasterRepository) usecase.DisasterUseCase {
	return usecase.NewDisasterUseCase(repo)
}

// ProvidePrefectureUseCase creates a new prefecture use case
func ProvidePrefectureUseCase(repo datastore.PrefectureRepository) usecase.PrefectureUseCase {
	return usecase.NewPrefectureUseCase(repo)
}

// ProvideTimelineUseCase creates a new timeline use case
func ProvideTimelineUseCase(repo datastore.TimelineRepository) usecase.TimelineUseCase {
	return usecase.NewTimelineUseCase(repo)
}

// ProvideSupportApplicationUseCase creates a new support application use case
func ProvideSupportApplicationUseCase(repo datastore.SupportApplicationRepository) usecase.SupportApplicationUseCase {
	return usecase.NewSupportApplicationUseCase(repo)
}

// ProvideDisasterHandler creates a new disaster handler
func ProvideDisasterHandler(l *logger.Logger, usecase usecase.DisasterUseCase) handler.Disaster {
	return handler.NewDisasterHandler(l, usecase)
}

// ProvidePrefectureHandler creates a new prefecture handler
func ProvidePrefectureHandler(l *logger.Logger, usecase usecase.PrefectureUseCase) handler.Prefecture {
	return handler.NewPrefectureHandler(l, usecase)
}

// ProvideTimelineHandler creates a new timeline handler
func ProvideTimelineHandler(l *logger.Logger, usecase usecase.TimelineUseCase) handler.Timeline {
	return handler.NewTimelineHandler(l, usecase)
}

// ProvideSupportApplicationHandler creates a new support application handler
func ProvideSupportApplicationHandler(l *logger.Logger, usecase usecase.SupportApplicationUseCase) handler.SupportApplication {
	return handler.NewSupportApplicationHandler(l, usecase)
}

// ProvideDamageLevelRepository creates a new damage level repository
func ProvideDamageLevelRepository(dbClient db.Client) datastore.DamageLevelRepository {
	return datastore.NewDamageLevelRepository(context.Background(), dbClient)
}

// ProvideDamageLevelUseCase creates a new damage level usecase
func ProvideDamageLevelUseCase(repo datastore.DamageLevelRepository) usecase.DamageLevelUseCase {
	return usecase.NewDamageLevelUseCase(repo)
}

// ProvideDamageLevelHandler creates a new damage level handler
func ProvideDamageLevelHandler(l *logger.Logger, usecase usecase.DamageLevelUseCase) handler.DamageLevel {
	return handler.NewDamageLevelHandler(l, usecase)
}

// ProvideFacilityEquipmentRepository creates a new facility equipment repository
func ProvideFacilityEquipmentRepository(dbClient db.Client) datastore.FacilityEquipmentRepository {
	return datastore.NewFacilityEquipmentRepository(context.Background(), dbClient)
}

// ProvideFacilityEquipmentUseCase creates a new facility equipment use case
func ProvideFacilityEquipmentUseCase(repo datastore.FacilityEquipmentRepository) usecase.FacilityEquipmentUseCase {
	return usecase.NewFacilityEquipmentUseCase(repo)
}

// ProvideFacilityEquipmentHandler creates a new facility equipment handler
func ProvideFacilityEquipmentHandler(l *logger.Logger, usecase usecase.FacilityEquipmentUseCase) handler.FacilityEquipment {
	return handler.NewFacilityEquipmentHandler(l, usecase)
}

// ProvideNotificationRepository creates a new notification repository
func ProvideNotificationRepository(dbClient db.Client) datastore.NotificationRepository {
	return datastore.NewNotificationRepository(context.Background(), dbClient)
}

// ProvideNotificationUseCase creates a new notification use case
func ProvideNotificationUseCase(repo datastore.NotificationRepository) usecase.NotificationUseCase {
	return usecase.NewNotificationUseCase(repo)
}

// ProvideNotificationHandler creates a new notification handler
func ProvideNotificationHandler(l *logger.Logger, usecase usecase.NotificationUseCase) handler.Notification {
	return handler.NewNotificationHandler(l, usecase)
}

// ProvideOrganizationRepository creates a new organization repository
func ProvideOrganizationRepository(client db.Client) datastore.OrganizationRepository {
	ctx := context.Background()
	return datastore.NewOrganizationRepository(ctx, client)
}

// ProvideOrganizationUseCase creates a new organization usecase
func ProvideOrganizationUseCase(repo datastore.OrganizationRepository) usecase.OrganizationUseCase {
	return usecase.NewOrganizationUseCase(repo)
}

// ProvideOrganizationHandler creates a new organization handler
func ProvideOrganizationHandler(l *logger.Logger, usecase usecase.OrganizationUseCase) handler.Organization {
	return handler.NewOrganizationHandler(l, usecase)
}

// ProvideUserRepository creates a new user repository
func ProvideUserRepository(client db.Client) datastore.UserRepository {
	ctx := context.Background()
	return datastore.NewUserRepository(ctx, client)
}

// ProvideUserUseCase creates a new user usecase
func ProvideUserUseCase(repo datastore.UserRepository) usecase.UserUseCase {
	return usecase.NewUserUseCase(repo)
}

// ProvideUserHandler creates a new user handler
func ProvideUserHandler(l *logger.Logger, usecase usecase.UserUseCase) handler.User {
	return handler.NewUserHandler(l, usecase)
}

// ProvideAuthHandler creates a new auth handler
func ProvideAuthHandler(ctx context.Context, l *logger.Logger, usecase usecase.UserUseCase, env *env.Values) (handler.Auth, error) {
	return handler.NewAuthHandler(ctx, l, usecase, env)
}

// RegisterRoutes registers all HTTP routes
func RegisterRoutes(
	lc fx.Lifecycle,
	r *gin.Engine,
	l *logger.Logger,
	dbClient db.Client,
	env *env.Values,
	disasterHandler handler.Disaster,
	prefectureHandler handler.Prefecture,
	timelineHandler handler.Timeline,
	supportApplicationHandler handler.SupportApplication,
	damageLevelHandler handler.DamageLevel,
	facilityEquipmentHandler handler.FacilityEquipment,
	notificationHandler handler.Notification,
	organizationHandler handler.Organization,
	userHandler handler.User,
	authHandler handler.Auth,
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

	// 災害関連のルート
	r.GET("/disasters", disasterHandler.ListDisasters)
	r.GET("/disasters/:id", disasterHandler.GetDisaster)
	r.POST("/disasters", disasterHandler.CreateDisaster)
	r.PUT("/disasters/:id", disasterHandler.UpdateDisaster)
	r.DELETE("/disasters/:id", disasterHandler.DeleteDisaster)

	// 都道府県関連のルート
	r.GET("/prefectures", prefectureHandler.ListPrefectures)
	r.GET("/prefectures/:code", prefectureHandler.GetPrefecture)

	// タイムライン関連のルート
	r.GET("/disasters/:id/timelines", timelineHandler.GetTimelinesByDisasterID)

	// 支援申請関連のルート
	r.GET("/support-applications", supportApplicationHandler.ListSupportApplications)
	r.GET("/support-applications/:id", supportApplicationHandler.GetSupportApplication)
	r.POST("/support-applications", supportApplicationHandler.CreateSupportApplication)

	// 被害程度関連のルート
	r.GET("/damage-levels", damageLevelHandler.ListDamageLevels)
	r.GET("/damage-levels/:id", damageLevelHandler.GetDamageLevel)
	r.POST("/damage-levels", damageLevelHandler.CreateDamageLevel)
	r.PUT("/damage-levels/:id", damageLevelHandler.UpdateDamageLevel)
	r.DELETE("/damage-levels/:id", damageLevelHandler.DeleteDamageLevel)

	// 施設設備関連のルート
	r.GET("/facility-equipment", facilityEquipmentHandler.ListFacilityEquipments)
	r.GET("/facility-equipment/:id", facilityEquipmentHandler.GetFacilityEquipment)
	r.POST("/facility-equipment", facilityEquipmentHandler.CreateFacilityEquipment)
	r.PUT("/facility-equipment/:id", facilityEquipmentHandler.UpdateFacilityEquipment)
	r.DELETE("/facility-equipment/:id", facilityEquipmentHandler.DeleteFacilityEquipment)

	// 通知関連のルート
	r.GET("/notifications", notificationHandler.ListNotifications)
	r.GET("/notifications/:id", notificationHandler.GetNotification)
	r.GET("/notifications/user/:user_id", notificationHandler.GetNotificationsByUserID)
	r.POST("/notifications", notificationHandler.CreateNotification)
	r.PUT("/notifications/:id", notificationHandler.UpdateNotification)
	r.DELETE("/notifications/:id", notificationHandler.DeleteNotification)
	r.PUT("/notifications/:id/read", notificationHandler.MarkAsRead)

	// 組織関連のルート
	r.GET("/organizations", organizationHandler.ListOrganizations)
	r.GET("/organizations/:id", organizationHandler.GetOrganization)
	r.POST("/organizations", organizationHandler.CreateOrganization)
	r.PUT("/organizations/:id", organizationHandler.UpdateOrganization)
	r.DELETE("/organizations/:id", organizationHandler.DeleteOrganization)

	// ユーザー関連のルート
	r.GET("/users", userHandler.ListUsers)
	r.GET("/users/:id", userHandler.GetUser)
	r.POST("/users", userHandler.CreateUser)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	// 認証関連のルート
	r.GET("/auth/login", authHandler.Login)
	r.GET("/auth/callback", authHandler.Callback)
	r.POST("/auth/logout", authHandler.Logout)

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
				l.Info(fmt.Sprintf("Starting server on :%s", env.ServerPort))
				if err := r.Run(fmt.Sprintf(":%s", env.ServerPort)); err != nil {
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

// ProvideAppContext provides a background context for the application
func ProvideAppContext() context.Context {
	return context.Background()
}

func main() {
	app := fx.New(
		// Provide all the constructors needed by the application
		fx.Provide(
			ProvideAppContext,
			ProvideLogger,
			ProvideEnvValues,
			ProvideDBClient,
			ProvideGinEngine,
			ProvideDisasterRepository,
			ProvidePrefectureRepository,
			ProvideTimelineRepository,
			ProvideSupportApplicationRepository,
			ProvideDamageLevelRepository,
			ProvideFacilityEquipmentRepository,
			ProvideNotificationRepository,
			ProvideOrganizationRepository,
			ProvideUserRepository,
			ProvideDisasterUseCase,
			ProvidePrefectureUseCase,
			ProvideTimelineUseCase,
			ProvideSupportApplicationUseCase,
			ProvideDamageLevelUseCase,
			ProvideFacilityEquipmentUseCase,
			ProvideNotificationUseCase,
			ProvideOrganizationUseCase,
			ProvideUserUseCase,
			ProvideDisasterHandler,
			ProvidePrefectureHandler,
			ProvideTimelineHandler,
			ProvideSupportApplicationHandler,
			ProvideDamageLevelHandler,
			ProvideFacilityEquipmentHandler,
			ProvideNotificationHandler,
			ProvideOrganizationHandler,
			ProvideUserHandler,
			ProvideAuthHandler,
		),
		// Register the lifecycle hooks
		fx.Invoke(RegisterRoutes),
	)

	// Run the application
	app.Run()
}
