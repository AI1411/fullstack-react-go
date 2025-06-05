package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/logger"
	"github.com/AI1411/fullstack-react-go/internal/usecase"
)

type Notification interface {
	ListNotifications(c *gin.Context)
	GetNotification(c *gin.Context)
	GetNotificationsByUserID(c *gin.Context)
	CreateNotification(c *gin.Context)
	UpdateNotification(c *gin.Context)
	DeleteNotification(c *gin.Context)
	MarkAsRead(c *gin.Context)
}

type notificationHandler struct {
	l                   *logger.Logger
	notificationUseCase usecase.NotificationUseCase
}

func NewNotificationHandler(
	l *logger.Logger,
	notificationUseCase usecase.NotificationUseCase,
) Notification {
	return &notificationHandler{
		l:                   l,
		notificationUseCase: notificationUseCase,
	}
}

type NotificationResponse struct {
	ID                int32      `json:"id"`
	UserID            string     `json:"user_id"`
	Title             string     `json:"title"`
	Message           string     `json:"message"`
	NotificationType  string     `json:"notification_type"`
	RelatedEntityType *string    `json:"related_entity_type,omitempty"`
	RelatedEntityID   *string    `json:"related_entity_id,omitempty"`
	IsRead            bool       `json:"is_read"`
	ReadAt            *time.Time `json:"read_at,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

type CreateNotificationRequest struct {
	UserID            string  `json:"user_id" binding:"required"`
	Title             string  `json:"title" binding:"required,max=200"`
	Message           string  `json:"message" binding:"required"`
	NotificationType  string  `json:"notification_type" binding:"required,max=50"`
	RelatedEntityType *string `json:"related_entity_type,omitempty" binding:"omitempty,max=50"`
	RelatedEntityID   *string `json:"related_entity_id,omitempty" binding:"omitempty,max=100"`
}

type UpdateNotificationRequest struct {
	Title             string  `json:"title" binding:"required,max=200"`
	Message           string  `json:"message" binding:"required"`
	NotificationType  string  `json:"notification_type" binding:"required,max=50"`
	RelatedEntityType *string `json:"related_entity_type,omitempty" binding:"omitempty,max=50"`
	RelatedEntityID   *string `json:"related_entity_id,omitempty" binding:"omitempty,max=100"`
}

// ListNotifications @title 通知一覧取得
// @id ListNotifications
// @tags notification
// @accept json
// @produce json
// @Summary 通知一覧取得
// @Success 200 {array} NotificationResponse
// @Router /notifications [get]
func (h *notificationHandler) ListNotifications(c *gin.Context) {
	ctx := c.Request.Context()
	notifications, err := h.notificationUseCase.ListNotifications(ctx)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to list notifications")
		c.JSON(500, gin.H{"error": "Internal Server Error"})

		return
	}

	var response []*NotificationResponse

	for _, notification := range notifications {
		resp := &NotificationResponse{
			ID:                notification.ID,
			UserID:            notification.UserID,
			Title:             notification.Title,
			Message:           notification.Message,
			NotificationType:  notification.NotificationType,
			RelatedEntityType: notification.RelatedEntityType,
			RelatedEntityID:   notification.RelatedEntityID,
			IsRead:            notification.IsRead,
			ReadAt:            notification.ReadAt,
			CreatedAt:         notification.CreatedAt,
			UpdatedAt:         notification.UpdatedAt,
		}

		response = append(response, resp)
	}

	h.l.InfoContext(ctx, "Successfully listed notifications", "count", len(response))
	c.JSON(http.StatusOK, response)
}

// GetNotification @title 通知詳細取得
// @id GetNotification
// @tags notification
// @accept json
// @produce json
// @Param id path int true "通知ID"
// @Summary 通知詳細取得
// @Success 200 {object} NotificationResponse
// @Failure 404 {object} map[string]string
// @Router /notifications/{id} [get]
func (h *notificationHandler) GetNotification(c *gin.Context) {
	idStr := c.Param("id")
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Invalid notification ID", "notification_id_str", idStr)
		c.JSON(400, gin.H{"error": "Invalid notification ID"})

		return
	}

	notification, err := h.notificationUseCase.GetNotificationByID(ctx, int32(id))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Notification not found", "notification_id", id)
		c.JSON(404, gin.H{"error": "Notification not found"})

		return
	}

	response := &NotificationResponse{
		ID:                notification.ID,
		UserID:            notification.UserID,
		Title:             notification.Title,
		Message:           notification.Message,
		NotificationType:  notification.NotificationType,
		RelatedEntityType: notification.RelatedEntityType,
		RelatedEntityID:   notification.RelatedEntityID,
		IsRead:            notification.IsRead,
		ReadAt:            notification.ReadAt,
		CreatedAt:         notification.CreatedAt,
		UpdatedAt:         notification.UpdatedAt,
	}

	h.l.InfoContext(ctx, "Successfully retrieved notification", "notification_id", id)
	c.JSON(http.StatusOK, response)
}

// GetNotificationsByUserID @title ユーザーIDによる通知一覧取得
// @id GetNotificationsByUserID
// @tags notification
// @accept json
// @produce json
// @Param user_id path int true "ユーザーID"
// @Summary ユーザーIDによる通知一覧取得
// @Success 200 {array} NotificationResponse
// @Failure 404 {object} map[string]string
// @Router /notifications/user/{user_id} [get]
func (h *notificationHandler) GetNotificationsByUserID(c *gin.Context) {
	userIDStr := c.Param("user_id")
	ctx := c.Request.Context()

	userID, err := strconv.ParseInt(userIDStr, 10, 32)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Invalid user ID", "user_id_str", userIDStr)
		c.JSON(400, gin.H{"error": "Invalid user ID"})

		return
	}

	notifications, err := h.notificationUseCase.GetNotificationsByUserID(ctx, int32(userID))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to get notifications for user", "user_id", userID)
		c.JSON(500, gin.H{"error": "Internal Server Error"})

		return
	}

	var response []*NotificationResponse

	for _, notification := range notifications {
		resp := &NotificationResponse{
			ID:                notification.ID,
			UserID:            notification.UserID,
			Title:             notification.Title,
			Message:           notification.Message,
			NotificationType:  notification.NotificationType,
			RelatedEntityType: notification.RelatedEntityType,
			RelatedEntityID:   notification.RelatedEntityID,
			IsRead:            notification.IsRead,
			ReadAt:            notification.ReadAt,
			CreatedAt:         notification.CreatedAt,
			UpdatedAt:         notification.UpdatedAt,
		}

		response = append(response, resp)
	}

	h.l.InfoContext(ctx, "Successfully retrieved notifications for user", "user_id", userID, "count", len(response))
	c.JSON(http.StatusOK, response)
}

// CreateNotification @title 通知作成
// @id CreateNotification
// @tags notification
// @accept json
// @produce json
// @Param request body CreateNotificationRequest true "通知作成リクエスト"
// @Summary 通知作成
// @Success 201 {object} NotificationResponse
// @Failure 400 {object} map[string]string
// @Router /notifications [post]
func (h *notificationHandler) CreateNotification(c *gin.Context) {
	var req CreateNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	notification := &model.Notification{
		UserID:            req.UserID,
		Title:             req.Title,
		Message:           req.Message,
		NotificationType:  req.NotificationType,
		RelatedEntityType: req.RelatedEntityType,
		RelatedEntityID:   req.RelatedEntityID,
		IsRead:            false,
	}

	ctx := c.Request.Context()
	err := h.notificationUseCase.CreateNotification(ctx, notification)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to create notification")
		c.JSON(500, gin.H{"error": "Failed to create notification"})

		return
	}

	response := &NotificationResponse{
		ID:                notification.ID,
		UserID:            notification.UserID,
		Title:             notification.Title,
		Message:           notification.Message,
		NotificationType:  notification.NotificationType,
		RelatedEntityType: notification.RelatedEntityType,
		RelatedEntityID:   notification.RelatedEntityID,
		IsRead:            notification.IsRead,
		ReadAt:            notification.ReadAt,
		CreatedAt:         notification.CreatedAt,
		UpdatedAt:         notification.UpdatedAt,
	}

	h.l.InfoContext(ctx, "Successfully created notification", "notification_id", notification.ID)
	c.JSON(http.StatusCreated, response)
}

// UpdateNotification @title 通知更新
// @id UpdateNotification
// @tags notification
// @accept json
// @produce json
// @Param id path int true "通知ID"
// @Param request body UpdateNotificationRequest true "通知更新リクエスト"
// @Summary 通知更新
// @Success 200 {object} NotificationResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /notifications/{id} [put]
func (h *notificationHandler) UpdateNotification(c *gin.Context) {
	idStr := c.Param("id")
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Invalid notification ID", "notification_id_str", idStr)
		c.JSON(400, gin.H{"error": "Invalid notification ID"})

		return
	}

	var req UpdateNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Check if the notification exists
	existingNotification, err := h.notificationUseCase.GetNotificationByID(ctx, int32(id))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Notification not found", "notification_id", id)
		c.JSON(404, gin.H{"error": "Notification not found"})

		return
	}

	// Update the notification
	existingNotification.Title = req.Title
	existingNotification.Message = req.Message
	existingNotification.NotificationType = req.NotificationType
	existingNotification.RelatedEntityType = req.RelatedEntityType
	existingNotification.RelatedEntityID = req.RelatedEntityID

	err = h.notificationUseCase.UpdateNotification(ctx, existingNotification)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to update notification", "notification_id", id)
		c.JSON(500, gin.H{"error": "Failed to update notification"})

		return
	}

	response := &NotificationResponse{
		ID:                existingNotification.ID,
		UserID:            existingNotification.UserID,
		Title:             existingNotification.Title,
		Message:           existingNotification.Message,
		NotificationType:  existingNotification.NotificationType,
		RelatedEntityType: existingNotification.RelatedEntityType,
		RelatedEntityID:   existingNotification.RelatedEntityID,
		IsRead:            existingNotification.IsRead,
		ReadAt:            existingNotification.ReadAt,
		CreatedAt:         existingNotification.CreatedAt,
		UpdatedAt:         existingNotification.UpdatedAt,
	}

	h.l.InfoContext(ctx, "Successfully updated notification", "notification_id", id)
	c.JSON(http.StatusOK, response)
}

// DeleteNotification @title 通知削除
// @id DeleteNotification
// @tags notification
// @accept json
// @produce json
// @Param id path int true "通知ID"
// @Summary 通知削除
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /notifications/{id} [delete]
func (h *notificationHandler) DeleteNotification(c *gin.Context) {
	idStr := c.Param("id")
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Invalid notification ID", "notification_id_str", idStr)
		c.JSON(400, gin.H{"error": "Invalid notification ID"})

		return
	}

	// Check if the notification exists
	_, err = h.notificationUseCase.GetNotificationByID(ctx, int32(id))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Notification not found", "notification_id", id)
		c.JSON(404, gin.H{"error": "Notification not found"})

		return
	}

	err = h.notificationUseCase.DeleteNotification(ctx, int32(id))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to delete notification", "notification_id", id)
		c.JSON(500, gin.H{"error": "Failed to delete notification"})

		return
	}

	h.l.InfoContext(ctx, "Successfully deleted notification", "notification_id", id)
	c.Status(http.StatusNoContent)
}

// MarkAsRead @title 通知を既読にする
// @id MarkAsRead
// @tags notification
// @accept json
// @produce json
// @Param id path int true "通知ID"
// @Summary 通知を既読にする
// @Success 200 {object} NotificationResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /notifications/{id}/read [put]
func (h *notificationHandler) MarkAsRead(c *gin.Context) {
	idStr := c.Param("id")
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Invalid notification ID", "notification_id_str", idStr)
		c.JSON(400, gin.H{"error": "Invalid notification ID"})

		return
	}

	// Check if the notification exists
	notification, err := h.notificationUseCase.GetNotificationByID(ctx, int32(id))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Notification not found", "notification_id", id)
		c.JSON(404, gin.H{"error": "Notification not found"})

		return
	}

	// If already read, return the notification as is
	if notification.IsRead {
		response := &NotificationResponse{
			ID:                notification.ID,
			UserID:            notification.UserID,
			Title:             notification.Title,
			Message:           notification.Message,
			NotificationType:  notification.NotificationType,
			RelatedEntityType: notification.RelatedEntityType,
			RelatedEntityID:   notification.RelatedEntityID,
			IsRead:            notification.IsRead,
			ReadAt:            notification.ReadAt,
			CreatedAt:         notification.CreatedAt,
			UpdatedAt:         notification.UpdatedAt,
		}

		h.l.InfoContext(ctx, "Notification already marked as read", "notification_id", id)
		c.JSON(http.StatusOK, response)

		return
	}

	err = h.notificationUseCase.MarkAsRead(ctx, int32(id))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to mark notification as read", "notification_id", id)
		c.JSON(500, gin.H{"error": "Failed to mark notification as read"})

		return
	}

	// Get the updated notification
	updatedNotification, err := h.notificationUseCase.GetNotificationByID(ctx, int32(id))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to get updated notification", "notification_id", id)
		c.JSON(500, gin.H{"error": "Failed to get updated notification"})

		return
	}

	response := &NotificationResponse{
		ID:                updatedNotification.ID,
		UserID:            updatedNotification.UserID,
		Title:             updatedNotification.Title,
		Message:           updatedNotification.Message,
		NotificationType:  updatedNotification.NotificationType,
		RelatedEntityType: updatedNotification.RelatedEntityType,
		RelatedEntityID:   updatedNotification.RelatedEntityID,
		IsRead:            updatedNotification.IsRead,
		ReadAt:            updatedNotification.ReadAt,
		CreatedAt:         updatedNotification.CreatedAt,
		UpdatedAt:         updatedNotification.UpdatedAt,
	}

	h.l.InfoContext(ctx, "Successfully marked notification as read", "notification_id", id)
	c.JSON(http.StatusOK, response)
}
