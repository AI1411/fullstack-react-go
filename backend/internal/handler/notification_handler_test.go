package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/handler"
	"github.com/AI1411/fullstack-react-go/internal/infra/logger"
	mockusecase "github.com/AI1411/fullstack-react-go/tests/mock/usecase"
)

func setupNotificationTest(t *testing.T) (*gin.Engine, *mockusecase.MockNotificationUseCase, handler.Notification) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	ctrl := gomock.NewController(t)
	mockUseCase := mockusecase.NewMockNotificationUseCase(ctrl)
	l := logger.New(logger.DefaultConfig())
	h := handler.NewNotificationHandler(l, mockUseCase)
	return r, mockUseCase, h
}

func TestNotificationHandler_ListNotifications(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupNotificationTest(t)
	r.GET("/notifications", h.ListNotifications)

	// Test cases
	tests := []struct {
		name           string
		mockSetup      func(mockUseCase *mockusecase.MockNotificationUseCase)
		expectedStatus int
		expectedBody   []*handler.NotificationResponse
	}{
		{
			name: "Success",
			mockSetup: func(mockUseCase *mockusecase.MockNotificationUseCase) {
				relatedEntityType := "disaster"
				relatedEntityID := "123"
				readAt := time.Now()

				notifications := []*model.Notification{
					{
						ID:                1,
						UserID:            1,
						Title:             "Test Notification 1",
						Message:           "This is a test notification",
						NotificationType:  "alert",
						RelatedEntityType: &relatedEntityType,
						RelatedEntityID:   &relatedEntityID,
						IsRead:            true,
						ReadAt:            &readAt,
						CreatedAt:         time.Now(),
						UpdatedAt:         time.Now(),
					},
					{
						ID:               2,
						UserID:           2,
						Title:            "Test Notification 2",
						Message:          "This is another test notification",
						NotificationType: "info",
						IsRead:           false,
						CreatedAt:        time.Now(),
						UpdatedAt:        time.Now(),
					},
				}
				mockUseCase.EXPECT().ListNotifications(gomock.Any()).Return(notifications, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: []*handler.NotificationResponse{
				{
					ID:                1,
					UserID:            1,
					Title:             "Test Notification 1",
					Message:           "This is a test notification",
					NotificationType:  "alert",
					RelatedEntityType: strPtr("disaster"),
					RelatedEntityID:   strPtr("123"),
					IsRead:            true,
				},
				{
					ID:               2,
					UserID:           2,
					Title:            "Test Notification 2",
					Message:          "This is another test notification",
					NotificationType: "info",
					IsRead:           false,
				},
			},
		},
		{
			name: "Error",
			mockSetup: func(mockUseCase *mockusecase.MockNotificationUseCase) {
				mockUseCase.EXPECT().ListNotifications(gomock.Any()).Return(nil, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   nil,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup(mockUseCase)

			// Make request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/notifications", nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response []*handler.NotificationResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, len(tt.expectedBody), len(response))

				for i, expected := range tt.expectedBody {
					assert.Equal(t, expected.ID, response[i].ID)
					assert.Equal(t, expected.UserID, response[i].UserID)
					assert.Equal(t, expected.Title, response[i].Title)
					assert.Equal(t, expected.Message, response[i].Message)
					assert.Equal(t, expected.NotificationType, response[i].NotificationType)
					assert.Equal(t, expected.IsRead, response[i].IsRead)

					// Check optional fields only if they exist in the expected response
					if expected.RelatedEntityType != nil {
						assert.Equal(t, *expected.RelatedEntityType, *response[i].RelatedEntityType)
					}
					if expected.RelatedEntityID != nil {
						assert.Equal(t, *expected.RelatedEntityID, *response[i].RelatedEntityID)
					}
				}
			}
		})
	}
}

func TestNotificationHandler_GetNotification(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupNotificationTest(t)
	r.GET("/notifications/:id", h.GetNotification)

	// Test cases
	tests := []struct {
		name           string
		notificationID string
		mockSetup      func(mockUseCase *mockusecase.MockNotificationUseCase)
		expectedStatus int
		expectedBody   *handler.NotificationResponse
	}{
		{
			name:           "Success",
			notificationID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockNotificationUseCase) {
				relatedEntityType := "disaster"
				relatedEntityID := "123"
				readAt := time.Now()

				notification := &model.Notification{
					ID:                1,
					UserID:            1,
					Title:             "Test Notification",
					Message:           "This is a test notification",
					NotificationType:  "alert",
					RelatedEntityType: &relatedEntityType,
					RelatedEntityID:   &relatedEntityID,
					IsRead:            true,
					ReadAt:            &readAt,
					CreatedAt:         time.Now(),
					UpdatedAt:         time.Now(),
				}
				mockUseCase.EXPECT().GetNotificationByID(gomock.Any(), int32(1)).Return(notification, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.NotificationResponse{
				ID:                1,
				UserID:            1,
				Title:             "Test Notification",
				Message:           "This is a test notification",
				NotificationType:  "alert",
				RelatedEntityType: strPtr("disaster"),
				RelatedEntityID:   strPtr("123"),
				IsRead:            true,
			},
		},
		{
			name:           "Invalid ID",
			notificationID: "invalid",
			mockSetup:      func(mockUseCase *mockusecase.MockNotificationUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name:           "Not Found",
			notificationID: "999",
			mockSetup: func(mockUseCase *mockusecase.MockNotificationUseCase) {
				mockUseCase.EXPECT().GetNotificationByID(gomock.Any(), int32(999)).Return(nil, errors.New("not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   nil,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup(mockUseCase)

			// Make request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/notifications/"+tt.notificationID, nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response handler.NotificationResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody.ID, response.ID)
				assert.Equal(t, tt.expectedBody.UserID, response.UserID)
				assert.Equal(t, tt.expectedBody.Title, response.Title)
				assert.Equal(t, tt.expectedBody.Message, response.Message)
				assert.Equal(t, tt.expectedBody.NotificationType, response.NotificationType)
				assert.Equal(t, tt.expectedBody.IsRead, response.IsRead)

				// Check optional fields only if they exist in the expected response
				if tt.expectedBody.RelatedEntityType != nil {
					assert.Equal(t, *tt.expectedBody.RelatedEntityType, *response.RelatedEntityType)
				}
				if tt.expectedBody.RelatedEntityID != nil {
					assert.Equal(t, *tt.expectedBody.RelatedEntityID, *response.RelatedEntityID)
				}
			}
		})
	}
}

func TestNotificationHandler_GetNotificationsByUserID(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupNotificationTest(t)
	r.GET("/notifications/user/:user_id", h.GetNotificationsByUserID)

	// Test cases
	tests := []struct {
		name           string
		userID         string
		mockSetup      func(mockUseCase *mockusecase.MockNotificationUseCase)
		expectedStatus int
		expectedBody   []*handler.NotificationResponse
	}{
		{
			name:   "Success",
			userID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockNotificationUseCase) {
				relatedEntityType := "disaster"
				relatedEntityID := "123"
				readAt := time.Now()

				notifications := []*model.Notification{
					{
						ID:                1,
						UserID:            1,
						Title:             "Test Notification 1",
						Message:           "This is a test notification",
						NotificationType:  "alert",
						RelatedEntityType: &relatedEntityType,
						RelatedEntityID:   &relatedEntityID,
						IsRead:            true,
						ReadAt:            &readAt,
						CreatedAt:         time.Now(),
						UpdatedAt:         time.Now(),
					},
					{
						ID:               2,
						UserID:           1,
						Title:            "Test Notification 2",
						Message:          "This is another test notification",
						NotificationType: "info",
						IsRead:           false,
						CreatedAt:        time.Now(),
						UpdatedAt:        time.Now(),
					},
				}
				mockUseCase.EXPECT().GetNotificationsByUserID(gomock.Any(), int32(1)).Return(notifications, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: []*handler.NotificationResponse{
				{
					ID:                1,
					UserID:            1,
					Title:             "Test Notification 1",
					Message:           "This is a test notification",
					NotificationType:  "alert",
					RelatedEntityType: strPtr("disaster"),
					RelatedEntityID:   strPtr("123"),
					IsRead:            true,
				},
				{
					ID:               2,
					UserID:           1,
					Title:            "Test Notification 2",
					Message:          "This is another test notification",
					NotificationType: "info",
					IsRead:           false,
				},
			},
		},
		{
			name:           "Invalid User ID",
			userID:         "invalid",
			mockSetup:      func(mockUseCase *mockusecase.MockNotificationUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name:   "Error",
			userID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockNotificationUseCase) {
				mockUseCase.EXPECT().GetNotificationsByUserID(gomock.Any(), int32(1)).Return(nil, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   nil,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup(mockUseCase)

			// Make request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/notifications/user/"+tt.userID, nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response []*handler.NotificationResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, len(tt.expectedBody), len(response))

				for i, expected := range tt.expectedBody {
					assert.Equal(t, expected.ID, response[i].ID)
					assert.Equal(t, expected.UserID, response[i].UserID)
					assert.Equal(t, expected.Title, response[i].Title)
					assert.Equal(t, expected.Message, response[i].Message)
					assert.Equal(t, expected.NotificationType, response[i].NotificationType)
					assert.Equal(t, expected.IsRead, response[i].IsRead)

					// Check optional fields only if they exist in the expected response
					if expected.RelatedEntityType != nil {
						assert.Equal(t, *expected.RelatedEntityType, *response[i].RelatedEntityType)
					}
					if expected.RelatedEntityID != nil {
						assert.Equal(t, *expected.RelatedEntityID, *response[i].RelatedEntityID)
					}
				}
			}
		})
	}
}

func TestNotificationHandler_CreateNotification(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupNotificationTest(t)
	r.POST("/notifications", h.CreateNotification)

	// Test cases
	tests := []struct {
		name           string
		requestBody    handler.CreateNotificationRequest
		mockSetup      func(mockUseCase *mockusecase.MockNotificationUseCase)
		expectedStatus int
		expectedBody   *handler.NotificationResponse
	}{
		{
			name: "Success",
			requestBody: handler.CreateNotificationRequest{
				UserID:            1,
				Title:             "New Notification",
				Message:           "This is a new notification",
				NotificationType:  "alert",
				RelatedEntityType: strPtr("disaster"),
				RelatedEntityID:   strPtr("123"),
			},
			mockSetup: func(mockUseCase *mockusecase.MockNotificationUseCase) {
				mockUseCase.EXPECT().
					CreateNotification(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ interface{}, notification *model.Notification) error {
						notification.ID = 1 // Simulate ID generation
						notification.CreatedAt = time.Now()
						notification.UpdatedAt = time.Now()
						return nil
					})
			},
			expectedStatus: http.StatusCreated,
			expectedBody: &handler.NotificationResponse{
				ID:                1,
				UserID:            1,
				Title:             "New Notification",
				Message:           "This is a new notification",
				NotificationType:  "alert",
				RelatedEntityType: strPtr("disaster"),
				RelatedEntityID:   strPtr("123"),
				IsRead:            false,
			},
		},
		{
			name: "Invalid Request - Missing Required Field",
			requestBody: handler.CreateNotificationRequest{
				// Missing Title
				UserID:           1,
				Message:          "This is a new notification",
				NotificationType: "alert",
			},
			mockSetup:      func(mockUseCase *mockusecase.MockNotificationUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name: "Database Error",
			requestBody: handler.CreateNotificationRequest{
				UserID:           1,
				Title:            "New Notification",
				Message:          "This is a new notification",
				NotificationType: "alert",
			},
			mockSetup: func(mockUseCase *mockusecase.MockNotificationUseCase) {
				mockUseCase.EXPECT().
					CreateNotification(gomock.Any(), gomock.Any()).
					Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   nil,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup(mockUseCase)

			// Prepare request body
			jsonBody, _ := json.Marshal(tt.requestBody)

			// Make request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/notifications", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusCreated {
				var response handler.NotificationResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				// Check fields
				assert.Equal(t, tt.requestBody.UserID, response.UserID)
				assert.Equal(t, tt.requestBody.Title, response.Title)
				assert.Equal(t, tt.requestBody.Message, response.Message)
				assert.Equal(t, tt.requestBody.NotificationType, response.NotificationType)
				assert.False(t, response.IsRead)

				// Check optional fields only if they exist in the expected response
				if tt.requestBody.RelatedEntityType != nil {
					assert.Equal(t, *tt.requestBody.RelatedEntityType, *response.RelatedEntityType)
				}
				if tt.requestBody.RelatedEntityID != nil {
					assert.Equal(t, *tt.requestBody.RelatedEntityID, *response.RelatedEntityID)
				}
			}
		})
	}
}

func TestNotificationHandler_UpdateNotification(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupNotificationTest(t)
	r.PUT("/notifications/:id", h.UpdateNotification)

	// Test cases
	tests := []struct {
		name           string
		notificationID string
		requestBody    handler.UpdateNotificationRequest
		mockSetup      func(mockUseCase *mockusecase.MockNotificationUseCase)
		expectedStatus int
		expectedBody   *handler.NotificationResponse
	}{
		{
			name:           "Success",
			notificationID: "1",
			requestBody: handler.UpdateNotificationRequest{
				Title:             "Updated Notification",
				Message:           "This is an updated notification",
				NotificationType:  "info",
				RelatedEntityType: strPtr("disaster"),
				RelatedEntityID:   strPtr("456"),
			},
			mockSetup: func(mockUseCase *mockusecase.MockNotificationUseCase) {
				oldRelatedEntityType := "disaster"
				oldRelatedEntityID := "123"
				readAt := time.Now()

				notification := &model.Notification{
					ID:                1,
					UserID:            1,
					Title:             "Old Notification",
					Message:           "This is an old notification",
					NotificationType:  "alert",
					RelatedEntityType: &oldRelatedEntityType,
					RelatedEntityID:   &oldRelatedEntityID,
					IsRead:            true,
					ReadAt:            &readAt,
					CreatedAt:         time.Now(),
					UpdatedAt:         time.Now(),
				}
				mockUseCase.EXPECT().GetNotificationByID(gomock.Any(), int32(1)).Return(notification, nil)
				mockUseCase.EXPECT().
					UpdateNotification(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ interface{}, n *model.Notification) error {
						assert.Equal(t, "Updated Notification", n.Title)
						assert.Equal(t, "This is an updated notification", n.Message)
						assert.Equal(t, "info", n.NotificationType)
						assert.Equal(t, "disaster", *n.RelatedEntityType)
						assert.Equal(t, "456", *n.RelatedEntityID)
						return nil
					})
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.NotificationResponse{
				ID:                1,
				UserID:            1,
				Title:             "Updated Notification",
				Message:           "This is an updated notification",
				NotificationType:  "info",
				RelatedEntityType: strPtr("disaster"),
				RelatedEntityID:   strPtr("456"),
				IsRead:            true,
			},
		},
		{
			name:           "Invalid ID",
			notificationID: "invalid",
			requestBody: handler.UpdateNotificationRequest{
				Title:            "Updated Notification",
				Message:          "This is an updated notification",
				NotificationType: "info",
			},
			mockSetup:      func(mockUseCase *mockusecase.MockNotificationUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name:           "Not Found",
			notificationID: "999",
			requestBody: handler.UpdateNotificationRequest{
				Title:            "Updated Notification",
				Message:          "This is an updated notification",
				NotificationType: "info",
			},
			mockSetup: func(mockUseCase *mockusecase.MockNotificationUseCase) {
				mockUseCase.EXPECT().GetNotificationByID(gomock.Any(), int32(999)).Return(nil, errors.New("not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   nil,
		},
		{
			name:           "Database Error",
			notificationID: "1",
			requestBody: handler.UpdateNotificationRequest{
				Title:            "Updated Notification",
				Message:          "This is an updated notification",
				NotificationType: "info",
			},
			mockSetup: func(mockUseCase *mockusecase.MockNotificationUseCase) {
				notification := &model.Notification{
					ID:               1,
					UserID:           1,
					Title:            "Old Notification",
					Message:          "This is an old notification",
					NotificationType: "alert",
					IsRead:           false,
					CreatedAt:        time.Now(),
					UpdatedAt:        time.Now(),
				}
				mockUseCase.EXPECT().GetNotificationByID(gomock.Any(), int32(1)).Return(notification, nil)
				mockUseCase.EXPECT().UpdateNotification(gomock.Any(), gomock.Any()).Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   nil,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup(mockUseCase)

			// Prepare request body
			jsonBody, _ := json.Marshal(tt.requestBody)

			// Make request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPut, "/notifications/"+tt.notificationID, bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response handler.NotificationResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.requestBody.Title, response.Title)
				assert.Equal(t, tt.requestBody.Message, response.Message)
				assert.Equal(t, tt.requestBody.NotificationType, response.NotificationType)

				if tt.requestBody.RelatedEntityType != nil {
					assert.Equal(t, *tt.requestBody.RelatedEntityType, *response.RelatedEntityType)
				}
				if tt.requestBody.RelatedEntityID != nil {
					assert.Equal(t, *tt.requestBody.RelatedEntityID, *response.RelatedEntityID)
				}
			}
		})
	}
}

func TestNotificationHandler_DeleteNotification(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupNotificationTest(t)
	r.DELETE("/notifications/:id", h.DeleteNotification)

	// Test cases
	tests := []struct {
		name           string
		notificationID string
		mockSetup      func(mockUseCase *mockusecase.MockNotificationUseCase)
		expectedStatus int
	}{
		{
			name:           "Success",
			notificationID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockNotificationUseCase) {
				notification := &model.Notification{
					ID:               1,
					UserID:           1,
					Title:            "Test Notification",
					Message:          "This is a test notification",
					NotificationType: "alert",
					IsRead:           false,
					CreatedAt:        time.Now(),
					UpdatedAt:        time.Now(),
				}
				mockUseCase.EXPECT().GetNotificationByID(gomock.Any(), int32(1)).Return(notification, nil)
				mockUseCase.EXPECT().DeleteNotification(gomock.Any(), int32(1)).Return(nil)
			},
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "Invalid ID",
			notificationID: "invalid",
			mockSetup:      func(mockUseCase *mockusecase.MockNotificationUseCase) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Not Found",
			notificationID: "999",
			mockSetup: func(mockUseCase *mockusecase.MockNotificationUseCase) {
				mockUseCase.EXPECT().GetNotificationByID(gomock.Any(), int32(999)).Return(nil, errors.New("not found"))
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Database Error",
			notificationID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockNotificationUseCase) {
				notification := &model.Notification{
					ID:               1,
					UserID:           1,
					Title:            "Test Notification",
					Message:          "This is a test notification",
					NotificationType: "alert",
					IsRead:           false,
					CreatedAt:        time.Now(),
					UpdatedAt:        time.Now(),
				}
				mockUseCase.EXPECT().GetNotificationByID(gomock.Any(), int32(1)).Return(notification, nil)
				mockUseCase.EXPECT().DeleteNotification(gomock.Any(), int32(1)).Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup(mockUseCase)

			// Make request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodDelete, "/notifications/"+tt.notificationID, nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func TestNotificationHandler_MarkAsRead(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupNotificationTest(t)
	r.PUT("/notifications/:id/read", h.MarkAsRead)

	// Test cases
	tests := []struct {
		name           string
		notificationID string
		mockSetup      func(mockUseCase *mockusecase.MockNotificationUseCase)
		expectedStatus int
		expectedBody   *handler.NotificationResponse
	}{
		{
			name:           "Success",
			notificationID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockNotificationUseCase) {
				// Initial notification (not read)
				notification := &model.Notification{
					ID:               1,
					UserID:           1,
					Title:            "Test Notification",
					Message:          "This is a test notification",
					NotificationType: "alert",
					IsRead:           false,
					CreatedAt:        time.Now(),
					UpdatedAt:        time.Now(),
				}
				mockUseCase.EXPECT().GetNotificationByID(gomock.Any(), int32(1)).Return(notification, nil)
				mockUseCase.EXPECT().MarkAsRead(gomock.Any(), int32(1)).Return(nil)

				// Updated notification (read)
				readAt := time.Now()
				updatedNotification := &model.Notification{
					ID:               1,
					UserID:           1,
					Title:            "Test Notification",
					Message:          "This is a test notification",
					NotificationType: "alert",
					IsRead:           true,
					ReadAt:           &readAt,
					CreatedAt:        notification.CreatedAt,
					UpdatedAt:        time.Now(),
				}
				mockUseCase.EXPECT().GetNotificationByID(gomock.Any(), int32(1)).Return(updatedNotification, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.NotificationResponse{
				ID:               1,
				UserID:           1,
				Title:            "Test Notification",
				Message:          "This is a test notification",
				NotificationType: "alert",
				IsRead:           true,
			},
		},
		{
			name:           "Already Read",
			notificationID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockNotificationUseCase) {
				readAt := time.Now()
				notification := &model.Notification{
					ID:               1,
					UserID:           1,
					Title:            "Test Notification",
					Message:          "This is a test notification",
					NotificationType: "alert",
					IsRead:           true,
					ReadAt:           &readAt,
					CreatedAt:        time.Now(),
					UpdatedAt:        time.Now(),
				}
				mockUseCase.EXPECT().GetNotificationByID(gomock.Any(), int32(1)).Return(notification, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.NotificationResponse{
				ID:               1,
				UserID:           1,
				Title:            "Test Notification",
				Message:          "This is a test notification",
				NotificationType: "alert",
				IsRead:           true,
			},
		},
		{
			name:           "Invalid ID",
			notificationID: "invalid",
			mockSetup:      func(mockUseCase *mockusecase.MockNotificationUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name:           "Not Found",
			notificationID: "999",
			mockSetup: func(mockUseCase *mockusecase.MockNotificationUseCase) {
				mockUseCase.EXPECT().GetNotificationByID(gomock.Any(), int32(999)).Return(nil, errors.New("not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   nil,
		},
		{
			name:           "Mark As Read Error",
			notificationID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockNotificationUseCase) {
				notification := &model.Notification{
					ID:               1,
					UserID:           1,
					Title:            "Test Notification",
					Message:          "This is a test notification",
					NotificationType: "alert",
					IsRead:           false,
					CreatedAt:        time.Now(),
					UpdatedAt:        time.Now(),
				}
				mockUseCase.EXPECT().GetNotificationByID(gomock.Any(), int32(1)).Return(notification, nil)
				mockUseCase.EXPECT().MarkAsRead(gomock.Any(), int32(1)).Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   nil,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup(mockUseCase)

			// Make request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPut, "/notifications/"+tt.notificationID+"/read", nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response handler.NotificationResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody.ID, response.ID)
				assert.Equal(t, tt.expectedBody.UserID, response.UserID)
				assert.Equal(t, tt.expectedBody.Title, response.Title)
				assert.Equal(t, tt.expectedBody.Message, response.Message)
				assert.Equal(t, tt.expectedBody.NotificationType, response.NotificationType)
				assert.Equal(t, tt.expectedBody.IsRead, response.IsRead)
			}
		})
	}
}

// strPtr is defined in another test file in the same package
