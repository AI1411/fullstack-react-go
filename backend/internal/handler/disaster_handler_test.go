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

func setupDisasterTest(t *testing.T) (*gin.Engine, *mockusecase.MockDisasterUseCase, handler.Disaster) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	ctrl := gomock.NewController(t)
	mockUseCase := mockusecase.NewMockDisasterUseCase(ctrl)
	l := logger.New(logger.DefaultConfig())
	h := handler.NewDisasterHandler(l, mockUseCase)
	return r, mockUseCase, h
}

func TestDisasterHandler_ListDisasters(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupDisasterTest(t)
	r.GET("/disasters", h.ListDisasters)

	// Test cases
	tests := []struct {
		name           string
		queryParams    string
		mockSetup      func(mockUseCase *mockusecase.MockDisasterUseCase)
		expectedStatus int
		expectedBody   *handler.ListDisastersResponse
	}{
		{
			name:        "Success",
			queryParams: "",
			mockSetup: func(mockUseCase *mockusecase.MockDisasterUseCase) {
				occurredAt, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
				disasters := []*model.Disaster{
					{
						ID:           "1",
						DisasterCode: "D2023-001",
						Name:         "Test Disaster 1",
						PrefectureID: 1,
						Prefecture: model.Prefecture{
							ID:   1,
							Name: "東京都",
						},
						OccurredAt:            occurredAt,
						Summary:               "Test Summary 1",
						DisasterType:          "earthquake",
						Status:                "pending",
						ImpactLevel:           "severe",
						AffectedAreaSize:      nil,
						EstimatedDamageAmount: nil,
					},
					{
						ID:           "2",
						DisasterCode: "D2023-002",
						Name:         "Test Disaster 2",
						PrefectureID: 2,
						Prefecture: model.Prefecture{
							ID:   2,
							Name: "大阪府",
						},
						OccurredAt:            occurredAt,
						Summary:               "Test Summary 2",
						DisasterType:          "flood",
						Status:                "in_progress",
						ImpactLevel:           "moderate",
						AffectedAreaSize:      nil,
						EstimatedDamageAmount: nil,
					},
				}
				mockUseCase.EXPECT().ListDisasters(gomock.Any(), gomock.Any()).Return(disasters, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.ListDisastersResponse{
				Disasters: []*handler.DisasterResponse{
					{
						ID:           "1",
						DisasterCode: "D2023-001",
						Name:         "Test Disaster 1",
						Prefecture: handler.PrefectureItem{
							Name: "東京都",
						},
						OccurredAt:            "2023-01-01 00:00:00",
						Summary:               "Test Summary 1",
						DisasterType:          "earthquake",
						Status:                "pending",
						ImpactLevel:           "severe",
						AffectedAreaSize:      nil,
						EstimatedDamageAmount: nil,
					},
					{
						ID:           "2",
						DisasterCode: "D2023-002",
						Name:         "Test Disaster 2",
						Prefecture: handler.PrefectureItem{
							Name: "大阪府",
						},
						OccurredAt:            "2023-01-01 00:00:00",
						Summary:               "Test Summary 2",
						DisasterType:          "flood",
						Status:                "in_progress",
						ImpactLevel:           "moderate",
						AffectedAreaSize:      nil,
						EstimatedDamageAmount: nil,
					},
				},
				Total: 2,
			},
		},
		{
			name:        "With Filters",
			queryParams: "?name=Test&disaster_type=earthquake&status=pending&prefecture_id=1&start_date=2023-01-01T00:00:00Z&end_date=2023-12-31T23:59:59Z",
			mockSetup: func(mockUseCase *mockusecase.MockDisasterUseCase) {
				occurredAt, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
				disasters := []*model.Disaster{
					{
						ID:           "1",
						DisasterCode: "D2023-001",
						Name:         "Test Disaster 1",
						PrefectureID: 1,
						Prefecture: model.Prefecture{
							ID:   1,
							Name: "東京都",
						},
						OccurredAt:            occurredAt,
						Summary:               "Test Summary 1",
						DisasterType:          "earthquake",
						Status:                "pending",
						ImpactLevel:           "severe",
						AffectedAreaSize:      nil,
						EstimatedDamageAmount: nil,
					},
				}
				mockUseCase.EXPECT().ListDisasters(gomock.Any(), gomock.Any()).Return(disasters, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.ListDisastersResponse{
				Disasters: []*handler.DisasterResponse{
					{
						ID:           "1",
						DisasterCode: "D2023-001",
						Name:         "Test Disaster 1",
						Prefecture: handler.PrefectureItem{
							Name: "東京都",
						},
						OccurredAt:            "2023-01-01 00:00:00",
						Summary:               "Test Summary 1",
						DisasterType:          "earthquake",
						Status:                "pending",
						ImpactLevel:           "severe",
						AffectedAreaSize:      nil,
						EstimatedDamageAmount: nil,
					},
				},
				Total: 1,
			},
		},
		{
			name:        "Error",
			queryParams: "",
			mockSetup: func(mockUseCase *mockusecase.MockDisasterUseCase) {
				mockUseCase.EXPECT().ListDisasters(gomock.Any(), gomock.Any()).Return(nil, errors.New("database error"))
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
			req, _ := http.NewRequest(http.MethodGet, "/disasters"+tt.queryParams, nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response handler.ListDisastersResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody.Total, response.Total)
				assert.Equal(t, len(tt.expectedBody.Disasters), len(response.Disasters))

				for i, expected := range tt.expectedBody.Disasters {
					assert.Equal(t, expected.ID, response.Disasters[i].ID)
					assert.Equal(t, expected.DisasterCode, response.Disasters[i].DisasterCode)
					assert.Equal(t, expected.Name, response.Disasters[i].Name)
					assert.Equal(t, expected.Prefecture.Name, response.Disasters[i].Prefecture.Name)
					assert.Equal(t, expected.OccurredAt, response.Disasters[i].OccurredAt)
					assert.Equal(t, expected.Summary, response.Disasters[i].Summary)
					assert.Equal(t, expected.DisasterType, response.Disasters[i].DisasterType)
					assert.Equal(t, expected.Status, response.Disasters[i].Status)
					assert.Equal(t, expected.ImpactLevel, response.Disasters[i].ImpactLevel)
				}
			}
		})
	}
}

func TestDisasterHandler_GetDisaster(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupDisasterTest(t)
	r.GET("/disasters/:id", h.GetDisaster)

	// Test cases
	tests := []struct {
		name           string
		disasterID     string
		mockSetup      func(mockUseCase *mockusecase.MockDisasterUseCase)
		expectedStatus int
		expectedBody   *handler.DisasterResponse
	}{
		{
			name:       "Success",
			disasterID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockDisasterUseCase) {
				occurredAt, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
				disaster := &model.Disaster{
					ID:           "1",
					DisasterCode: "D2023-001",
					Name:         "Test Disaster",
					PrefectureID: 1,
					Prefecture: model.Prefecture{
						ID:   1,
						Name: "東京都",
					},
					OccurredAt:            occurredAt,
					Summary:               "Test Summary",
					DisasterType:          "earthquake",
					Status:                "pending",
					ImpactLevel:           "severe",
					AffectedAreaSize:      nil,
					EstimatedDamageAmount: nil,
				}
				mockUseCase.EXPECT().GetDisasterByID(gomock.Any(), "1").Return(disaster, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.DisasterResponse{
				ID:           "1",
				DisasterCode: "D2023-001",
				Name:         "Test Disaster",
				Prefecture: handler.PrefectureItem{
					Name: "東京都",
				},
				OccurredAt:            "2023-01-01 00:00:00",
				Summary:               "Test Summary",
				DisasterType:          "earthquake",
				Status:                "pending",
				ImpactLevel:           "severe",
				AffectedAreaSize:      nil,
				EstimatedDamageAmount: nil,
			},
		},
		{
			name:       "Not Found",
			disasterID: "999",
			mockSetup: func(mockUseCase *mockusecase.MockDisasterUseCase) {
				mockUseCase.EXPECT().GetDisasterByID(gomock.Any(), "999").Return(nil, errors.New("not found"))
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
			req, _ := http.NewRequest(http.MethodGet, "/disasters/"+tt.disasterID, nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response handler.DisasterResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody.ID, response.ID)
				assert.Equal(t, tt.expectedBody.DisasterCode, response.DisasterCode)
				assert.Equal(t, tt.expectedBody.Name, response.Name)
				assert.Equal(t, tt.expectedBody.Prefecture.Name, response.Prefecture.Name)
				assert.Equal(t, tt.expectedBody.OccurredAt, response.OccurredAt)
				assert.Equal(t, tt.expectedBody.Summary, response.Summary)
				assert.Equal(t, tt.expectedBody.DisasterType, response.DisasterType)
				assert.Equal(t, tt.expectedBody.Status, response.Status)
				assert.Equal(t, tt.expectedBody.ImpactLevel, response.ImpactLevel)
			}
		})
	}
}

func TestDisasterHandler_CreateDisaster(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupDisasterTest(t)
	r.POST("/disasters", h.CreateDisaster)

	// Test cases
	tests := []struct {
		name           string
		requestBody    handler.CreateDisasterRequest
		mockSetup      func(mockUseCase *mockusecase.MockDisasterUseCase)
		expectedStatus int
		expectedBody   *handler.DisasterResponse
	}{
		{
			name: "Success",
			requestBody: handler.CreateDisasterRequest{
				DisasterCode: "D2023-001",
				Name:         "Test Disaster",
				PrefectureID: 1,
				OccurredAt:   "2023-01-01T00:00:00Z",
				Summary:      "Test Summary",
				DisasterType: "earthquake",
				Status:       "pending",
				ImpactLevel:  "severe",
			},
			mockSetup: func(mockUseCase *mockusecase.MockDisasterUseCase) {
				mockUseCase.EXPECT().
					CreateDisaster(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ interface{}, disaster *model.Disaster) error {
						disaster.ID = "1" // Simulate ID generation
						return nil
					})
			},
			expectedStatus: http.StatusCreated,
			expectedBody: &handler.DisasterResponse{
				ID:           "1",
				DisasterCode: "D2023-001",
				Name:         "Test Disaster",
				OccurredAt:   "2023-01-01 00:00:00",
				Summary:      "Test Summary",
				DisasterType: "earthquake",
				Status:       "pending",
				ImpactLevel:  "severe",
			},
		},
		{
			name: "Invalid Request - Missing Required Field",
			requestBody: handler.CreateDisasterRequest{
				// Missing DisasterCode
				Name:         "Test Disaster",
				PrefectureID: 1,
				OccurredAt:   "2023-01-01T00:00:00Z",
				Summary:      "Test Summary",
				DisasterType: "earthquake",
				ImpactLevel:  "severe",
			},
			mockSetup:      func(mockUseCase *mockusecase.MockDisasterUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name: "Invalid Request - Invalid Date Format",
			requestBody: handler.CreateDisasterRequest{
				DisasterCode: "D2023-001",
				Name:         "Test Disaster",
				PrefectureID: 1,
				OccurredAt:   "2023-01-01", // Invalid format
				Summary:      "Test Summary",
				DisasterType: "earthquake",
				ImpactLevel:  "severe",
			},
			mockSetup:      func(mockUseCase *mockusecase.MockDisasterUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name: "Database Error",
			requestBody: handler.CreateDisasterRequest{
				DisasterCode: "D2023-001",
				Name:         "Test Disaster",
				PrefectureID: 1,
				OccurredAt:   "2023-01-01T00:00:00Z",
				Summary:      "Test Summary",
				DisasterType: "earthquake",
				Status:       "pending",
				ImpactLevel:  "severe",
			},
			mockSetup: func(mockUseCase *mockusecase.MockDisasterUseCase) {
				mockUseCase.EXPECT().
					CreateDisaster(gomock.Any(), gomock.Any()).
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
			req, _ := http.NewRequest(http.MethodPost, "/disasters", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusCreated {
				var response handler.DisasterResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				// Check only fields we can predict
				assert.Equal(t, tt.expectedBody.DisasterCode, response.DisasterCode)
				assert.Equal(t, tt.expectedBody.Name, response.Name)
				assert.Equal(t, tt.expectedBody.Summary, response.Summary)
				assert.Equal(t, tt.expectedBody.DisasterType, response.DisasterType)
				assert.Equal(t, tt.expectedBody.Status, response.Status)
				assert.Equal(t, tt.expectedBody.ImpactLevel, response.ImpactLevel)
			}
		})
	}
}

func TestDisasterHandler_UpdateDisaster(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupDisasterTest(t)
	r.PUT("/disasters/:id", h.UpdateDisaster)

	// Test cases
	tests := []struct {
		name           string
		disasterID     string
		requestBody    handler.UpdateDisasterRequest
		mockSetup      func(mockUseCase *mockusecase.MockDisasterUseCase)
		expectedStatus int
		expectedBody   *handler.DisasterResponse
	}{
		{
			name:       "Success",
			disasterID: "1",
			requestBody: handler.UpdateDisasterRequest{
				Name:         "Updated Disaster",
				DisasterType: "flood",
				Status:       "in_progress",
			},
			mockSetup: func(mockUseCase *mockusecase.MockDisasterUseCase) {
				occurredAt, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
				disaster := &model.Disaster{
					ID:           "1",
					DisasterCode: "D2023-001",
					Name:         "Test Disaster",
					PrefectureID: 1,
					Prefecture: model.Prefecture{
						ID:   1,
						Name: "東京都",
					},
					OccurredAt:            occurredAt,
					Summary:               "Test Summary",
					DisasterType:          "earthquake",
					Status:                "pending",
					ImpactLevel:           "severe",
					AffectedAreaSize:      nil,
					EstimatedDamageAmount: nil,
				}
				mockUseCase.EXPECT().GetDisasterByID(gomock.Any(), "1").Return(disaster, nil)
				mockUseCase.EXPECT().
					UpdateDisaster(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ interface{}, d *model.Disaster) error {
						assert.Equal(t, "Updated Disaster", d.Name)
						assert.Equal(t, "flood", d.DisasterType)
						assert.Equal(t, "in_progress", d.Status)
						return nil
					})
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.DisasterResponse{
				ID:           "1",
				DisasterCode: "D2023-001",
				Name:         "Updated Disaster",
				Prefecture: handler.PrefectureItem{
					Name: "東京都",
				},
				OccurredAt:            "2023-01-01 00:00:00",
				Summary:               "Test Summary",
				DisasterType:          "flood",
				Status:                "in_progress",
				ImpactLevel:           "severe",
				AffectedAreaSize:      nil,
				EstimatedDamageAmount: nil,
			},
		},
		{
			name:       "Not Found",
			disasterID: "999",
			requestBody: handler.UpdateDisasterRequest{
				Name: "Updated Disaster",
			},
			mockSetup: func(mockUseCase *mockusecase.MockDisasterUseCase) {
				mockUseCase.EXPECT().GetDisasterByID(gomock.Any(), "999").Return(nil, errors.New("not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   nil,
		},
		{
			name:       "Invalid Date Format",
			disasterID: "1",
			requestBody: handler.UpdateDisasterRequest{
				OccurredAt: "2023-01-01", // Invalid format
			},
			mockSetup: func(mockUseCase *mockusecase.MockDisasterUseCase) {
				occurredAt, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
				disaster := &model.Disaster{
					ID:           "1",
					DisasterCode: "D2023-001",
					Name:         "Test Disaster",
					PrefectureID: 1,
					OccurredAt:   occurredAt,
					Summary:      "Test Summary",
					DisasterType: "earthquake",
					Status:       "pending",
					ImpactLevel:  "severe",
				}
				mockUseCase.EXPECT().GetDisasterByID(gomock.Any(), "1").Return(disaster, nil)
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name:       "Database Error",
			disasterID: "1",
			requestBody: handler.UpdateDisasterRequest{
				Name: "Updated Disaster",
			},
			mockSetup: func(mockUseCase *mockusecase.MockDisasterUseCase) {
				occurredAt, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
				disaster := &model.Disaster{
					ID:           "1",
					DisasterCode: "D2023-001",
					Name:         "Test Disaster",
					PrefectureID: 1,
					OccurredAt:   occurredAt,
					Summary:      "Test Summary",
					DisasterType: "earthquake",
					Status:       "pending",
					ImpactLevel:  "severe",
				}
				mockUseCase.EXPECT().GetDisasterByID(gomock.Any(), "1").Return(disaster, nil)
				mockUseCase.EXPECT().UpdateDisaster(gomock.Any(), gomock.Any()).Return(errors.New("database error"))
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
			req, _ := http.NewRequest(http.MethodPut, "/disasters/"+tt.disasterID, bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response handler.DisasterResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				if tt.requestBody.Name != "" {
					assert.Equal(t, tt.requestBody.Name, response.Name)
				}
				if tt.requestBody.DisasterType != "" {
					assert.Equal(t, tt.requestBody.DisasterType, response.DisasterType)
				}
				if tt.requestBody.Status != "" {
					assert.Equal(t, tt.requestBody.Status, response.Status)
				}
			}
		})
	}
}

func TestDisasterHandler_DeleteDisaster(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupDisasterTest(t)
	r.DELETE("/disasters/:id", h.DeleteDisaster)

	// Test cases
	tests := []struct {
		name           string
		disasterID     string
		mockSetup      func(mockUseCase *mockusecase.MockDisasterUseCase)
		expectedStatus int
	}{
		{
			name:       "Success",
			disasterID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockDisasterUseCase) {
				occurredAt, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
				disaster := &model.Disaster{
					ID:           "1",
					DisasterCode: "D2023-001",
					Name:         "Test Disaster",
					PrefectureID: 1,
					OccurredAt:   occurredAt,
					Summary:      "Test Summary",
					DisasterType: "earthquake",
					Status:       "pending",
					ImpactLevel:  "severe",
				}
				mockUseCase.EXPECT().GetDisasterByID(gomock.Any(), "1").Return(disaster, nil)
				mockUseCase.EXPECT().DeleteDisaster(gomock.Any(), "1").Return(nil)
			},
			expectedStatus: http.StatusNoContent,
		},
		{
			name:       "Not Found",
			disasterID: "999",
			mockSetup: func(mockUseCase *mockusecase.MockDisasterUseCase) {
				mockUseCase.EXPECT().GetDisasterByID(gomock.Any(), "999").Return(nil, errors.New("not found"))
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name:       "Database Error",
			disasterID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockDisasterUseCase) {
				occurredAt, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
				disaster := &model.Disaster{
					ID:           "1",
					DisasterCode: "D2023-001",
					Name:         "Test Disaster",
					PrefectureID: 1,
					OccurredAt:   occurredAt,
					Summary:      "Test Summary",
					DisasterType: "earthquake",
					Status:       "pending",
					ImpactLevel:  "severe",
				}
				mockUseCase.EXPECT().GetDisasterByID(gomock.Any(), "1").Return(disaster, nil)
				mockUseCase.EXPECT().DeleteDisaster(gomock.Any(), "1").Return(errors.New("database error"))
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
			req, _ := http.NewRequest(http.MethodDelete, "/disasters/"+tt.disasterID, nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
