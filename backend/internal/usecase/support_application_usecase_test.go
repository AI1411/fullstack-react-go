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

func setupSupportApplicationTest(t *testing.T) (*mockdatastore.MockSupportApplicationRepository, usecase.SupportApplicationUseCase) {
	ctrl := gomock.NewController(t)
	mockRepo := mockdatastore.NewMockSupportApplicationRepository(ctrl)
	useCase := usecase.NewSupportApplicationUseCase(mockRepo)
	return mockRepo, useCase
}

func TestSupportApplicationUseCase_ListSupportApplications(t *testing.T) {
	// Setup
	mockRepo, useCase := setupSupportApplicationTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		mockSetup     func(mockRepo *mockdatastore.MockSupportApplicationRepository)
		expectedError bool
		expectedLen   int
	}{
		{
			name: "Success",
			mockSetup: func(mockRepo *mockdatastore.MockSupportApplicationRepository) {
				now := time.Now()
				reviewedAt := now.Add(-2 * time.Hour)
				approvedAt := now.Add(-1 * time.Hour)
				notes1 := "緊急性が高いため優先的に処理"

				notes2 := "追加書類の提出待ち"

				supportApplications := []*model.SupportApplication{
					{
						ApplicationID:   "A001",
						ApplicationDate: now.Add(-24 * time.Hour),
						ApplicantName:   "山田太郎",
						DisasterName:    "東京地震",
						RequestedAmount: 500000,
						Status:          "承認済",
						ReviewedAt:      &reviewedAt,
						ApprovedAt:      &approvedAt,
						CompletedAt:     nil,
						Notes:           &notes1,
						CreatedAt:       now.Add(-24 * time.Hour),
						UpdatedAt:       now.Add(-1 * time.Hour),
					},
					{
						ApplicationID:   "A002",
						ApplicationDate: now.Add(-12 * time.Hour),
						ApplicantName:   "鈴木一郎",
						DisasterName:    "大阪洪水",
						RequestedAmount: 300000,
						Status:          "書類確認中",
						ReviewedAt:      nil,
						ApprovedAt:      nil,
						CompletedAt:     nil,
						Notes:           &notes2,
						CreatedAt:       now.Add(-12 * time.Hour),
						UpdatedAt:       now.Add(-6 * time.Hour),
					},
				}
				mockRepo.EXPECT().Find(gomock.Any()).Return(supportApplications, nil)
			},
			expectedError: false,
			expectedLen:   2,
		},
		{
			name: "Error",
			mockSetup: func(mockRepo *mockdatastore.MockSupportApplicationRepository) {
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
			supportApplications, err := useCase.ListSupportApplications(ctx)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, supportApplications)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, supportApplications)
				assert.Equal(t, tt.expectedLen, len(supportApplications))
			}
		})
	}
}

func TestSupportApplicationUseCase_GetSupportApplicationByID(t *testing.T) {
	// Setup
	mockRepo, useCase := setupSupportApplicationTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		id            string
		mockSetup     func(mockRepo *mockdatastore.MockSupportApplicationRepository)
		expectedError bool
	}{
		{
			name: "Success",
			id:   "A001",
			mockSetup: func(mockRepo *mockdatastore.MockSupportApplicationRepository) {
				now := time.Now()
				reviewedAt := now.Add(-2 * time.Hour)
				approvedAt := now.Add(-1 * time.Hour)
				notes := "緊急性が高いため優先的に処理"

				supportApplication := &model.SupportApplication{
					ApplicationID:   "A001",
					ApplicationDate: now.Add(-24 * time.Hour),
					ApplicantName:   "山田太郎",
					DisasterName:    "東京地震",
					RequestedAmount: 500000,
					Status:          "承認済",
					ReviewedAt:      &reviewedAt,
					ApprovedAt:      &approvedAt,
					CompletedAt:     nil,
					Notes:           &notes,
					CreatedAt:       now.Add(-24 * time.Hour),
					UpdatedAt:       now.Add(-1 * time.Hour),
				}
				mockRepo.EXPECT().FindByID(gomock.Any(), "A001").Return(supportApplication, nil)
			},
			expectedError: false,
		},
		{
			name: "Not Found",
			id:   "A999",
			mockSetup: func(mockRepo *mockdatastore.MockSupportApplicationRepository) {
				mockRepo.EXPECT().FindByID(gomock.Any(), "A999").Return(nil, errors.New("not found"))
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
			supportApplication, err := useCase.GetSupportApplicationByID(ctx, tt.id)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, supportApplication)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, supportApplication)
				assert.Equal(t, tt.id, supportApplication.ApplicationID)
			}
		})
	}
}

func TestSupportApplicationUseCase_CreateSupportApplication(t *testing.T) {
	// Setup
	mockRepo, useCase := setupSupportApplicationTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name               string
		supportApplication *model.SupportApplication
		mockSetup          func(mockRepo *mockdatastore.MockSupportApplicationRepository)
		expectedError      bool
	}{
		{
			name: "Success",
			supportApplication: func() *model.SupportApplication {
				now := time.Now()
				notes := "被害状況の写真添付済み"

				return &model.SupportApplication{
					ApplicationID:   "A003",
					ApplicationDate: now,
					ApplicantName:   "佐藤花子",
					DisasterName:    "福岡台風",
					RequestedAmount: 450000,
					Status:          "審査中",
					Notes:           &notes,
				}
			}(),
			mockSetup: func(mockRepo *mockdatastore.MockSupportApplicationRepository) {
				mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			supportApplication: func() *model.SupportApplication {
				now := time.Now()

				return &model.SupportApplication{
					ApplicationID:   "A003",
					ApplicationDate: now,
					ApplicantName:   "佐藤花子",
					DisasterName:    "福岡台風",
					RequestedAmount: 450000,
					Status:          "審査中",
				}
			}(),
			mockSetup: func(mockRepo *mockdatastore.MockSupportApplicationRepository) {
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
			err := useCase.CreateSupportApplication(ctx, tt.supportApplication)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
