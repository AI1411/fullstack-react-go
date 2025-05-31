package handler_test

import (
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

func setupTimelineTest(t *testing.T) (*gin.Engine, *mockusecase.MockTimelineUseCase, handler.Timeline) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	ctrl := gomock.NewController(t)
	mockUseCase := mockusecase.NewMockTimelineUseCase(ctrl)
	l := logger.New(logger.DefaultConfig())
	h := handler.NewTimelineHandler(l, mockUseCase)
	return r, mockUseCase, h
}

func TestTimelineHandler_GetTimelinesByDisasterID(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupTimelineTest(t)
	r.GET("/disasters/:id/timelines", h.GetTimelinesByDisasterID)

	// Test cases
	tests := []struct {
		name           string
		disasterID     string
		mockSetup      func(mockUseCase *mockusecase.MockTimelineUseCase)
		expectedStatus int
		expectedBody   *handler.ListTimelinesResponse
	}{
		{
			name:       "Success",
			disasterID: "disaster-001",
			mockSetup: func(mockUseCase *mockusecase.MockTimelineUseCase) {
				severity1 := "high"
				severity2 := "medium"

				timelines := []*model.Timeline{
					{
						ID:          1,
						DisasterID:  "disaster-001",
						EventName:   "Initial Impact",
						EventTime:   time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
						Description: "Disaster first occurred",
						Severity:    &severity1,
					},
					{
						ID:          2,
						DisasterID:  "disaster-001",
						EventName:   "Emergency Response",
						EventTime:   time.Date(2023, 1, 1, 14, 0, 0, 0, time.UTC),
						Description: "Emergency teams deployed",
						Severity:    &severity2,
					},
				}
				mockUseCase.EXPECT().GetTimelinesByDisasterID(gomock.Any(), "disaster-001").Return(timelines, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.ListTimelinesResponse{
				Timelines: []*handler.TimelineResponse{
					{
						ID:          1,
						DisasterID:  "disaster-001",
						EventName:   "Initial Impact",
						EventTime:   "2023-01-01 12:00:00",
						Description: "Disaster first occurred",
						Severity:    "high",
					},
					{
						ID:          2,
						DisasterID:  "disaster-001",
						EventName:   "Emergency Response",
						EventTime:   "2023-01-01 14:00:00",
						Description: "Emergency teams deployed",
						Severity:    "medium",
					},
				},
				Total: 2,
			},
		},
		{
			name:       "Empty Result",
			disasterID: "disaster-002",
			mockSetup: func(mockUseCase *mockusecase.MockTimelineUseCase) {
				mockUseCase.EXPECT().GetTimelinesByDisasterID(gomock.Any(), "disaster-002").Return([]*model.Timeline{}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.ListTimelinesResponse{
				Timelines: []*handler.TimelineResponse{},
			},
		},
		{
			name:           "Missing Disaster ID",
			disasterID:     "",
			mockSetup:      func(mockUseCase *mockusecase.MockTimelineUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name:       "Database Error",
			disasterID: "disaster-001",
			mockSetup: func(mockUseCase *mockusecase.MockTimelineUseCase) {
				mockUseCase.EXPECT().GetTimelinesByDisasterID(gomock.Any(), "disaster-001").Return(nil, errors.New("database error"))
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
			var req *http.Request
			if tt.disasterID == "" {
				req, _ = http.NewRequest(http.MethodGet, "/disasters//timelines", nil)
			} else {
				req, _ = http.NewRequest(http.MethodGet, "/disasters/"+tt.disasterID+"/timelines", nil)
			}
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response handler.ListTimelinesResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				if tt.expectedBody.Total > 0 {
					assert.Equal(t, tt.expectedBody.Total, response.Total)
				}
				assert.Equal(t, len(tt.expectedBody.Timelines), len(response.Timelines))

				for i, expected := range tt.expectedBody.Timelines {
					assert.Equal(t, expected.ID, response.Timelines[i].ID)
					assert.Equal(t, expected.DisasterID, response.Timelines[i].DisasterID)
					assert.Equal(t, expected.EventName, response.Timelines[i].EventName)
					assert.Equal(t, expected.EventTime, response.Timelines[i].EventTime)
					assert.Equal(t, expected.Description, response.Timelines[i].Description)
					assert.Equal(t, expected.Severity, response.Timelines[i].Severity)
				}
			}
		})
	}
}
