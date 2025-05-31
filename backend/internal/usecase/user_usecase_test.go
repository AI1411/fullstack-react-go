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

func setupUserTest(t *testing.T) (*mockdatastore.MockUserRepository, usecase.UserUseCase) {
	ctrl := gomock.NewController(t)
	mockRepo := mockdatastore.NewMockUserRepository(ctrl)
	useCase := usecase.NewUserUseCase(mockRepo)
	return mockRepo, useCase
}

func TestUserUseCase_ListUsers(t *testing.T) {
	// Setup
	mockRepo, useCase := setupUserTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		mockSetup     func(mockRepo *mockdatastore.MockUserRepository)
		expectedError bool
		expectedLen   int
	}{
		{
			name: "Success",
			mockSetup: func(mockRepo *mockdatastore.MockUserRepository) {
				now := time.Now()

				users := []*model.User{
					{
						ID:        1,
						Name:      "山田太郎",
						Email:     "yamada@example.com",
						Password:  "hashed_password_1",
						CreatedAt: &now,
						UpdatedAt: &now,
					},
					{
						ID:        2,
						Name:      "佐藤花子",
						Email:     "sato@example.com",
						Password:  "hashed_password_2",
						CreatedAt: &now,
						UpdatedAt: &now,
					},
				}
				mockRepo.EXPECT().Find(gomock.Any()).Return(users, nil)
			},
			expectedError: false,
			expectedLen:   2,
		},
		{
			name: "Error",
			mockSetup: func(mockRepo *mockdatastore.MockUserRepository) {
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
			users, err := useCase.ListUsers(ctx)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, users)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, users)
				assert.Equal(t, tt.expectedLen, len(users))
			}
		})
	}
}

func TestUserUseCase_GetUserByID(t *testing.T) {
	// Setup
	mockRepo, useCase := setupUserTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		id            int32
		mockSetup     func(mockRepo *mockdatastore.MockUserRepository)
		expectedError bool
	}{
		{
			name: "Success",
			id:   1,
			mockSetup: func(mockRepo *mockdatastore.MockUserRepository) {
				now := time.Now()

				user := &model.User{
					ID:        1,
					Name:      "山田太郎",
					Email:     "yamada@example.com",
					Password:  "hashed_password_1",
					CreatedAt: &now,
					UpdatedAt: &now,
				}
				mockRepo.EXPECT().FindByID(gomock.Any(), int32(1)).Return(user, nil)
			},
			expectedError: false,
		},
		{
			name: "Not Found",
			id:   999,
			mockSetup: func(mockRepo *mockdatastore.MockUserRepository) {
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
			user, err := useCase.GetUserByID(ctx, tt.id)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.id, user.ID)
			}
		})
	}
}

func TestUserUseCase_CreateUser(t *testing.T) {
	// Setup
	mockRepo, useCase := setupUserTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		user          *model.User
		mockSetup     func(mockRepo *mockdatastore.MockUserRepository)
		expectedError bool
	}{
		{
			name: "Success",
			user: &model.User{
				Name:     "鈴木一郎",
				Email:    "suzuki@example.com",
				Password: "password123",
			},
			mockSetup: func(mockRepo *mockdatastore.MockUserRepository) {
				mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			user: &model.User{
				Name:     "鈴木一郎",
				Email:    "suzuki@example.com",
				Password: "password123",
			},
			mockSetup: func(mockRepo *mockdatastore.MockUserRepository) {
				mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(errors.New("database error"))
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
			err := useCase.CreateUser(ctx, tt.user)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserUseCase_UpdateUser(t *testing.T) {
	// Setup
	mockRepo, useCase := setupUserTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		user          *model.User
		mockSetup     func(mockRepo *mockdatastore.MockUserRepository)
		expectedError bool
	}{
		{
			name: "Success",
			user: &model.User{
				ID:       1,
				Name:     "山田太郎（更新）",
				Email:    "yamada_new@example.com",
				Password: "new_password",
			},
			mockSetup: func(mockRepo *mockdatastore.MockUserRepository) {
				mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			user: &model.User{
				ID:       1,
				Name:     "山田太郎（更新）",
				Email:    "yamada_new@example.com",
				Password: "new_password",
			},
			mockSetup: func(mockRepo *mockdatastore.MockUserRepository) {
				mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.New("database error"))
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
			err := useCase.UpdateUser(ctx, tt.user)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserUseCase_DeleteUser(t *testing.T) {
	// Setup
	mockRepo, useCase := setupUserTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		id            int32
		mockSetup     func(mockRepo *mockdatastore.MockUserRepository)
		expectedError bool
	}{
		{
			name: "Success",
			id:   1,
			mockSetup: func(mockRepo *mockdatastore.MockUserRepository) {
				mockRepo.EXPECT().Delete(gomock.Any(), int32(1)).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			id:   1,
			mockSetup: func(mockRepo *mockdatastore.MockUserRepository) {
				mockRepo.EXPECT().Delete(gomock.Any(), int32(1)).Return(errors.New("database error"))
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
			err := useCase.DeleteUser(ctx, tt.id)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
