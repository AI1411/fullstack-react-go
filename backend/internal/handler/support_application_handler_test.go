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

func setupSupportApplicationTest(t *testing.T) (*gin.Engine, *mockusecase.MockSupportApplicationUseCase, handler.SupportApplication) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	ctrl := gomock.NewController(t)
	mockUseCase := mockusecase.NewMockSupportApplicationUseCase(ctrl)
	l := logger.New(logger.DefaultConfig())
	h := handler.NewSupportApplicationHandler(l, mockUseCase)
	return r, mockUseCase, h
}

func TestSupportApplicationHandler_ListSupportApplications(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupSupportApplicationTest(t)
	r.GET("/support-applications", h.ListSupportApplications)

	// Test cases
	tests := []struct {
		name           string
		mockSetup      func(mockUseCase *mockusecase.MockSupportApplicationUseCase)
		expectedStatus int
		expectedBody   *handler.ListSupportApplicationsResponse
	}{
		{
			name: "Success",
			mockSetup: func(mockUseCase *mockusecase.MockSupportApplicationUseCase) {
				notes := "Test notes"
				reviewedAt := time.Now()
				approvedAt := time.Now()
				completedAt := time.Now()

				supportApplications := []*model.SupportApplication{
					{
						ApplicationID:   "APP-001",
						ApplicationDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
						ApplicantName:   "Test Applicant 1",
						DisasterName:    "Test Disaster 1",
						RequestedAmount: 100000,
						Status:          "審査中",
						ReviewedAt:      &reviewedAt,
						ApprovedAt:      &approvedAt,
						CompletedAt:     &completedAt,
						Notes:           &notes,
						CreatedAt:       time.Now(),
						UpdatedAt:       time.Now(),
					},
					{
						ApplicationID:   "APP-002",
						ApplicationDate: time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
						ApplicantName:   "Test Applicant 2",
						DisasterName:    "Test Disaster 2",
						RequestedAmount: 200000,
						Status:          "承認済み",
						CreatedAt:       time.Now(),
						UpdatedAt:       time.Now(),
					},
				}
				mockUseCase.EXPECT().ListSupportApplications(gomock.Any()).Return(supportApplications, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.ListSupportApplicationsResponse{
				SupportApplications: []*handler.SupportApplicationResponse{
					{
						ApplicationID:   "APP-001",
						ApplicationDate: "2023-01-01",
						ApplicantName:   "Test Applicant 1",
						DisasterName:    "Test Disaster 1",
						RequestedAmount: 100000,
						Status:          "審査中",
						ReviewedAt:      strPtr(time.Now().Format(time.DateTime)),
						ApprovedAt:      strPtr(time.Now().Format(time.DateTime)),
						CompletedAt:     strPtr(time.Now().Format(time.DateTime)),
						Notes:           strPtr("Test notes"),
					},
					{
						ApplicationID:   "APP-002",
						ApplicationDate: "2023-02-01",
						ApplicantName:   "Test Applicant 2",
						DisasterName:    "Test Disaster 2",
						RequestedAmount: 200000,
						Status:          "承認済み",
					},
				},
				Total: 2,
			},
		},
		{
			name: "Error",
			mockSetup: func(mockUseCase *mockusecase.MockSupportApplicationUseCase) {
				mockUseCase.EXPECT().ListSupportApplications(gomock.Any()).Return(nil, errors.New("database error"))
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
			req, _ := http.NewRequest(http.MethodGet, "/support-applications", nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response handler.ListSupportApplicationsResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody.Total, response.Total)
				assert.Equal(t, len(tt.expectedBody.SupportApplications), len(response.SupportApplications))

				for i, expected := range tt.expectedBody.SupportApplications {
					assert.Equal(t, expected.ApplicationID, response.SupportApplications[i].ApplicationID)
					assert.Equal(t, expected.ApplicationDate, response.SupportApplications[i].ApplicationDate)
					assert.Equal(t, expected.ApplicantName, response.SupportApplications[i].ApplicantName)
					assert.Equal(t, expected.DisasterName, response.SupportApplications[i].DisasterName)
					assert.Equal(t, expected.RequestedAmount, response.SupportApplications[i].RequestedAmount)
					assert.Equal(t, expected.Status, response.SupportApplications[i].Status)

					// Check optional fields only if they exist in the expected response
					if expected.Notes != nil {
						assert.NotNil(t, response.SupportApplications[i].Notes)
					}
					if expected.ReviewedAt != nil {
						assert.NotNil(t, response.SupportApplications[i].ReviewedAt)
					}
					if expected.ApprovedAt != nil {
						assert.NotNil(t, response.SupportApplications[i].ApprovedAt)
					}
					if expected.CompletedAt != nil {
						assert.NotNil(t, response.SupportApplications[i].CompletedAt)
					}
				}
			}
		})
	}
}

func TestSupportApplicationHandler_GetSupportApplication(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupSupportApplicationTest(t)
	r.GET("/support-applications/:id", h.GetSupportApplication)

	// Test cases
	tests := []struct {
		name           string
		applicationID  string
		mockSetup      func(mockUseCase *mockusecase.MockSupportApplicationUseCase)
		expectedStatus int
		expectedBody   *handler.SupportApplicationResponse
	}{
		{
			name:          "Success",
			applicationID: "APP-001",
			mockSetup: func(mockUseCase *mockusecase.MockSupportApplicationUseCase) {
				notes := "Test notes"
				reviewedAt := time.Now()
				approvedAt := time.Now()
				completedAt := time.Now()

				supportApplication := &model.SupportApplication{
					ApplicationID:   "APP-001",
					ApplicationDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					ApplicantName:   "Test Applicant",
					DisasterName:    "Test Disaster",
					RequestedAmount: 100000,
					Status:          "審査中",
					ReviewedAt:      &reviewedAt,
					ApprovedAt:      &approvedAt,
					CompletedAt:     &completedAt,
					Notes:           &notes,
					CreatedAt:       time.Now(),
					UpdatedAt:       time.Now(),
				}
				mockUseCase.EXPECT().GetSupportApplicationByID(gomock.Any(), "APP-001").Return(supportApplication, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.SupportApplicationResponse{
				ApplicationID:   "APP-001",
				ApplicationDate: "2023-01-01",
				ApplicantName:   "Test Applicant",
				DisasterName:    "Test Disaster",
				RequestedAmount: 100000,
				Status:          "審査中",
				ReviewedAt:      strPtr(time.Now().Format(time.DateTime)),
				ApprovedAt:      strPtr(time.Now().Format(time.DateTime)),
				CompletedAt:     strPtr(time.Now().Format(time.DateTime)),
				Notes:           strPtr("Test notes"),
			},
		},
		{
			name:          "Not Found",
			applicationID: "APP-999",
			mockSetup: func(mockUseCase *mockusecase.MockSupportApplicationUseCase) {
				mockUseCase.EXPECT().GetSupportApplicationByID(gomock.Any(), "APP-999").Return(nil, errors.New("not found"))
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
			req, _ := http.NewRequest(http.MethodGet, "/support-applications/"+tt.applicationID, nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response handler.SupportApplicationResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody.ApplicationID, response.ApplicationID)
				assert.Equal(t, tt.expectedBody.ApplicationDate, response.ApplicationDate)
				assert.Equal(t, tt.expectedBody.ApplicantName, response.ApplicantName)
				assert.Equal(t, tt.expectedBody.DisasterName, response.DisasterName)
				assert.Equal(t, tt.expectedBody.RequestedAmount, response.RequestedAmount)
				assert.Equal(t, tt.expectedBody.Status, response.Status)

				// Check optional fields only if they exist in the expected response
				if tt.expectedBody.Notes != nil {
					assert.NotNil(t, response.Notes)
				}
				if tt.expectedBody.ReviewedAt != nil {
					assert.NotNil(t, response.ReviewedAt)
				}
				if tt.expectedBody.ApprovedAt != nil {
					assert.NotNil(t, response.ApprovedAt)
				}
				if tt.expectedBody.CompletedAt != nil {
					assert.NotNil(t, response.CompletedAt)
				}
			}
		})
	}
}

func TestSupportApplicationHandler_CreateSupportApplication(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupSupportApplicationTest(t)
	r.POST("/support-applications", h.CreateSupportApplication)

	// Test cases
	tests := []struct {
		name           string
		requestBody    handler.CreateSupportApplicationRequest
		mockSetup      func(mockUseCase *mockusecase.MockSupportApplicationUseCase)
		expectedStatus int
		expectedBody   *handler.SupportApplicationResponse
	}{
		{
			name: "Success",
			requestBody: handler.CreateSupportApplicationRequest{
				ApplicationID:   "APP-003",
				ApplicationDate: "2023-03-01",
				ApplicantName:   "New Applicant",
				DisasterName:    "New Disaster",
				RequestedAmount: 300000,
				Status:          "審査中",
				Notes:           strPtr("New notes"),
			},
			mockSetup: func(mockUseCase *mockusecase.MockSupportApplicationUseCase) {
				mockUseCase.EXPECT().
					CreateSupportApplication(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ interface{}, application *model.SupportApplication) error {
						// Verify the application data
						assert.Equal(t, "APP-003", application.ApplicationID)
						assert.Equal(t, time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC), application.ApplicationDate)
						assert.Equal(t, "New Applicant", application.ApplicantName)
						assert.Equal(t, "New Disaster", application.DisasterName)
						assert.Equal(t, int64(300000), application.RequestedAmount)
						assert.Equal(t, "審査中", application.Status)
						assert.Equal(t, "New notes", *application.Notes)
						return nil
					})
			},
			expectedStatus: http.StatusCreated,
			expectedBody: &handler.SupportApplicationResponse{
				ApplicationID:   "APP-003",
				ApplicationDate: "2023-03-01",
				ApplicantName:   "New Applicant",
				DisasterName:    "New Disaster",
				RequestedAmount: 300000,
				Status:          "審査中",
				Notes:           strPtr("New notes"),
			},
		},
		{
			name: "Invalid Request - Missing Required Field",
			requestBody: handler.CreateSupportApplicationRequest{
				// Missing ApplicationID
				ApplicationDate: "2023-03-01",
				ApplicantName:   "New Applicant",
				DisasterName:    "New Disaster",
				RequestedAmount: 300000,
			},
			mockSetup:      func(mockUseCase *mockusecase.MockSupportApplicationUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name: "Invalid Date Format",
			requestBody: handler.CreateSupportApplicationRequest{
				ApplicationID:   "APP-003",
				ApplicationDate: "2023/03/01", // Wrong format
				ApplicantName:   "New Applicant",
				DisasterName:    "New Disaster",
				RequestedAmount: 300000,
			},
			mockSetup:      func(mockUseCase *mockusecase.MockSupportApplicationUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name: "Database Error",
			requestBody: handler.CreateSupportApplicationRequest{
				ApplicationID:   "APP-003",
				ApplicationDate: "2023-03-01",
				ApplicantName:   "New Applicant",
				DisasterName:    "New Disaster",
				RequestedAmount: 300000,
			},
			mockSetup: func(mockUseCase *mockusecase.MockSupportApplicationUseCase) {
				mockUseCase.EXPECT().
					CreateSupportApplication(gomock.Any(), gomock.Any()).
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
			req, _ := http.NewRequest(http.MethodPost, "/support-applications", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusCreated {
				var response handler.SupportApplicationResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				// Check fields
				assert.Equal(t, tt.requestBody.ApplicationID, response.ApplicationID)
				assert.Equal(t, tt.requestBody.ApplicationDate, response.ApplicationDate)
				assert.Equal(t, tt.requestBody.ApplicantName, response.ApplicantName)
				assert.Equal(t, tt.requestBody.DisasterName, response.DisasterName)
				assert.Equal(t, tt.requestBody.RequestedAmount, response.RequestedAmount)
				assert.Equal(t, tt.requestBody.Status, response.Status)

				// Check optional fields only if they exist in the expected response
				if tt.requestBody.Notes != nil {
					assert.Equal(t, *tt.requestBody.Notes, *response.Notes)
				}
			}
		})
	}
}
