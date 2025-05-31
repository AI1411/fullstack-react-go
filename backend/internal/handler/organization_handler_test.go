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

func setupOrganizationTest(t *testing.T) (*gin.Engine, *mockusecase.MockOrganizationUseCase, handler.Organization) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	ctrl := gomock.NewController(t)
	mockUseCase := mockusecase.NewMockOrganizationUseCase(ctrl)
	l := logger.New(logger.DefaultConfig())
	h := handler.NewOrganizationHandler(l, mockUseCase)
	return r, mockUseCase, h
}

func TestOrganizationHandler_ListOrganizations(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupOrganizationTest(t)
	r.GET("/organizations", h.ListOrganizations)

	// Test cases
	tests := []struct {
		name           string
		mockSetup      func(mockUseCase *mockusecase.MockOrganizationUseCase)
		expectedStatus int
		expectedBody   []*handler.OrganizationResponse
	}{
		{
			name: "Success",
			mockSetup: func(mockUseCase *mockusecase.MockOrganizationUseCase) {
				prefectureID := int32(1)
				parentID := int32(2)
				description := "Test Organization Description"

				organizations := []*model.Organization{
					{
						ID:           1,
						Name:         "Test Organization 1",
						Type:         "government",
						PrefectureID: &prefectureID,
						ParentID:     &parentID,
						Description:  &description,
						CreatedAt:    time.Now(),
						UpdatedAt:    time.Now(),
					},
					{
						ID:        2,
						Name:      "Test Organization 2",
						Type:      "ngo",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				}
				mockUseCase.EXPECT().ListOrganizations(gomock.Any()).Return(organizations, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: []*handler.OrganizationResponse{
				{
					ID:           1,
					Name:         "Test Organization 1",
					Type:         "government",
					PrefectureID: int32Ptr(1),
					ParentID:     int32Ptr(2),
					Description:  strPtr("Test Organization Description"),
				},
				{
					ID:   2,
					Name: "Test Organization 2",
					Type: "ngo",
				},
			},
		},
		{
			name: "Error",
			mockSetup: func(mockUseCase *mockusecase.MockOrganizationUseCase) {
				mockUseCase.EXPECT().ListOrganizations(gomock.Any()).Return(nil, errors.New("database error"))
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
			req, _ := http.NewRequest(http.MethodGet, "/organizations", nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response []*handler.OrganizationResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, len(tt.expectedBody), len(response))

				for i, expected := range tt.expectedBody {
					assert.Equal(t, expected.ID, response[i].ID)
					assert.Equal(t, expected.Name, response[i].Name)
					assert.Equal(t, expected.Type, response[i].Type)

					// Check optional fields only if they exist in the expected response
					if expected.PrefectureID != nil {
						assert.Equal(t, *expected.PrefectureID, *response[i].PrefectureID)
					}
					if expected.ParentID != nil {
						assert.Equal(t, *expected.ParentID, *response[i].ParentID)
					}
					if expected.Description != nil {
						assert.Equal(t, *expected.Description, *response[i].Description)
					}
				}
			}
		})
	}
}

func TestOrganizationHandler_GetOrganization(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupOrganizationTest(t)
	r.GET("/organizations/:id", h.GetOrganization)

	// Test cases
	tests := []struct {
		name           string
		organizationID string
		mockSetup      func(mockUseCase *mockusecase.MockOrganizationUseCase)
		expectedStatus int
		expectedBody   *handler.OrganizationResponse
	}{
		{
			name:           "Success",
			organizationID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockOrganizationUseCase) {
				prefectureID := int32(1)
				parentID := int32(2)
				description := "Test Organization Description"

				organization := &model.Organization{
					ID:           1,
					Name:         "Test Organization",
					Type:         "government",
					PrefectureID: &prefectureID,
					ParentID:     &parentID,
					Description:  &description,
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					Users: []model.User{
						{
							ID:        1,
							Name:      "Test User",
							Email:     "test@example.com",
							CreatedAt: timePtr(time.Now()),
							UpdatedAt: timePtr(time.Now()),
						},
					},
				}
				mockUseCase.EXPECT().GetOrganizationByID(gomock.Any(), int32(1)).Return(organization, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.OrganizationResponse{
				ID:           1,
				Name:         "Test Organization",
				Type:         "government",
				PrefectureID: int32Ptr(1),
				ParentID:     int32Ptr(2),
				Description:  strPtr("Test Organization Description"),
				Users: []handler.UserResponse{
					{
						ID:        1,
						Name:      "Test User",
						Email:     "test@example.com",
						CreatedAt: nil,
						UpdatedAt: nil,
					},
				},
			},
		},
		{
			name:           "Invalid ID",
			organizationID: "invalid",
			mockSetup:      func(mockUseCase *mockusecase.MockOrganizationUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name:           "Not Found",
			organizationID: "999",
			mockSetup: func(mockUseCase *mockusecase.MockOrganizationUseCase) {
				mockUseCase.EXPECT().GetOrganizationByID(gomock.Any(), int32(999)).Return(nil, errors.New("not found"))
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
			req, _ := http.NewRequest(http.MethodGet, "/organizations/"+tt.organizationID, nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response handler.OrganizationResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody.ID, response.ID)
				assert.Equal(t, tt.expectedBody.Name, response.Name)
				assert.Equal(t, tt.expectedBody.Type, response.Type)

				// Check optional fields only if they exist in the expected response
				if tt.expectedBody.PrefectureID != nil {
					assert.Equal(t, *tt.expectedBody.PrefectureID, *response.PrefectureID)
				}
				if tt.expectedBody.ParentID != nil {
					assert.Equal(t, *tt.expectedBody.ParentID, *response.ParentID)
				}
				if tt.expectedBody.Description != nil {
					assert.Equal(t, *tt.expectedBody.Description, *response.Description)
				}

				// Check users
				if len(tt.expectedBody.Users) > 0 {
					assert.Equal(t, len(tt.expectedBody.Users), len(response.Users))
					for i, user := range tt.expectedBody.Users {
						assert.Equal(t, user.ID, response.Users[i].ID)
						assert.Equal(t, user.Name, response.Users[i].Name)
						assert.Equal(t, user.Email, response.Users[i].Email)
					}
				}
			}
		})
	}
}

func TestOrganizationHandler_CreateOrganization(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupOrganizationTest(t)
	r.POST("/organizations", h.CreateOrganization)

	// Test cases
	tests := []struct {
		name           string
		requestBody    handler.CreateOrganizationRequest
		mockSetup      func(mockUseCase *mockusecase.MockOrganizationUseCase)
		expectedStatus int
		expectedBody   *handler.OrganizationResponse
	}{
		{
			name: "Success",
			requestBody: handler.CreateOrganizationRequest{
				Name:         "New Organization",
				Type:         "government",
				PrefectureID: int32Ptr(1),
				ParentID:     int32Ptr(2),
				Description:  strPtr("New Organization Description"),
			},
			mockSetup: func(mockUseCase *mockusecase.MockOrganizationUseCase) {
				mockUseCase.EXPECT().
					CreateOrganization(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ interface{}, organization *model.Organization) error {
						organization.ID = 1 // Simulate ID generation
						organization.CreatedAt = time.Now()
						organization.UpdatedAt = time.Now()
						return nil
					})
			},
			expectedStatus: http.StatusCreated,
			expectedBody: &handler.OrganizationResponse{
				ID:           1,
				Name:         "New Organization",
				Type:         "government",
				PrefectureID: int32Ptr(1),
				ParentID:     int32Ptr(2),
				Description:  strPtr("New Organization Description"),
			},
		},
		{
			name: "Invalid Request - Missing Required Field",
			requestBody: handler.CreateOrganizationRequest{
				// Missing Name
				Type: "government",
			},
			mockSetup:      func(mockUseCase *mockusecase.MockOrganizationUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name: "Database Error",
			requestBody: handler.CreateOrganizationRequest{
				Name: "New Organization",
				Type: "government",
			},
			mockSetup: func(mockUseCase *mockusecase.MockOrganizationUseCase) {
				mockUseCase.EXPECT().
					CreateOrganization(gomock.Any(), gomock.Any()).
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
			req, _ := http.NewRequest(http.MethodPost, "/organizations", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusCreated {
				var response handler.OrganizationResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				// Check fields
				assert.Equal(t, tt.requestBody.Name, response.Name)
				assert.Equal(t, tt.requestBody.Type, response.Type)

				// Check optional fields only if they exist in the expected response
				if tt.requestBody.PrefectureID != nil {
					assert.Equal(t, *tt.requestBody.PrefectureID, *response.PrefectureID)
				}
				if tt.requestBody.ParentID != nil {
					assert.Equal(t, *tt.requestBody.ParentID, *response.ParentID)
				}
				if tt.requestBody.Description != nil {
					assert.Equal(t, *tt.requestBody.Description, *response.Description)
				}
			}
		})
	}
}

func TestOrganizationHandler_UpdateOrganization(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupOrganizationTest(t)
	r.PUT("/organizations/:id", h.UpdateOrganization)

	// Test cases
	tests := []struct {
		name           string
		organizationID string
		requestBody    handler.UpdateOrganizationRequest
		mockSetup      func(mockUseCase *mockusecase.MockOrganizationUseCase)
		expectedStatus int
		expectedBody   *handler.OrganizationResponse
	}{
		{
			name:           "Success",
			organizationID: "1",
			requestBody: handler.UpdateOrganizationRequest{
				Name:         "Updated Organization",
				Type:         "ngo",
				PrefectureID: int32Ptr(3),
				ParentID:     int32Ptr(4),
				Description:  strPtr("Updated Organization Description"),
			},
			mockSetup: func(mockUseCase *mockusecase.MockOrganizationUseCase) {
				oldPrefectureID := int32(1)
				oldParentID := int32(2)
				oldDescription := "Old Organization Description"

				organization := &model.Organization{
					ID:           1,
					Name:         "Old Organization",
					Type:         "government",
					PrefectureID: &oldPrefectureID,
					ParentID:     &oldParentID,
					Description:  &oldDescription,
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				}
				mockUseCase.EXPECT().GetOrganizationByID(gomock.Any(), int32(1)).Return(organization, nil)
				mockUseCase.EXPECT().
					UpdateOrganization(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ interface{}, o *model.Organization) error {
						assert.Equal(t, "Updated Organization", o.Name)
						assert.Equal(t, "ngo", o.Type)
						assert.Equal(t, int32(3), *o.PrefectureID)
						assert.Equal(t, int32(4), *o.ParentID)
						assert.Equal(t, "Updated Organization Description", *o.Description)
						return nil
					})
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.OrganizationResponse{
				ID:           1,
				Name:         "Updated Organization",
				Type:         "ngo",
				PrefectureID: int32Ptr(3),
				ParentID:     int32Ptr(4),
				Description:  strPtr("Updated Organization Description"),
			},
		},
		{
			name:           "Invalid ID",
			organizationID: "invalid",
			requestBody: handler.UpdateOrganizationRequest{
				Name: "Updated Organization",
				Type: "ngo",
			},
			mockSetup:      func(mockUseCase *mockusecase.MockOrganizationUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name:           "Not Found",
			organizationID: "999",
			requestBody: handler.UpdateOrganizationRequest{
				Name: "Updated Organization",
				Type: "ngo",
			},
			mockSetup: func(mockUseCase *mockusecase.MockOrganizationUseCase) {
				mockUseCase.EXPECT().GetOrganizationByID(gomock.Any(), int32(999)).Return(nil, errors.New("not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   nil,
		},
		{
			name:           "Database Error",
			organizationID: "1",
			requestBody: handler.UpdateOrganizationRequest{
				Name: "Updated Organization",
				Type: "ngo",
			},
			mockSetup: func(mockUseCase *mockusecase.MockOrganizationUseCase) {
				organization := &model.Organization{
					ID:        1,
					Name:      "Old Organization",
					Type:      "government",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
				mockUseCase.EXPECT().GetOrganizationByID(gomock.Any(), int32(1)).Return(organization, nil)
				mockUseCase.EXPECT().UpdateOrganization(gomock.Any(), gomock.Any()).Return(errors.New("database error"))
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
			req, _ := http.NewRequest(http.MethodPut, "/organizations/"+tt.organizationID, bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response handler.OrganizationResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.requestBody.Name, response.Name)
				assert.Equal(t, tt.requestBody.Type, response.Type)

				if tt.requestBody.PrefectureID != nil {
					assert.Equal(t, *tt.requestBody.PrefectureID, *response.PrefectureID)
				}
				if tt.requestBody.ParentID != nil {
					assert.Equal(t, *tt.requestBody.ParentID, *response.ParentID)
				}
				if tt.requestBody.Description != nil {
					assert.Equal(t, *tt.requestBody.Description, *response.Description)
				}
			}
		})
	}
}

func TestOrganizationHandler_DeleteOrganization(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupOrganizationTest(t)
	r.DELETE("/organizations/:id", h.DeleteOrganization)

	// Test cases
	tests := []struct {
		name           string
		organizationID string
		mockSetup      func(mockUseCase *mockusecase.MockOrganizationUseCase)
		expectedStatus int
	}{
		{
			name:           "Success",
			organizationID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockOrganizationUseCase) {
				organization := &model.Organization{
					ID:        1,
					Name:      "Test Organization",
					Type:      "government",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
				mockUseCase.EXPECT().GetOrganizationByID(gomock.Any(), int32(1)).Return(organization, nil)
				mockUseCase.EXPECT().DeleteOrganization(gomock.Any(), int32(1)).Return(nil)
			},
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "Invalid ID",
			organizationID: "invalid",
			mockSetup:      func(mockUseCase *mockusecase.MockOrganizationUseCase) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Not Found",
			organizationID: "999",
			mockSetup: func(mockUseCase *mockusecase.MockOrganizationUseCase) {
				mockUseCase.EXPECT().GetOrganizationByID(gomock.Any(), int32(999)).Return(nil, errors.New("not found"))
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Database Error",
			organizationID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockOrganizationUseCase) {
				organization := &model.Organization{
					ID:        1,
					Name:      "Test Organization",
					Type:      "government",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
				mockUseCase.EXPECT().GetOrganizationByID(gomock.Any(), int32(1)).Return(organization, nil)
				mockUseCase.EXPECT().DeleteOrganization(gomock.Any(), int32(1)).Return(errors.New("database error"))
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
			req, _ := http.NewRequest(http.MethodDelete, "/organizations/"+tt.organizationID, nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

// Helper function for creating pointers to int32
func int32Ptr(i int32) *int32 {
	return &i
}

// Helper function for creating pointers to time.Time
func timePtr(t time.Time) *time.Time {
	return &t
}
