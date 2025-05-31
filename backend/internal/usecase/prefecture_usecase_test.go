package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/usecase"
	mockdatastore "github.com/AI1411/fullstack-react-go/tests/mock/datastore"
)

func setupPrefectureTest(t *testing.T) (*mockdatastore.MockPrefectureRepository, usecase.PrefectureUseCase) {
	ctrl := gomock.NewController(t)
	mockRepo := mockdatastore.NewMockPrefectureRepository(ctrl)
	useCase := usecase.NewPrefectureUseCase(mockRepo)
	return mockRepo, useCase
}

func TestPrefectureUseCase_ListPrefectures(t *testing.T) {
	// Setup
	mockRepo, useCase := setupPrefectureTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		mockSetup     func(mockRepo *mockdatastore.MockPrefectureRepository)
		expectedError bool
		expectedLen   int
	}{
		{
			name: "Success",
			mockSetup: func(mockRepo *mockdatastore.MockPrefectureRepository) {
				prefectures := []*model.Prefecture{
					{
						ID:       1,
						Name:     "北海道",
						RegionID: 1,
					},
					{
						ID:       13,
						Name:     "東京都",
						RegionID: 3,
					},
					{
						ID:       27,
						Name:     "大阪府",
						RegionID: 5,
					},
				}
				mockRepo.EXPECT().Find(gomock.Any()).Return(prefectures, nil)
			},
			expectedError: false,
			expectedLen:   3,
		},
		{
			name: "Error",
			mockSetup: func(mockRepo *mockdatastore.MockPrefectureRepository) {
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
			prefectures, err := useCase.ListPrefectures(ctx)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, prefectures)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, prefectures)
				assert.Equal(t, tt.expectedLen, len(prefectures))
			}
		})
	}
}

func TestPrefectureUseCase_GetPrefectureByID(t *testing.T) {
	// Setup
	mockRepo, useCase := setupPrefectureTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		id            int32
		mockSetup     func(mockRepo *mockdatastore.MockPrefectureRepository)
		expectedError bool
	}{
		{
			name: "Success",
			id:   13,
			mockSetup: func(mockRepo *mockdatastore.MockPrefectureRepository) {
				prefecture := &model.Prefecture{
					ID:       13,
					Name:     "東京都",
					RegionID: 3,
				}
				mockRepo.EXPECT().FindByID(gomock.Any(), int32(13)).Return(prefecture, nil)
			},
			expectedError: false,
		},
		{
			name: "Not Found",
			id:   999,
			mockSetup: func(mockRepo *mockdatastore.MockPrefectureRepository) {
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
			prefecture, err := useCase.GetPrefectureByID(ctx, tt.id)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, prefecture)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, prefecture)
				assert.Equal(t, tt.id, prefecture.ID)
			}
		})
	}
}
