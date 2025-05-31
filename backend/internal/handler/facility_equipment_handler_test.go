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

func setupFacilityEquipmentTest(t *testing.T) (*gin.Engine, *mockusecase.MockFacilityEquipmentUseCase, handler.FacilityEquipment) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	ctrl := gomock.NewController(t)
	mockUseCase := mockusecase.NewMockFacilityEquipmentUseCase(ctrl)
	l := logger.New(logger.DefaultConfig())
	h := handler.NewFacilityEquipmentHandler(l, mockUseCase)
	return r, mockUseCase, h
}

func TestFacilityEquipmentHandler_ListFacilityEquipments(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupFacilityEquipmentTest(t)
	r.GET("/facility-equipment", h.ListFacilityEquipments)

	// Test cases
	tests := []struct {
		name           string
		mockSetup      func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase)
		expectedStatus int
		expectedBody   []*handler.FacilityEquipmentResponse
	}{
		{
			name: "Success",
			mockSetup: func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase) {
				modelNumber := "Model-123"
				manufacturer := "Test Manufacturer"
				installationDate := time.Now()
				locationDesc := "Building A, Floor 2"
				latitude := 35.6812
				longitude := 139.7671
				notes := "Test notes"

				facilityEquipments := []*model.FacilityEquipment{
					{
						ID:                  1,
						Name:                "Test Equipment 1",
						FacilityTypeID:      1,
						ModelNumber:         &modelNumber,
						Manufacturer:        &manufacturer,
						InstallationDate:    &installationDate,
						Status:              "稼働中",
						LocationDescription: &locationDesc,
						LocationLatitude:    &latitude,
						LocationLongitude:   &longitude,
						Notes:               &notes,
					},
					{
						ID:             2,
						Name:           "Test Equipment 2",
						FacilityTypeID: 2,
						Status:         "メンテナンス中",
					},
				}
				mockUseCase.EXPECT().ListFacilityEquipments(gomock.Any()).Return(facilityEquipments, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: []*handler.FacilityEquipmentResponse{
				{
					ID:                  1,
					Name:                "Test Equipment 1",
					FacilityTypeID:      1,
					ModelNumber:         strPtr("Model-123"),
					Manufacturer:        strPtr("Test Manufacturer"),
					Status:              "稼働中",
					LocationDescription: strPtr("Building A, Floor 2"),
					LocationLatitude:    float64Ptr(35.6812),
					LocationLongitude:   float64Ptr(139.7671),
					Notes:               strPtr("Test notes"),
				},
				{
					ID:             2,
					Name:           "Test Equipment 2",
					FacilityTypeID: 2,
					Status:         "メンテナンス中",
				},
			},
		},
		{
			name: "Error",
			mockSetup: func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase) {
				mockUseCase.EXPECT().ListFacilityEquipments(gomock.Any()).Return(nil, errors.New("database error"))
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
			req, _ := http.NewRequest(http.MethodGet, "/facility-equipment", nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response []*handler.FacilityEquipmentResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, len(tt.expectedBody), len(response))

				for i, expected := range tt.expectedBody {
					assert.Equal(t, expected.ID, response[i].ID)
					assert.Equal(t, expected.Name, response[i].Name)
					assert.Equal(t, expected.FacilityTypeID, response[i].FacilityTypeID)
					assert.Equal(t, expected.Status, response[i].Status)

					// Check optional fields only if they exist in the expected response
					if expected.ModelNumber != nil {
						assert.Equal(t, *expected.ModelNumber, *response[i].ModelNumber)
					}
					if expected.Manufacturer != nil {
						assert.Equal(t, *expected.Manufacturer, *response[i].Manufacturer)
					}
					if expected.LocationDescription != nil {
						assert.Equal(t, *expected.LocationDescription, *response[i].LocationDescription)
					}
					if expected.LocationLatitude != nil {
						assert.Equal(t, *expected.LocationLatitude, *response[i].LocationLatitude)
					}
					if expected.LocationLongitude != nil {
						assert.Equal(t, *expected.LocationLongitude, *response[i].LocationLongitude)
					}
					if expected.Notes != nil {
						assert.Equal(t, *expected.Notes, *response[i].Notes)
					}
				}
			}
		})
	}
}

func TestFacilityEquipmentHandler_GetFacilityEquipment(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupFacilityEquipmentTest(t)
	r.GET("/facility-equipment/:id", h.GetFacilityEquipment)

	// Test cases
	tests := []struct {
		name           string
		equipmentID    string
		mockSetup      func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase)
		expectedStatus int
		expectedBody   *handler.FacilityEquipmentResponse
	}{
		{
			name:        "Success",
			equipmentID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase) {
				modelNumber := "Model-123"
				manufacturer := "Test Manufacturer"
				installationDate := time.Now()
				locationDesc := "Building A, Floor 2"
				latitude := 35.6812
				longitude := 139.7671
				notes := "Test notes"

				equipment := &model.FacilityEquipment{
					ID:                  1,
					Name:                "Test Equipment",
					FacilityTypeID:      1,
					ModelNumber:         &modelNumber,
					Manufacturer:        &manufacturer,
					InstallationDate:    &installationDate,
					Status:              "稼働中",
					LocationDescription: &locationDesc,
					LocationLatitude:    &latitude,
					LocationLongitude:   &longitude,
					Notes:               &notes,
				}
				mockUseCase.EXPECT().GetFacilityEquipmentByID(gomock.Any(), int32(1)).Return(equipment, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.FacilityEquipmentResponse{
				ID:                  1,
				Name:                "Test Equipment",
				FacilityTypeID:      1,
				ModelNumber:         strPtr("Model-123"),
				Manufacturer:        strPtr("Test Manufacturer"),
				Status:              "稼働中",
				LocationDescription: strPtr("Building A, Floor 2"),
				LocationLatitude:    float64Ptr(35.6812),
				LocationLongitude:   float64Ptr(139.7671),
				Notes:               strPtr("Test notes"),
			},
		},
		{
			name:           "Invalid ID",
			equipmentID:    "invalid",
			mockSetup:      func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name:        "Not Found",
			equipmentID: "999",
			mockSetup: func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase) {
				mockUseCase.EXPECT().GetFacilityEquipmentByID(gomock.Any(), int32(999)).Return(nil, errors.New("not found"))
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
			req, _ := http.NewRequest(http.MethodGet, "/facility-equipment/"+tt.equipmentID, nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response handler.FacilityEquipmentResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody.ID, response.ID)
				assert.Equal(t, tt.expectedBody.Name, response.Name)
				assert.Equal(t, tt.expectedBody.FacilityTypeID, response.FacilityTypeID)
				assert.Equal(t, tt.expectedBody.Status, response.Status)

				// Check optional fields only if they exist in the expected response
				if tt.expectedBody.ModelNumber != nil {
					assert.Equal(t, *tt.expectedBody.ModelNumber, *response.ModelNumber)
				}
				if tt.expectedBody.Manufacturer != nil {
					assert.Equal(t, *tt.expectedBody.Manufacturer, *response.Manufacturer)
				}
				if tt.expectedBody.LocationDescription != nil {
					assert.Equal(t, *tt.expectedBody.LocationDescription, *response.LocationDescription)
				}
				if tt.expectedBody.LocationLatitude != nil {
					assert.Equal(t, *tt.expectedBody.LocationLatitude, *response.LocationLatitude)
				}
				if tt.expectedBody.LocationLongitude != nil {
					assert.Equal(t, *tt.expectedBody.LocationLongitude, *response.LocationLongitude)
				}
				if tt.expectedBody.Notes != nil {
					assert.Equal(t, *tt.expectedBody.Notes, *response.Notes)
				}
			}
		})
	}
}

func TestFacilityEquipmentHandler_CreateFacilityEquipment(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupFacilityEquipmentTest(t)
	r.POST("/facility-equipment", h.CreateFacilityEquipment)

	// Helper function to create a time pointer
	now := time.Now()

	// Test cases
	tests := []struct {
		name           string
		requestBody    handler.CreateFacilityEquipmentRequest
		mockSetup      func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase)
		expectedStatus int
		expectedBody   *handler.FacilityEquipmentResponse
	}{
		{
			name: "Success",
			requestBody: handler.CreateFacilityEquipmentRequest{
				Name:             "New Equipment",
				FacilityTypeID:   1,
				ModelNumber:      strPtr("Model-XYZ"),
				Manufacturer:     strPtr("New Manufacturer"),
				InstallationDate: &now,
				Status:           "稼働中",
			},
			mockSetup: func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase) {
				mockUseCase.EXPECT().
					CreateFacilityEquipment(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ interface{}, equipment *model.FacilityEquipment) error {
						equipment.ID = 1 // Simulate ID generation
						return nil
					})
			},
			expectedStatus: http.StatusCreated,
			expectedBody: &handler.FacilityEquipmentResponse{
				ID:               1,
				Name:             "New Equipment",
				FacilityTypeID:   1,
				ModelNumber:      strPtr("Model-XYZ"),
				Manufacturer:     strPtr("New Manufacturer"),
				InstallationDate: &now,
				Status:           "稼働中",
			},
		},
		{
			name: "Invalid Request - Missing Required Field",
			requestBody: handler.CreateFacilityEquipmentRequest{
				// Missing Name
				FacilityTypeID: 1,
				Status:         "稼働中",
			},
			mockSetup:      func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name: "Database Error",
			requestBody: handler.CreateFacilityEquipmentRequest{
				Name:           "New Equipment",
				FacilityTypeID: 1,
				Status:         "稼働中",
			},
			mockSetup: func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase) {
				mockUseCase.EXPECT().
					CreateFacilityEquipment(gomock.Any(), gomock.Any()).
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
			req, _ := http.NewRequest(http.MethodPost, "/facility-equipment", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusCreated {
				var response handler.FacilityEquipmentResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				// Check fields
				assert.Equal(t, tt.expectedBody.Name, response.Name)
				assert.Equal(t, tt.expectedBody.FacilityTypeID, response.FacilityTypeID)
				assert.Equal(t, tt.expectedBody.Status, response.Status)

				// Check optional fields only if they exist in the expected response
				if tt.expectedBody.ModelNumber != nil {
					assert.Equal(t, *tt.expectedBody.ModelNumber, *response.ModelNumber)
				}
				if tt.expectedBody.Manufacturer != nil {
					assert.Equal(t, *tt.expectedBody.Manufacturer, *response.Manufacturer)
				}
			}
		})
	}
}

func TestFacilityEquipmentHandler_UpdateFacilityEquipment(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupFacilityEquipmentTest(t)
	r.PUT("/facility-equipment/:id", h.UpdateFacilityEquipment)

	// Helper function to create a time pointer
	now := time.Now()

	// Test cases
	tests := []struct {
		name           string
		equipmentID    string
		requestBody    handler.UpdateFacilityEquipmentRequest
		mockSetup      func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase)
		expectedStatus int
		expectedBody   *handler.FacilityEquipmentResponse
	}{
		{
			name:        "Success",
			equipmentID: "1",
			requestBody: handler.UpdateFacilityEquipmentRequest{
				Name:             "Updated Equipment",
				FacilityTypeID:   2,
				ModelNumber:      strPtr("Updated-Model"),
				Manufacturer:     strPtr("Updated Manufacturer"),
				InstallationDate: &now,
				Status:           "メンテナンス中",
			},
			mockSetup: func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase) {
				oldModelNumber := "Old-Model"
				oldManufacturer := "Old Manufacturer"
				oldInstallationDate := time.Now().Add(-24 * time.Hour)
				oldLocationDesc := "Old Location"
				oldLatitude := 35.6812
				oldLongitude := 139.7671
				oldNotes := "Old notes"

				equipment := &model.FacilityEquipment{
					ID:                  1,
					Name:                "Old Equipment",
					FacilityTypeID:      1,
					ModelNumber:         &oldModelNumber,
					Manufacturer:        &oldManufacturer,
					InstallationDate:    &oldInstallationDate,
					Status:              "稼働中",
					LocationDescription: &oldLocationDesc,
					LocationLatitude:    &oldLatitude,
					LocationLongitude:   &oldLongitude,
					Notes:               &oldNotes,
				}
				mockUseCase.EXPECT().GetFacilityEquipmentByID(gomock.Any(), int32(1)).Return(equipment, nil)
				mockUseCase.EXPECT().
					UpdateFacilityEquipment(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ interface{}, e *model.FacilityEquipment) error {
						assert.Equal(t, "Updated Equipment", e.Name)
						assert.Equal(t, int32(2), e.FacilityTypeID)
						assert.Equal(t, "Updated-Model", *e.ModelNumber)
						assert.Equal(t, "Updated Manufacturer", *e.Manufacturer)
						assert.Equal(t, "メンテナンス中", e.Status)
						return nil
					})
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.FacilityEquipmentResponse{
				ID:               1,
				Name:             "Updated Equipment",
				FacilityTypeID:   2,
				ModelNumber:      strPtr("Updated-Model"),
				Manufacturer:     strPtr("Updated Manufacturer"),
				InstallationDate: &now,
				Status:           "メンテナンス中",
			},
		},
		{
			name:        "Invalid ID",
			equipmentID: "invalid",
			requestBody: handler.UpdateFacilityEquipmentRequest{
				Name:           "Updated Equipment",
				FacilityTypeID: 2,
				Status:         "メンテナンス中",
			},
			mockSetup:      func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name:        "Not Found",
			equipmentID: "999",
			requestBody: handler.UpdateFacilityEquipmentRequest{
				Name:           "Updated Equipment",
				FacilityTypeID: 2,
				Status:         "メンテナンス中",
			},
			mockSetup: func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase) {
				mockUseCase.EXPECT().GetFacilityEquipmentByID(gomock.Any(), int32(999)).Return(nil, errors.New("not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   nil,
		},
		{
			name:        "Database Error",
			equipmentID: "1",
			requestBody: handler.UpdateFacilityEquipmentRequest{
				Name:           "Updated Equipment",
				FacilityTypeID: 2,
				Status:         "メンテナンス中",
			},
			mockSetup: func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase) {
				equipment := &model.FacilityEquipment{
					ID:             1,
					Name:           "Old Equipment",
					FacilityTypeID: 1,
					Status:         "稼働中",
				}
				mockUseCase.EXPECT().GetFacilityEquipmentByID(gomock.Any(), int32(1)).Return(equipment, nil)
				mockUseCase.EXPECT().UpdateFacilityEquipment(gomock.Any(), gomock.Any()).Return(errors.New("database error"))
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
			req, _ := http.NewRequest(http.MethodPut, "/facility-equipment/"+tt.equipmentID, bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response handler.FacilityEquipmentResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.requestBody.Name, response.Name)
				assert.Equal(t, tt.requestBody.FacilityTypeID, response.FacilityTypeID)
				assert.Equal(t, tt.requestBody.Status, response.Status)

				if tt.requestBody.ModelNumber != nil {
					assert.Equal(t, *tt.requestBody.ModelNumber, *response.ModelNumber)
				}
				if tt.requestBody.Manufacturer != nil {
					assert.Equal(t, *tt.requestBody.Manufacturer, *response.Manufacturer)
				}
			}
		})
	}
}

func TestFacilityEquipmentHandler_DeleteFacilityEquipment(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupFacilityEquipmentTest(t)
	r.DELETE("/facility-equipment/:id", h.DeleteFacilityEquipment)

	// Test cases
	tests := []struct {
		name           string
		equipmentID    string
		mockSetup      func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase)
		expectedStatus int
	}{
		{
			name:        "Success",
			equipmentID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase) {
				equipment := &model.FacilityEquipment{
					ID:             1,
					Name:           "Test Equipment",
					FacilityTypeID: 1,
					Status:         "稼働中",
				}
				mockUseCase.EXPECT().GetFacilityEquipmentByID(gomock.Any(), int32(1)).Return(equipment, nil)
				mockUseCase.EXPECT().DeleteFacilityEquipment(gomock.Any(), int32(1)).Return(nil)
			},
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "Invalid ID",
			equipmentID:    "invalid",
			mockSetup:      func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:        "Not Found",
			equipmentID: "999",
			mockSetup: func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase) {
				mockUseCase.EXPECT().GetFacilityEquipmentByID(gomock.Any(), int32(999)).Return(nil, errors.New("not found"))
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name:        "Database Error",
			equipmentID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockFacilityEquipmentUseCase) {
				equipment := &model.FacilityEquipment{
					ID:             1,
					Name:           "Test Equipment",
					FacilityTypeID: 1,
					Status:         "稼働中",
				}
				mockUseCase.EXPECT().GetFacilityEquipmentByID(gomock.Any(), int32(1)).Return(equipment, nil)
				mockUseCase.EXPECT().DeleteFacilityEquipment(gomock.Any(), int32(1)).Return(errors.New("database error"))
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
			req, _ := http.NewRequest(http.MethodDelete, "/facility-equipment/"+tt.equipmentID, nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

// Helper function for creating pointers to float64
func float64Ptr(f float64) *float64 {
	return &f
}
