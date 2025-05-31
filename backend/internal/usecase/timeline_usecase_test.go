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

func setupTimelineTest(t *testing.T) (*mockdatastore.MockTimelineRepository, usecase.TimelineUseCase) {
	ctrl := gomock.NewController(t)
	mockRepo := mockdatastore.NewMockTimelineRepository(ctrl)
	useCase := usecase.NewTimelineUseCase(mockRepo)
	return mockRepo, useCase
}

func TestTimelineUseCase_GetTimelinesByDisasterID(t *testing.T) {
	// Setup
	mockRepo, useCase := setupTimelineTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		disasterID    string
		mockSetup     func(mockRepo *mockdatastore.MockTimelineRepository)
		expectedError bool
		expectedLen   int
	}{
		{
			name:       "Success",
			disasterID: "1",
			mockSetup: func(mockRepo *mockdatastore.MockTimelineRepository) {
				now := time.Now()
				severity1 := "高"
				severity2 := "中"

				timelines := []*model.Timeline{
					{
						ID:          1,
						DisasterID:  "1",
						EventName:   "地震発生",
						EventTime:   now.Add(-24 * time.Hour),
						Description: "マグニチュード6.5の地震が発生",
						Severity:    &severity1,
						CreatedAt:   now.Add(-24 * time.Hour),
						UpdatedAt:   now.Add(-24 * time.Hour),
					},
					{
						ID:          2,
						DisasterID:  "1",
						EventName:   "避難指示発令",
						EventTime:   now.Add(-23 * time.Hour),
						Description: "被災地域に避難指示が発令されました",
						Severity:    &severity2,
						CreatedAt:   now.Add(-23 * time.Hour),
						UpdatedAt:   now.Add(-23 * time.Hour),
					},
				}
				mockRepo.EXPECT().FindByDisasterID(gomock.Any(), "1").Return(timelines, nil)
			},
			expectedError: false,
			expectedLen:   2,
		},
		{
			name:       "No Timelines",
			disasterID: "2",
			mockSetup: func(mockRepo *mockdatastore.MockTimelineRepository) {
				mockRepo.EXPECT().FindByDisasterID(gomock.Any(), "2").Return([]*model.Timeline{}, nil)
			},
			expectedError: false,
			expectedLen:   0,
		},
		{
			name:       "Error",
			disasterID: "1",
			mockSetup: func(mockRepo *mockdatastore.MockTimelineRepository) {
				mockRepo.EXPECT().FindByDisasterID(gomock.Any(), "1").Return(nil, errors.New("database error"))
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
			timelines, err := useCase.GetTimelinesByDisasterID(ctx, tt.disasterID)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, timelines)
			} else {
				assert.NoError(t, err)
				if tt.expectedLen > 0 {
					assert.NotNil(t, timelines)
				}
				assert.Equal(t, tt.expectedLen, len(timelines))

				if tt.expectedLen > 0 {
					// Check that all timelines have the correct disaster ID
					for _, timeline := range timelines {
						assert.Equal(t, tt.disasterID, timeline.DisasterID)
					}
				}
			}
		})
	}
}
