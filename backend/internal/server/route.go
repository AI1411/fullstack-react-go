package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"github.com/AI1411/fullstack-react-go/internal/env"
	"github.com/AI1411/fullstack-react-go/internal/handler"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
	"github.com/AI1411/fullstack-react-go/internal/infra/logger"
)

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
	r.POST("/auth/register", authHandler.Register)

	// Swagger JSON エンドポイント
	r.GET("/docs", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.File("./docs/api/swagger.json")
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
