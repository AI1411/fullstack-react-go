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

func setupUserTest(t *testing.T) (*gin.Engine, *mockusecase.MockUserUseCase, handler.User) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	ctrl := gomock.NewController(t)
	mockUseCase := mockusecase.NewMockUserUseCase(ctrl)
	l := logger.New(logger.DefaultConfig())
	h := handler.NewUserHandler(l, mockUseCase)
	return r, mockUseCase, h
}

func TestUserHandler_ListUsers(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupUserTest(t)
	r.GET("/users", h.ListUsers)

	// Test cases
	tests := []struct {
		name           string
		mockSetup      func(mockUseCase *mockusecase.MockUserUseCase)
		expectedStatus int
		expectedBody   []handler.UserResponse
	}{
		{
			name: "Success",
			mockSetup: func(mockUseCase *mockusecase.MockUserUseCase) {
				now := time.Now()
				users := []*model.User{
					{
						ID:        1,
						Name:      "User 1",
						Email:     "user1@example.com",
						CreatedAt: &now,
						UpdatedAt: &now,
					},
					{
						ID:        2,
						Name:      "User 2",
						Email:     "user2@example.com",
						CreatedAt: &now,
						UpdatedAt: &now,
					},
				}
				mockUseCase.EXPECT().ListUsers(gomock.Any()).Return(users, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: []handler.UserResponse{
				{
					ID:    1,
					Name:  "User 1",
					Email: "user1@example.com",
				},
				{
					ID:    2,
					Name:  "User 2",
					Email: "user2@example.com",
				},
			},
		},
		{
			name: "Error",
			mockSetup: func(mockUseCase *mockusecase.MockUserUseCase) {
				mockUseCase.EXPECT().ListUsers(gomock.Any()).Return(nil, errors.New("database error"))
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
			req, _ := http.NewRequest(http.MethodGet, "/users", nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response []handler.UserResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				// Compare only ID, Name, and Email fields since CreatedAt and UpdatedAt will be different
				for i, expectedUser := range tt.expectedBody {
					assert.Equal(t, expectedUser.ID, response[i].ID)
					assert.Equal(t, expectedUser.Name, response[i].Name)
					assert.Equal(t, expectedUser.Email, response[i].Email)
				}
			}
		})
	}
}

func TestUserHandler_GetUser(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupUserTest(t)
	r.GET("/users/:id", h.GetUser)

	// Test cases
	tests := []struct {
		name           string
		userID         string
		mockSetup      func(mockUseCase *mockusecase.MockUserUseCase)
		expectedStatus int
		expectedBody   *handler.UserResponse
	}{
		{
			name:   "Success",
			userID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockUserUseCase) {
				now := time.Now()
				user := &model.User{
					ID:        1,
					Name:      "User 1",
					Email:     "user1@example.com",
					CreatedAt: &now,
					UpdatedAt: &now,
				}
				mockUseCase.EXPECT().GetUserByID(gomock.Any(), int32(1)).Return(user, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.UserResponse{
				ID:    1,
				Name:  "User 1",
				Email: "user1@example.com",
			},
		},
		{
			name:           "Invalid ID",
			userID:         "invalid",
			mockSetup:      func(mockUseCase *mockusecase.MockUserUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name:   "User Not Found",
			userID: "999",
			mockSetup: func(mockUseCase *mockusecase.MockUserUseCase) {
				mockUseCase.EXPECT().GetUserByID(gomock.Any(), int32(999)).Return(nil, errors.New("user not found"))
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
			req, _ := http.NewRequest(http.MethodGet, "/users/"+tt.userID, nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response handler.UserResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				// Compare only ID, Name, and Email fields
				assert.Equal(t, tt.expectedBody.ID, response.ID)
				assert.Equal(t, tt.expectedBody.Name, response.Name)
				assert.Equal(t, tt.expectedBody.Email, response.Email)
			}
		})
	}
}

func TestUserHandler_CreateUser(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupUserTest(t)
	r.POST("/users", h.CreateUser)

	// Test cases
	tests := []struct {
		name           string
		requestBody    handler.CreateUserRequest
		mockSetup      func(mockUseCase *mockusecase.MockUserUseCase)
		expectedStatus int
		expectedBody   *handler.UserResponse
	}{
		{
			name: "Success",
			requestBody: handler.CreateUserRequest{
				Name:     "New User",
				Email:    "newuser@example.com",
				Password: "password123",
			},
			mockSetup: func(mockUseCase *mockusecase.MockUserUseCase) {
				mockUseCase.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ interface{}, user *model.User) error {
						// Verify user properties
						if user.Name != "New User" || user.Email != "newuser@example.com" || user.Password != "password123" {
							t.Errorf("Expected user with name 'New User', email 'newuser@example.com', and password 'password123', got name '%s', email '%s', and password '%s'",
								user.Name, user.Email, user.Password)
						}
						// Set ID to simulate database insertion
						user.ID = 1
						return nil
					})
			},
			expectedStatus: http.StatusCreated,
			expectedBody: &handler.UserResponse{
				ID:    1,
				Name:  "New User",
				Email: "newuser@example.com",
			},
		},
		{
			name: "Invalid Request - Missing Name",
			requestBody: handler.CreateUserRequest{
				Email:    "newuser@example.com",
				Password: "password123",
			},
			mockSetup:      func(mockUseCase *mockusecase.MockUserUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name: "Invalid Request - Invalid Email",
			requestBody: handler.CreateUserRequest{
				Name:     "New User",
				Email:    "invalid-email",
				Password: "password123",
			},
			mockSetup:      func(mockUseCase *mockusecase.MockUserUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name: "Invalid Request - Short Password",
			requestBody: handler.CreateUserRequest{
				Name:     "New User",
				Email:    "newuser@example.com",
				Password: "short",
			},
			mockSetup:      func(mockUseCase *mockusecase.MockUserUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name: "Database Error",
			requestBody: handler.CreateUserRequest{
				Name:     "New User",
				Email:    "newuser@example.com",
				Password: "password123",
			},
			mockSetup: func(mockUseCase *mockusecase.MockUserUseCase) {
				mockUseCase.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
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
			req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusCreated {
				var response handler.UserResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				// Compare only ID, Name, and Email fields
				assert.Equal(t, tt.expectedBody.ID, response.ID)
				assert.Equal(t, tt.expectedBody.Name, response.Name)
				assert.Equal(t, tt.expectedBody.Email, response.Email)
			}
		})
	}
}

func TestUserHandler_UpdateUser(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupUserTest(t)
	r.PUT("/users/:id", h.UpdateUser)

	// Test cases
	tests := []struct {
		name           string
		userID         string
		requestBody    handler.UpdateUserRequest
		mockSetup      func(mockUseCase *mockusecase.MockUserUseCase)
		expectedStatus int
		expectedBody   *handler.UserResponse
	}{
		{
			name:   "Success",
			userID: "1",
			requestBody: handler.UpdateUserRequest{
				Name:     "Updated User",
				Email:    "updated@example.com",
				Password: "newpassword123",
			},
			mockSetup: func(mockUseCase *mockusecase.MockUserUseCase) {
				now := time.Now()
				existingUser := &model.User{
					ID:        1,
					Name:      "Original User",
					Email:     "original@example.com",
					Password:  "oldpassword",
					CreatedAt: &now,
					UpdatedAt: &now,
				}

				mockUseCase.EXPECT().GetUserByID(gomock.Any(), int32(1)).Return(existingUser, nil)

				mockUseCase.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ interface{}, user *model.User) error {
						// Verify user properties
						if user.ID != 1 || user.Name != "Updated User" || user.Email != "updated@example.com" || user.Password != "newpassword123" {
							t.Errorf("Expected user with ID 1, name 'Updated User', email 'updated@example.com', and password 'newpassword123', got ID %d, name '%s', email '%s', and password '%s'",
								user.ID, user.Name, user.Email, user.Password)
						}
						return nil
					})
			},
			expectedStatus: http.StatusOK,
			expectedBody: &handler.UserResponse{
				ID:    1,
				Name:  "Updated User",
				Email: "updated@example.com",
			},
		},
		{
			name:           "Invalid ID",
			userID:         "invalid",
			requestBody:    handler.UpdateUserRequest{},
			mockSetup:      func(mockUseCase *mockusecase.MockUserUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name:   "Invalid Request - Missing Name",
			userID: "1",
			requestBody: handler.UpdateUserRequest{
				Email:    "updated@example.com",
				Password: "newpassword123",
			},
			mockSetup:      func(mockUseCase *mockusecase.MockUserUseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name:   "User Not Found",
			userID: "999",
			requestBody: handler.UpdateUserRequest{
				Name:     "Updated User",
				Email:    "updated@example.com",
				Password: "newpassword123",
			},
			mockSetup: func(mockUseCase *mockusecase.MockUserUseCase) {
				mockUseCase.EXPECT().GetUserByID(gomock.Any(), int32(999)).Return(nil, errors.New("user not found"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   nil,
		},
		{
			name:   "Database Error",
			userID: "1",
			requestBody: handler.UpdateUserRequest{
				Name:     "Updated User",
				Email:    "updated@example.com",
				Password: "newpassword123",
			},
			mockSetup: func(mockUseCase *mockusecase.MockUserUseCase) {
				now := time.Now()
				existingUser := &model.User{
					ID:        1,
					Name:      "Original User",
					Email:     "original@example.com",
					Password:  "oldpassword",
					CreatedAt: &now,
					UpdatedAt: &now,
				}

				mockUseCase.EXPECT().GetUserByID(gomock.Any(), int32(1)).Return(existingUser, nil)

				mockUseCase.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
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
			req, _ := http.NewRequest(http.MethodPut, "/users/"+tt.userID, bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response handler.UserResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				// Compare only ID, Name, and Email fields
				assert.Equal(t, tt.expectedBody.ID, response.ID)
				assert.Equal(t, tt.expectedBody.Name, response.Name)
				assert.Equal(t, tt.expectedBody.Email, response.Email)
			}
		})
	}
}

func TestUserHandler_DeleteUser(t *testing.T) {
	// Setup
	r, mockUseCase, h := setupUserTest(t)
	r.DELETE("/users/:id", h.DeleteUser)

	// Test cases
	tests := []struct {
		name           string
		userID         string
		mockSetup      func(mockUseCase *mockusecase.MockUserUseCase)
		expectedStatus int
	}{
		{
			name:   "Success",
			userID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockUserUseCase) {
				mockUseCase.EXPECT().DeleteUser(gomock.Any(), int32(1)).Return(nil)
			},
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "Invalid ID",
			userID:         "invalid",
			mockSetup:      func(mockUseCase *mockusecase.MockUserUseCase) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:   "Database Error",
			userID: "1",
			mockSetup: func(mockUseCase *mockusecase.MockUserUseCase) {
				mockUseCase.EXPECT().DeleteUser(gomock.Any(), int32(1)).Return(errors.New("database error"))
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
			req, _ := http.NewRequest(http.MethodDelete, "/users/"+tt.userID, nil)
			r.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
