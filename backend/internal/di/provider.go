package di

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	domain "github.com/AI1411/fullstack-react-go/internal/domain/repository"
	"github.com/AI1411/fullstack-react-go/internal/env"
	"github.com/AI1411/fullstack-react-go/internal/handler"
	"github.com/AI1411/fullstack-react-go/internal/infra/auth"
	"github.com/AI1411/fullstack-react-go/internal/infra/datastore"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
	"github.com/AI1411/fullstack-react-go/internal/infra/logger"
	middleware2 "github.com/AI1411/fullstack-react-go/internal/server/middleware"
	"github.com/AI1411/fullstack-react-go/internal/usecase"
)

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

func ProvideJWTClient(env *env.Values) (domain.JWT, error) {
	jwtConfig := auth.JWTConfig{
		SecretKey:  env.Auth.JWTSecret,
		Expiration: time.Duration(env.JWTExpiration),
		Issuer:     env.Auth.JWTIssuer,
	}

	return auth.NewJWTClient(jwtConfig), nil
}

// ProvideGinEngine creates and configures a new Gin engine
func ProvideGinEngine(l *logger.Logger) *gin.Engine {
	r := gin.Default()

	// ミドルウェアの設定
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware2.NewLogging(l))
	r.Use(middleware2.CORSMiddleware())

	return r
}

// ProvideDisasterRepository creates a new disaster repository
func ProvideDisasterRepository(dbClient db.Client) datastore.DisasterRepository {
	ctx := context.Background()
	return datastore.NewDisasterRepository(ctx, dbClient)
}

// ProvidePrefectureRepository creates a new prefecture repository
func ProvidePrefectureRepository(dbClient db.Client) domain.PrefectureRepository {
	ctx := context.Background()
	return datastore.NewPrefectureRepository(ctx, dbClient)
}

// ProvideTimelineRepository creates a new timeline repository
func ProvideTimelineRepository(dbClient db.Client) datastore.TimelineRepository {
	ctx := context.Background()
	return datastore.NewTimelineRepository(ctx, dbClient)
}

// ProvideSupportApplicationRepository creates a new support application repository
func ProvideSupportApplicationRepository(dbClient db.Client) domain.SupportApplicationRepository {
	ctx := context.Background()
	return datastore.NewSupportApplicationRepository(ctx, dbClient)
}

// ProvideEmailVarificationTokenRepository creates a new email verification token repository
func ProvideEmailVarificationTokenRepository(ctx context.Context, dbClient db.Client) domain.EmailVarificationTokenRepository {
	return datastore.NewEmailVarificationTokenRepository(ctx, dbClient)
}

// ProvideDisasterUseCase creates a new disaster use case
func ProvideDisasterUseCase(repo datastore.DisasterRepository) usecase.DisasterUseCase {
	return usecase.NewDisasterUseCase(repo)
}

// ProvidePrefectureUseCase creates a new prefecture use case
func ProvidePrefectureUseCase(repo domain.PrefectureRepository) usecase.PrefectureUseCase {
	return usecase.NewPrefectureUseCase(repo)
}

// ProvideTimelineUseCase creates a new timeline use case
func ProvideTimelineUseCase(repo datastore.TimelineRepository) usecase.TimelineUseCase {
	return usecase.NewTimelineUseCase(repo)
}

// ProvideSupportApplicationUseCase creates a new support application use case
func ProvideSupportApplicationUseCase(repo domain.SupportApplicationRepository) usecase.SupportApplicationUseCase {
	return usecase.NewSupportApplicationUseCase(repo)
}

// ProvideEmailVarificationTokenUseCase creates a new support application use case
func ProvideEmailVarificationTokenUseCase(repo domain.EmailVarificationTokenRepository) usecase.EmailVarificationTokenUsecase {
	return usecase.NewEmailVarificationTokenUsecase(repo)
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
func ProvideDamageLevelRepository(dbClient db.Client) domain.DamageLevelRepository {
	return datastore.NewDamageLevelRepository(context.Background(), dbClient)
}

// ProvideDamageLevelUseCase creates a new damage level usecase
func ProvideDamageLevelUseCase(repo domain.DamageLevelRepository) usecase.DamageLevelUseCase {
	return usecase.NewDamageLevelUseCase(repo)
}

// ProvideDamageLevelHandler creates a new damage level handler
func ProvideDamageLevelHandler(l *logger.Logger, usecase usecase.DamageLevelUseCase) handler.DamageLevel {
	return handler.NewDamageLevelHandler(l, usecase)
}

// ProvideFacilityEquipmentRepository creates a new facility equipment repository
func ProvideFacilityEquipmentRepository(dbClient db.Client) domain.FacilityEquipmentRepository {
	return datastore.NewFacilityEquipmentRepository(context.Background(), dbClient)
}

// ProvideFacilityEquipmentUseCase creates a new facility equipment use case
func ProvideFacilityEquipmentUseCase(repo domain.FacilityEquipmentRepository) usecase.FacilityEquipmentUseCase {
	return usecase.NewFacilityEquipmentUseCase(repo)
}

// ProvideFacilityEquipmentHandler creates a new facility equipment handler
func ProvideFacilityEquipmentHandler(l *logger.Logger, usecase usecase.FacilityEquipmentUseCase) handler.FacilityEquipment {
	return handler.NewFacilityEquipmentHandler(l, usecase)
}

// ProvideNotificationRepository creates a new notification repository
func ProvideNotificationRepository(dbClient db.Client) domain.NotificationRepository {
	return datastore.NewNotificationRepository(context.Background(), dbClient)
}

// ProvideNotificationUseCase creates a new notification use case
func ProvideNotificationUseCase(repo domain.NotificationRepository) usecase.NotificationUseCase {
	return usecase.NewNotificationUseCase(repo)
}

// ProvideNotificationHandler creates a new notification handler
func ProvideNotificationHandler(l *logger.Logger, usecase usecase.NotificationUseCase) handler.Notification {
	return handler.NewNotificationHandler(l, usecase)
}

// ProvideOrganizationRepository creates a new organization repository
func ProvideOrganizationRepository(client db.Client) domain.OrganizationRepository {
	ctx := context.Background()
	return datastore.NewOrganizationRepository(ctx, client)
}

// ProvideOrganizationUseCase creates a new organization usecase
func ProvideOrganizationUseCase(repo domain.OrganizationRepository) usecase.OrganizationUseCase {
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

// ProvideEmailHistoryRepository creates a new email history repository
func ProvideEmailHistoryRepository(client db.Client) domain.EmailHistoryRepository {
	ctx := context.Background()
	return datastore.NewEmailHistoryRepository(ctx, client)
}

// ProvideUserUseCase creates a new user usecase
func ProvideUserUseCase(repo datastore.UserRepository, emailRepo domain.EmailHistoryRepository, emailVarificationTokenRepo domain.EmailVarificationTokenRepository) usecase.UserUseCase {
	return usecase.NewUserUseCase(repo, emailRepo, emailVarificationTokenRepo)
}

// ProvideUserHandler creates a new user handler
func ProvideUserHandler(l *logger.Logger, usecase usecase.UserUseCase) handler.User {
	return handler.NewUserHandler(l, usecase)
}

// ProvideAuthUsecase creates a new auth usecase
func ProvideAuthUsecase(jwtClient domain.JWT, emailVarificationTokenRepo domain.EmailVarificationTokenRepository) usecase.AuthUsecase {
	return usecase.NewAuthUsecase(jwtClient, emailVarificationTokenRepo)
}

// ProvideAuthHandler creates a new auth handler
func ProvideAuthHandler(l *logger.Logger, env *env.Values, userUseCase usecase.UserUseCase, authUsecase usecase.AuthUsecase, emailVarificationTokenUsecase usecase.EmailVarificationTokenUsecase) (handler.Auth, error) {
	return handler.NewAuthHandler(l, env, userUseCase, authUsecase, emailVarificationTokenUsecase)
}

// ProvideAppContext provides a background context for the application
func ProvideAppContext() context.Context {
	return context.Background()
}

func Provider() fx.Option {
	return fx.Provide(
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
		ProvideEmailHistoryRepository,
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
		ProvideJWTClient,
		ProvideAuthUsecase,
		ProvideAuthHandler,
		ProvideEmailVarificationTokenRepository,
		ProvideEmailVarificationTokenUseCase,
	)
}
