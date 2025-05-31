package handler_test

import (
	"bytes"
	"context"
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

func setupTest(t *testing.T) (*gin.Engine, *mockusecase.MockDamageLevelUseCase, handler.DamageLevel) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	ctrl := gomock.NewController(t)
	mockUseCase := mockusecase.NewMockDamageLevelUseCase(ctrl)
	l := logger.New(logger.DefaultConfig())
	h := handler.NewDamageLevelHandler(l, mockUseCase)
	return r, mockUseCase, h
}

func TestDamageLevelHandler_ListDamageLevels(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupTest(t)
	r.GET("/damage-levels", h.ListDamageLevels)

	// Test cases
	tests := []struct {
		name           string
		mockSetup      func(mockUseCase *mockusecase.MockDamageLevelUseCase)
		expectedStatus int
		expectedBody   []*handler.DamageLevelResponse
	}{
		{
			name: "Success",
			mockSetup: func(mockUseCase *mockusecase.MockDamageLevelUseCase) {
				description := "Test description"
				damageLevels := []*model.DamageLevel{
					{
						ID:          1,
						Name:        "軽微",
						Description: &description,
					},
					{
						ID:          2,
						Name:        "中程度",
						Description: nil,
					},
				}
				mockUseCase.EXPECT().ListDamageLevels(gomock.Any()).Return(damageLevels, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: []*handler.DamageLevelResponse{
				{
					ID:          1,
					Name:        "軽微",
					Description: strPtr("Test description"),
				},
				{
					ID:          2,
					Name:        "中程度",
					Description: nil,
				},
			},
		},
		{
			name: "Error",
			mockSetup: func(mockUseCase *mockusecase.MockDamageLevelUseCase) {
				mockUseCase.EXPECT().ListDamageLevels(gomock.Any()).Return(nil, errors.New("database error"))
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
			req, _ := http.NewRequest(http.MethodGet, "/damage-levels", nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response []*handler.DamageLevelResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, response)
			}
		})
	}
}

func TestDamageLevelHandler_GetDamageLevel(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupTest(t)
	r.GET("/damage-levels/:id", h.GetDamageLevel)

	// Test cases
	tests := []struct {
		name           string
		damageLevelID  string
		mockSetup      func(mockUseCase *mockusecase.MockDamageLevelUseCase)
		expectedStatus int
		expectedBody   *handler.DamageLevelResponse
	}{
		{
			name:          "Success",
			damageLevelID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockDamageLevelUseCase) {
				description := "Test description"
				damageLevel := &model.DamageLevel{
					ID:          1,
					Name:        "軽微",
					Description: &description,
				}
				mockUseCase.EXPECT().GetDamageLevelByID(gomock.Any(), int32(1)).Return(damageLevel, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.DamageLevelResponse{
				ID:          1,
				Name:        "軽微",
				Description: strPtr("Test description"),
			},
		},
		{
			name:           "Invalid ID",
			damageLevelID:  "invalid",
			mockSetup:      func(mockUseCase *mockusecase.MockDamageLevelUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name:          "Not Found",
			damageLevelID: "999",
			mockSetup: func(mockUseCase *mockusecase.MockDamageLevelUseCase) {
				mockUseCase.EXPECT().GetDamageLevelByID(gomock.Any(), int32(999)).Return(nil, errors.New("not found"))
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
			req, _ := http.NewRequest(http.MethodGet, "/damage-levels/"+tt.damageLevelID, nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response handler.DamageLevelResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, &response)
			}
		})
	}
}

func TestDamageLevelHandler_CreateDamageLevel(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupTest(t)
	r.POST("/damage-levels", h.CreateDamageLevel)

	// Test cases
	tests := []struct {
		name           string
		requestBody    handler.CreateDamageLevelRequest
		mockSetup      func(mockUseCase *mockusecase.MockDamageLevelUseCase, req handler.CreateDamageLevelRequest)
		expectedStatus int
		expectedBody   *handler.DamageLevelResponse
	}{
		{
			name: "Success",
			requestBody: handler.CreateDamageLevelRequest{
				Name:        "新規被害程度",
				Description: strPtr("新規の説明"),
			},
			mockSetup: func(mockUseCase *mockusecase.MockDamageLevelUseCase, req handler.CreateDamageLevelRequest) {
				mockUseCase.EXPECT().
					CreateDamageLevel(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ context.Context, dl *model.DamageLevel) error {
						// Verify the damage level properties
						if dl.Name != req.Name || *dl.Description != *req.Description {
							t.Errorf("Expected damage level with name %s and description %s, got name %s and description %s",
								req.Name, *req.Description, dl.Name, *dl.Description)
						}
						// Simulate ID assignment
						dl.ID = 1
						return nil
					})
			},
			expectedStatus: http.StatusCreated,
			expectedBody: &handler.DamageLevelResponse{
				ID:          1,
				Name:        "新規被害程度",
				Description: strPtr("新規の説明"),
			},
		},
		{
			name: "Invalid Request",
			requestBody: handler.CreateDamageLevelRequest{
				// Missing required Name field
				Description: strPtr("説明"),
			},
			mockSetup:      func(mockUseCase *mockusecase.MockDamageLevelUseCase, req handler.CreateDamageLevelRequest) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name: "Creation Error",
			requestBody: handler.CreateDamageLevelRequest{
				Name:        "エラー被害程度",
				Description: strPtr("エラーの説明"),
			},
			mockSetup: func(mockUseCase *mockusecase.MockDamageLevelUseCase, req handler.CreateDamageLevelRequest) {
				mockUseCase.EXPECT().
					CreateDamageLevel(gomock.Any(), gomock.Any()).
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
			tt.mockSetup(mockUseCase, tt.requestBody)

			// Prepare request
			reqBody, _ := json.Marshal(tt.requestBody)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/damage-levels", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusCreated {
				var response handler.DamageLevelResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, &response)
			}
		})
	}
}

func TestDamageLevelHandler_UpdateDamageLevel(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupTest(t)
	r.PUT("/damage-levels/:id", h.UpdateDamageLevel)

	// Test cases
	tests := []struct {
		name           string
		damageLevelID  string
		requestBody    handler.UpdateDamageLevelRequest
		mockSetup      func(mockUseCase *mockusecase.MockDamageLevelUseCase)
		expectedStatus int
		expectedBody   *handler.DamageLevelResponse
	}{
		{
			name:          "Success",
			damageLevelID: "1",
			requestBody: handler.UpdateDamageLevelRequest{
				Name:        "更新被害程度",
				Description: strPtr("更新の説明"),
			},
			mockSetup: func(mockUseCase *mockusecase.MockDamageLevelUseCase) {
				// First mock the GetDamageLevelByID call
				description := "元の説明"
				existingDamageLevel := &model.DamageLevel{
					ID:          1,
					Name:        "元の被害程度",
					Description: &description,
				}
				mockUseCase.EXPECT().GetDamageLevelByID(gomock.Any(), int32(1)).Return(existingDamageLevel, nil)

				// Then mock the UpdateDamageLevel call
				mockUseCase.EXPECT().
					UpdateDamageLevel(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ context.Context, dl *model.DamageLevel) error {
						// Verify the damage level properties
						if dl.ID != 1 || dl.Name != "更新被害程度" || *dl.Description != "更新の説明" {
							t.Errorf("Expected damage level with ID 1, name 更新被害程度 and description 更新の説明, got ID %d, name %s and description %s",
								dl.ID, dl.Name, *dl.Description)
						}
						return nil
					})
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.DamageLevelResponse{
				ID:          1,
				Name:        "更新被害程度",
				Description: strPtr("更新の説明"),
			},
		},
		{
			name:          "Invalid ID",
			damageLevelID: "invalid",
			requestBody: handler.UpdateDamageLevelRequest{
				Name: "更新被害程度",
			},
			mockSetup:      func(mockUseCase *mockusecase.MockDamageLevelUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name:          "Not Found",
			damageLevelID: "999",
			requestBody: handler.UpdateDamageLevelRequest{
				Name: "更新被害程度",
			},
			mockSetup: func(mockUseCase *mockusecase.MockDamageLevelUseCase) {
				mockUseCase.EXPECT().GetDamageLevelByID(gomock.Any(), int32(999)).Return(nil, errors.New("not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   nil,
		},
		{
			name:          "Update Error",
			damageLevelID: "1",
			requestBody: handler.UpdateDamageLevelRequest{
				Name: "エラー被害程度",
			},
			mockSetup: func(mockUseCase *mockusecase.MockDamageLevelUseCase) {
				// First mock the GetDamageLevelByID call
				description := "元の説明"
				existingDamageLevel := &model.DamageLevel{
					ID:          1,
					Name:        "元の被害程度",
					Description: &description,
				}
				mockUseCase.EXPECT().GetDamageLevelByID(gomock.Any(), int32(1)).Return(existingDamageLevel, nil)

				// Then mock the UpdateDamageLevel call with an error
				mockUseCase.EXPECT().
					UpdateDamageLevel(gomock.Any(), gomock.Any()).
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

			// Prepare request
			reqBody, _ := json.Marshal(tt.requestBody)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPut, "/damage-levels/"+tt.damageLevelID, bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response handler.DamageLevelResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, &response)
			}
		})
	}
}

func TestDamageLevelHandler_DeleteDamageLevel(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupTest(t)
	r.DELETE("/damage-levels/:id", h.DeleteDamageLevel)

	// Test cases
	tests := []struct {
		name           string
		damageLevelID  string
		mockSetup      func(mockUseCase *mockusecase.MockDamageLevelUseCase)
		expectedStatus int
	}{
		{
			name:          "Success",
			damageLevelID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockDamageLevelUseCase) {
				// First mock the GetDamageLevelByID call to check existence
				description := "説明"
				existingDamageLevel := &model.DamageLevel{
					ID:          1,
					Name:        "被害程度",
					Description: &description,
				}
				mockUseCase.EXPECT().GetDamageLevelByID(gomock.Any(), int32(1)).Return(existingDamageLevel, nil)

				// Then mock the DeleteDamageLevel call
				mockUseCase.EXPECT().DeleteDamageLevel(gomock.Any(), int32(1)).Return(nil)
			},
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "Invalid ID",
			damageLevelID:  "invalid",
			mockSetup:      func(mockUseCase *mockusecase.MockDamageLevelUseCase) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:          "Not Found",
			damageLevelID: "999",
			mockSetup: func(mockUseCase *mockusecase.MockDamageLevelUseCase) {
				mockUseCase.EXPECT().GetDamageLevelByID(gomock.Any(), int32(999)).Return(nil, errors.New("not found"))
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name:          "Delete Error",
			damageLevelID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockDamageLevelUseCase) {
				// First mock the GetDamageLevelByID call to check existence
				description := "説明"
				existingDamageLevel := &model.DamageLevel{
					ID:          1,
					Name:        "被害程度",
					Description: &description,
				}
				mockUseCase.EXPECT().GetDamageLevelByID(gomock.Any(), int32(1)).Return(existingDamageLevel, nil)

				// Then mock the DeleteDamageLevel call with an error
				mockUseCase.EXPECT().DeleteDamageLevel(gomock.Any(), int32(1)).Return(errors.New("database error"))
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
			req, _ := http.NewRequest(http.MethodDelete, "/damage-levels/"+tt.damageLevelID, nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

// Helper function to create string pointers
func strPtr(s string) *string {
	return &s
}
