package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/handler"
	"github.com/AI1411/fullstack-react-go/internal/infra/logger"
	mockusecase "github.com/AI1411/fullstack-react-go/tests/mock/usecase"
)

func setupPrefectureTest(t *testing.T) (*gin.Engine, *mockusecase.MockPrefectureUseCase, handler.Prefecture) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	ctrl := gomock.NewController(t)
	mockUseCase := mockusecase.NewMockPrefectureUseCase(ctrl)
	l := logger.New(logger.DefaultConfig())
	h := handler.NewPrefectureHandler(l, mockUseCase)
	return r, mockUseCase, h
}

func TestPrefectureHandler_ListPrefectures(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupPrefectureTest(t)
	r.GET("/prefectures", h.ListPrefectures)

	// Test cases
	tests := []struct {
		name           string
		mockSetup      func(mockUseCase *mockusecase.MockPrefectureUseCase)
		expectedStatus int
		expectedBody   []*handler.PrefectureResponse
	}{
		{
			name: "Success",
			mockSetup: func(mockUseCase *mockusecase.MockPrefectureUseCase) {
				prefectures := []*model.Prefecture{
					{
						ID:   1,
						Name: "東京都",
					},
					{
						ID:   2,
						Name: "大阪府",
					},
				}
				mockUseCase.EXPECT().ListPrefectures(gomock.Any()).Return(prefectures, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: []*handler.PrefectureResponse{
				{
					ID:   1,
					Name: "東京都",
				},
				{
					ID:   2,
					Name: "大阪府",
				},
			},
		},
		{
			name: "Error",
			mockSetup: func(mockUseCase *mockusecase.MockPrefectureUseCase) {
				mockUseCase.EXPECT().ListPrefectures(gomock.Any()).Return(nil, errors.New("database error"))
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
			req, _ := http.NewRequest(http.MethodGet, "/prefectures", nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response []*handler.PrefectureResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, response)
			}
		})
	}
}

func TestPrefectureHandler_GetPrefecture(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupPrefectureTest(t)
	r.GET("/prefectures/:code", h.GetPrefecture)

	// Test cases
	tests := []struct {
		name           string
		prefectureID   string
		mockSetup      func(mockUseCase *mockusecase.MockPrefectureUseCase)
		expectedStatus int
		expectedBody   *handler.PrefectureResponse
	}{
		{
			name:         "Success",
			prefectureID: "01",
			mockSetup: func(mockUseCase *mockusecase.MockPrefectureUseCase) {
				prefecture := &model.Prefecture{
					ID:   1,
					Name: "東京都",
				}
				mockUseCase.EXPECT().GetPrefectureByID(gomock.Any(), "01").Return(prefecture, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.PrefectureResponse{
				ID:   1,
				Name: "東京都",
			},
		},
		{
			name:         "Not Found",
			prefectureID: "999",
			mockSetup: func(mockUseCase *mockusecase.MockPrefectureUseCase) {
				mockUseCase.EXPECT().GetPrefectureByID(gomock.Any(), "999").Return(nil, errors.New("not found"))
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
			req, _ := http.NewRequest(http.MethodGet, "/prefectures/"+tt.prefectureID, nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response handler.PrefectureResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, &response)
			}
		})
	}
}
