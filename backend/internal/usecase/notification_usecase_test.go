package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/usecase"
	mockdatastore "github.com/AI1411/fullstack-react-go/tests/mock/datastore"
)

func setupNotificationTest(t *testing.T) (*mockdatastore.MockNotificationRepository, usecase.NotificationUseCase) {
	ctrl := gomock.NewController(t)
	mockRepo := mockdatastore.NewMockNotificationRepository(ctrl)
	useCase := usecase.NewNotificationUseCase(mockRepo)
	return mockRepo, useCase
}

func TestNotificationUseCase_ListNotifications(t *testing.T) {
	// Setup
	mockRepo, useCase := setupNotificationTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		mockSetup     func(mockRepo *mockdatastore.MockNotificationRepository)
		expectedError bool
		expectedLen   int
	}{
		{
			name: "Success",
			mockSetup: func(mockRepo *mockdatastore.MockNotificationRepository) {
				now := time.Now()
				relatedEntityType1 := "disaster"
				relatedEntityID1 := "1"
				readAt1 := now.Add(-1 * time.Hour)

				relatedEntityType2 := "support_application"
				relatedEntityID2 := "2"

				notifications := []*model.Notification{
					{
						ID:                1,
						UserID:            1,
						Title:             "災害情報更新",
						Message:           "東京地震の状況が更新されました。",
						NotificationType:  "disaster_update",
						RelatedEntityType: &relatedEntityType1,
						RelatedEntityID:   &relatedEntityID1,
						IsRead:            true,
						ReadAt:            &readAt1,
						CreatedAt:         now.Add(-2 * time.Hour),
						UpdatedAt:         now.Add(-1 * time.Hour),
					},
					{
						ID:                2,
						UserID:            1,
						Title:             "支援申請受付",
						Message:           "あなたの支援申請が受け付けられました。",
						NotificationType:  "application_received",
						RelatedEntityType: &relatedEntityType2,
						RelatedEntityID:   &relatedEntityID2,
						IsRead:            false,
						ReadAt:            nil,
						CreatedAt:         now.Add(-30 * time.Minute),
						UpdatedAt:         now.Add(-30 * time.Minute),
					},
				}
				mockRepo.EXPECT().Find(gomock.Any()).Return(notifications, nil)
			},
			expectedError: false,
			expectedLen:   2,
		},
		{
			name: "Error",
			mockSetup: func(mockRepo *mockdatastore.MockNotificationRepository) {
				mockRepo.EXPECT().Find(gomock.Any()).Return(nil, errors.New("database error"))
			},
			expectedError: true,
			expectedLen:   0,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup(mockRepo)

			// Call the method
			notifications, err := useCase.ListNotifications(ctx)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, notifications)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, notifications)
				assert.Equal(t, tt.expectedLen, len(notifications))
			}
		})
	}
}

func TestNotificationUseCase_GetNotificationByID(t *testing.T) {
	// Setup
	mockRepo, useCase := setupNotificationTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		id            int32
		mockSetup     func(mockRepo *mockdatastore.MockNotificationRepository)
		expectedError bool
	}{
		{
			name: "Success",
			id:   1,
			mockSetup: func(mockRepo *mockdatastore.MockNotificationRepository) {
				now := time.Now()
				relatedEntityType := "disaster"
				relatedEntityID := "1"
				readAt := now.Add(-1 * time.Hour)

				notification := &model.Notification{
					ID:                1,
					UserID:            1,
					Title:             "災害情報更新",
					Message:           "東京地震の状況が更新されました。",
					NotificationType:  "disaster_update",
					RelatedEntityType: &relatedEntityType,
					RelatedEntityID:   &relatedEntityID,
					IsRead:            true,
					ReadAt:            &readAt,
					CreatedAt:         now.Add(-2 * time.Hour),
					UpdatedAt:         now.Add(-1 * time.Hour),
				}
				mockRepo.EXPECT().FindByID(gomock.Any(), int32(1)).Return(notification, nil)
			},
			expectedError: false,
		},
		{
			name: "Not Found",
			id:   999,
			mockSetup: func(mockRepo *mockdatastore.MockNotificationRepository) {
				mockRepo.EXPECT().FindByID(gomock.Any(), int32(999)).Return(nil, errors.New("not found"))
			},
			expectedError: true,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup(mockRepo)

			// Call the method
			notification, err := useCase.GetNotificationByID(ctx, tt.id)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, notification)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, notification)
				assert.Equal(t, tt.id, notification.ID)
			}
		})
	}
}

func TestNotificationUseCase_GetNotificationsByUserID(t *testing.T) {
	// Setup
	mockRepo, useCase := setupNotificationTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		userID        int32
		mockSetup     func(mockRepo *mockdatastore.MockNotificationRepository)
		expectedError bool
		expectedLen   int
	}{
		{
			name:   "Success",
			userID: 1,
			mockSetup: func(mockRepo *mockdatastore.MockNotificationRepository) {
				now := time.Now()
				relatedEntityType1 := "disaster"
				relatedEntityID1 := "1"
				readAt1 := now.Add(-1 * time.Hour)

				relatedEntityType2 := "support_application"
				relatedEntityID2 := "2"

				notifications := []*model.Notification{
					{
						ID:                1,
						UserID:            1,
						Title:             "災害情報更新",
						Message:           "東京地震の状況が更新されました。",
						NotificationType:  "disaster_update",
						RelatedEntityType: &relatedEntityType1,
						RelatedEntityID:   &relatedEntityID1,
						IsRead:            true,
						ReadAt:            &readAt1,
						CreatedAt:         now.Add(-2 * time.Hour),
						UpdatedAt:         now.Add(-1 * time.Hour),
					},
					{
						ID:                2,
						UserID:            1,
						Title:             "支援申請受付",
						Message:           "あなたの支援申請が受け付けられました。",
						NotificationType:  "application_received",
						RelatedEntityType: &relatedEntityType2,
						RelatedEntityID:   &relatedEntityID2,
						IsRead:            false,
						ReadAt:            nil,
						CreatedAt:         now.Add(-30 * time.Minute),
						UpdatedAt:         now.Add(-30 * time.Minute),
					},
				}
				mockRepo.EXPECT().FindByUserID(gomock.Any(), int32(1)).Return(notifications, nil)
			},
			expectedError: false,
			expectedLen:   2,
		},
		{
			name:   "No Notifications",
			userID: 2,
			mockSetup: func(mockRepo *mockdatastore.MockNotificationRepository) {
				mockRepo.EXPECT().FindByUserID(gomock.Any(), int32(2)).Return([]*model.Notification{}, nil)
			},
			expectedError: false,
			expectedLen:   0,
		},
		{
			name:   "Error",
			userID: 1,
			mockSetup: func(mockRepo *mockdatastore.MockNotificationRepository) {
				mockRepo.EXPECT().FindByUserID(gomock.Any(), int32(1)).Return(nil, errors.New("database error"))
			},
			expectedError: true,
			expectedLen:   0,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup(mockRepo)

			// Call the method
			notifications, err := useCase.GetNotificationsByUserID(ctx, tt.userID)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, notifications)
			} else {
				assert.NoError(t, err)
				if tt.expectedLen > 0 {
					assert.NotNil(t, notifications)
				}
				assert.Equal(t, tt.expectedLen, len(notifications))
			}
		})
	}
}

func TestNotificationUseCase_CreateNotification(t *testing.T) {
	// Setup
	mockRepo, useCase := setupNotificationTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		notification  *model.Notification
		mockSetup     func(mockRepo *mockdatastore.MockNotificationRepository)
		expectedError bool
	}{
		{
			name: "Success",
			notification: func() *model.Notification {
				relatedEntityType := "facility_equipment"
				relatedEntityID := "3"

				return &model.Notification{
					UserID:            2,
					Title:             "設備点検通知",
					Message:           "設備の定期点検が予定されています。",
					NotificationType:  "maintenance_scheduled",
					RelatedEntityType: &relatedEntityType,
					RelatedEntityID:   &relatedEntityID,
					IsRead:            false,
				}
			}(),
			mockSetup: func(mockRepo *mockdatastore.MockNotificationRepository) {
				mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			notification: func() *model.Notification {
				return &model.Notification{
					UserID:           2,
					Title:            "設備点検通知",
					Message:          "設備の定期点検が予定されています。",
					NotificationType: "maintenance_scheduled",
					IsRead:           false,
				}
			}(),
			mockSetup: func(mockRepo *mockdatastore.MockNotificationRepository) {
				mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(errors.New("database error"))
			},
			expectedError: true,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup(mockRepo)

			// Call the method
			err := useCase.CreateNotification(ctx, tt.notification)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNotificationUseCase_UpdateNotification(t *testing.T) {
	// Setup
	mockRepo, useCase := setupNotificationTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		notification  *model.Notification
		mockSetup     func(mockRepo *mockdatastore.MockNotificationRepository)
		expectedError bool
	}{
		{
			name: "Success",
			notification: func() *model.Notification {
				relatedEntityType := "disaster"
				relatedEntityID := "1"

				return &model.Notification{
					ID:                1,
					UserID:            1,
					Title:             "災害情報更新（修正）",
					Message:           "東京地震の状況が更新されました。詳細をご確認ください。",
					NotificationType:  "disaster_update",
					RelatedEntityType: &relatedEntityType,
					RelatedEntityID:   &relatedEntityID,
					IsRead:            false,
				}
			}(),
			mockSetup: func(mockRepo *mockdatastore.MockNotificationRepository) {
				mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			notification: func() *model.Notification {
				return &model.Notification{
					ID:               1,
					UserID:           1,
					Title:            "災害情報更新（修正）",
					Message:          "東京地震の状況が更新されました。詳細をご確認ください。",
					NotificationType: "disaster_update",
					IsRead:           false,
				}
			}(),
			mockSetup: func(mockRepo *mockdatastore.MockNotificationRepository) {
				mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.New("database error"))
			},
			expectedError: true,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup(mockRepo)

			// Call the method
			err := useCase.UpdateNotification(ctx, tt.notification)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNotificationUseCase_DeleteNotification(t *testing.T) {
	// Setup
	mockRepo, useCase := setupNotificationTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		id            int32
		mockSetup     func(mockRepo *mockdatastore.MockNotificationRepository)
		expectedError bool
	}{
		{
			name: "Success",
			id:   1,
			mockSetup: func(mockRepo *mockdatastore.MockNotificationRepository) {
				mockRepo.EXPECT().Delete(gomock.Any(), int32(1)).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			id:   1,
			mockSetup: func(mockRepo *mockdatastore.MockNotificationRepository) {
				mockRepo.EXPECT().Delete(gomock.Any(), int32(1)).Return(errors.New("database error"))
			},
			expectedError: true,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup(mockRepo)

			// Call the method
			err := useCase.DeleteNotification(ctx, tt.id)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNotificationUseCase_MarkAsRead(t *testing.T) {
	// Setup
	mockRepo, useCase := setupNotificationTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		id            int32
		mockSetup     func(mockRepo *mockdatastore.MockNotificationRepository)
		expectedError bool
	}{
		{
			name: "Success",
			id:   2,
			mockSetup: func(mockRepo *mockdatastore.MockNotificationRepository) {
				mockRepo.EXPECT().MarkAsRead(gomock.Any(), int32(2)).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			id:   2,
			mockSetup: func(mockRepo *mockdatastore.MockNotificationRepository) {
				mockRepo.EXPECT().MarkAsRead(gomock.Any(), int32(2)).Return(errors.New("database error"))
			},
			expectedError: true,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup(mockRepo)

			// Call the method
			err := useCase.MarkAsRead(ctx, tt.id)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
