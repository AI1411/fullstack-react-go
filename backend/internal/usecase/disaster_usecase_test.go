package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/datastore"
	"github.com/AI1411/fullstack-react-go/internal/usecase"
	mockdatastore "github.com/AI1411/fullstack-react-go/tests/mock/datastore"
)

func setupDisasterTest(t *testing.T) (*mockdatastore.MockDisasterRepository, usecase.DisasterUseCase) {
	ctrl := gomock.NewController(t)
	mockRepo := mockdatastore.NewMockDisasterRepository(ctrl)
	useCase := usecase.NewDisasterUseCase(mockRepo)
	return mockRepo, useCase
}

func TestDisasterUseCase_ListDisasters(t *testing.T) {
	// Setup
	mockRepo, useCase := setupDisasterTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		params        *datastore.DisasterSearchParams
		mockSetup     func(mockRepo *mockdatastore.MockDisasterRepository)
		expectedError bool
		expectedLen   int
	}{
		{
			name:   "Success with no params",
			params: nil,
			mockSetup: func(mockRepo *mockdatastore.MockDisasterRepository) {
				now := time.Now()
				disasters := []*model.Disaster{
					{
						ID:           "1",
						DisasterCode: "D2024-001",
						Name:         "東京地震",
						PrefectureID: 13,
						OccurredAt:   now,
						Summary:      "東京で発生した地震",
						DisasterType: "地震",
						Status:       "in_progress",
						ImpactLevel:  "中程度",
						CreatedAt:    now,
						UpdatedAt:    now,
					},
					{
						ID:           "2",
						DisasterCode: "D2024-002",
						Name:         "大阪洪水",
						PrefectureID: 27,
						OccurredAt:   now,
						Summary:      "大阪で発生した洪水",
						DisasterType: "洪水",
						Status:       "completed",
						ImpactLevel:  "深刻",
						CreatedAt:    now,
						UpdatedAt:    now,
					},
				}
				mockRepo.EXPECT().Find(gomock.Any(), nil).Return(disasters, nil)
			},
			expectedError: false,
			expectedLen:   2,
		},
		{
			name: "Success with params",
			params: &datastore.DisasterSearchParams{
				DisasterType: "地震",
				Status:       "in_progress",
			},
			mockSetup: func(mockRepo *mockdatastore.MockDisasterRepository) {
				now := time.Now()
				disasters := []*model.Disaster{
					{
						ID:           "1",
						DisasterCode: "D2024-001",
						Name:         "東京地震",
						PrefectureID: 13,
						OccurredAt:   now,
						Summary:      "東京で発生した地震",
						DisasterType: "地震",
						Status:       "in_progress",
						ImpactLevel:  "中程度",
						CreatedAt:    now,
						UpdatedAt:    now,
					},
				}
				mockRepo.EXPECT().Find(gomock.Any(), gomock.Any()).Return(disasters, nil)
			},
			expectedError: false,
			expectedLen:   1,
		},
		{
			name:   "Error",
			params: nil,
			mockSetup: func(mockRepo *mockdatastore.MockDisasterRepository) {
				mockRepo.EXPECT().Find(gomock.Any(), nil).Return(nil, errors.New("database error"))
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
			disasters, err := useCase.ListDisasters(ctx, tt.params)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, disasters)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, disasters)
				assert.Equal(t, tt.expectedLen, len(disasters))
			}
		})
	}
}

func TestDisasterUseCase_GetDisasterByID(t *testing.T) {
	// Setup
	mockRepo, useCase := setupDisasterTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		id            string
		mockSetup     func(mockRepo *mockdatastore.MockDisasterRepository)
		expectedError bool
	}{
		{
			name: "Success",
			id:   "1",
			mockSetup: func(mockRepo *mockdatastore.MockDisasterRepository) {
				now := time.Now()
				disaster := &model.Disaster{
					ID:           "1",
					DisasterCode: "D2024-001",
					Name:         "東京地震",
					PrefectureID: 13,
					OccurredAt:   now,
					Summary:      "東京で発生した地震",
					DisasterType: "地震",
					Status:       "in_progress",
					ImpactLevel:  "中程度",
					CreatedAt:    now,
					UpdatedAt:    now,
				}
				mockRepo.EXPECT().FindByID(gomock.Any(), "1").Return(disaster, nil)
			},
			expectedError: false,
		},
		{
			name: "Not Found",
			id:   "999",
			mockSetup: func(mockRepo *mockdatastore.MockDisasterRepository) {
				mockRepo.EXPECT().FindByID(gomock.Any(), "999").Return(nil, errors.New("not found"))
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
			disaster, err := useCase.GetDisasterByID(ctx, tt.id)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, disaster)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, disaster)
				assert.Equal(t, tt.id, disaster.ID)
			}
		})
	}
}

func TestDisasterUseCase_CreateDisaster(t *testing.T) {
	// Setup
	mockRepo, useCase := setupDisasterTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		disaster      *model.Disaster
		mockSetup     func(mockRepo *mockdatastore.MockDisasterRepository)
		expectedError bool
	}{
		{
			name: "Success",
			disaster: &model.Disaster{
				DisasterCode: "D2024-003",
				Name:         "福岡台風",
				PrefectureID: 40,
				OccurredAt:   time.Now(),
				Summary:      "福岡で発生した台風",
				DisasterType: "台風",
				Status:       "pending",
				ImpactLevel:  "軽微",
			},
			mockSetup: func(mockRepo *mockdatastore.MockDisasterRepository) {
				mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			disaster: &model.Disaster{
				DisasterCode: "D2024-003",
				Name:         "福岡台風",
				PrefectureID: 40,
				OccurredAt:   time.Now(),
				Summary:      "福岡で発生した台風",
				DisasterType: "台風",
				Status:       "pending",
				ImpactLevel:  "軽微",
			},
			mockSetup: func(mockRepo *mockdatastore.MockDisasterRepository) {
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
			err := useCase.CreateDisaster(ctx, tt.disaster)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDisasterUseCase_UpdateDisaster(t *testing.T) {
	// Setup
	mockRepo, useCase := setupDisasterTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		disaster      *model.Disaster
		mockSetup     func(mockRepo *mockdatastore.MockDisasterRepository)
		expectedError bool
	}{
		{
			name: "Success",
			disaster: &model.Disaster{
				ID:           "1",
				DisasterCode: "D2024-001",
				Name:         "東京地震（更新）",
				PrefectureID: 13,
				OccurredAt:   time.Now(),
				Summary:      "東京で発生した地震（更新）",
				DisasterType: "地震",
				Status:       "completed",
				ImpactLevel:  "深刻",
			},
			mockSetup: func(mockRepo *mockdatastore.MockDisasterRepository) {
				mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			disaster: &model.Disaster{
				ID:           "1",
				DisasterCode: "D2024-001",
				Name:         "東京地震（更新）",
				PrefectureID: 13,
				OccurredAt:   time.Now(),
				Summary:      "東京で発生した地震（更新）",
				DisasterType: "地震",
				Status:       "completed",
				ImpactLevel:  "深刻",
			},
			mockSetup: func(mockRepo *mockdatastore.MockDisasterRepository) {
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
			err := useCase.UpdateDisaster(ctx, tt.disaster)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDisasterUseCase_DeleteDisaster(t *testing.T) {
	// Setup
	mockRepo, useCase := setupDisasterTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		id            string
		mockSetup     func(mockRepo *mockdatastore.MockDisasterRepository)
		expectedError bool
	}{
		{
			name: "Success",
			id:   "1",
			mockSetup: func(mockRepo *mockdatastore.MockDisasterRepository) {
				mockRepo.EXPECT().Delete(gomock.Any(), "1").Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			id:   "1",
			mockSetup: func(mockRepo *mockdatastore.MockDisasterRepository) {
				mockRepo.EXPECT().Delete(gomock.Any(), "1").Return(errors.New("database error"))
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
			err := useCase.DeleteDisaster(ctx, tt.id)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
