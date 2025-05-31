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

func setupDamageLevelTest(t *testing.T) (*mockdatastore.MockDamageLevelRepository, usecase.DamageLevelUseCase) {
	ctrl := gomock.NewController(t)
	mockRepo := mockdatastore.NewMockDamageLevelRepository(ctrl)
	useCase := usecase.NewDamageLevelUseCase(mockRepo)
	return mockRepo, useCase
}

func TestDamageLevelUseCase_ListDamageLevels(t *testing.T) {
	// Setup
	mockRepo, useCase := setupDamageLevelTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		mockSetup     func(mockRepo *mockdatastore.MockDamageLevelRepository)
		expectedError bool
		expectedLen   int
	}{
		{
			name: "Success",
			mockSetup: func(mockRepo *mockdatastore.MockDamageLevelRepository) {
				description1 := "軽微な被害の説明"
				description2 := "中程度の被害の説明"
				now := time.Now()

				damageLevels := []*model.DamageLevel{
					{
						ID:          1,
						Name:        "軽微",
						Description: &description1,
						CreatedAt:   now,
						UpdatedAt:   now,
					},
					{
						ID:          2,
						Name:        "中程度",
						Description: &description2,
						CreatedAt:   now,
						UpdatedAt:   now,
					},
				}
				mockRepo.EXPECT().Find(gomock.Any()).Return(damageLevels, nil)
			},
			expectedError: false,
			expectedLen:   2,
		},
		{
			name: "Error",
			mockSetup: func(mockRepo *mockdatastore.MockDamageLevelRepository) {
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
			damageLevels, err := useCase.ListDamageLevels(ctx)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, damageLevels)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, damageLevels)
				assert.Equal(t, tt.expectedLen, len(damageLevels))
			}
		})
	}
}

func TestDamageLevelUseCase_GetDamageLevelByID(t *testing.T) {
	// Setup
	mockRepo, useCase := setupDamageLevelTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		id            int32
		mockSetup     func(mockRepo *mockdatastore.MockDamageLevelRepository)
		expectedError bool
	}{
		{
			name: "Success",
			id:   1,
			mockSetup: func(mockRepo *mockdatastore.MockDamageLevelRepository) {
				description := "軽微な被害の説明"
				now := time.Now()

				damageLevel := &model.DamageLevel{
					ID:          1,
					Name:        "軽微",
					Description: &description,
					CreatedAt:   now,
					UpdatedAt:   now,
				}
				mockRepo.EXPECT().FindByID(gomock.Any(), int32(1)).Return(damageLevel, nil)
			},
			expectedError: false,
		},
		{
			name: "Not Found",
			id:   999,
			mockSetup: func(mockRepo *mockdatastore.MockDamageLevelRepository) {
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
			damageLevel, err := useCase.GetDamageLevelByID(ctx, tt.id)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, damageLevel)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, damageLevel)
				assert.Equal(t, tt.id, damageLevel.ID)
			}
		})
	}
}

func TestDamageLevelUseCase_CreateDamageLevel(t *testing.T) {
	// Setup
	mockRepo, useCase := setupDamageLevelTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		damageLevel   *model.DamageLevel
		mockSetup     func(mockRepo *mockdatastore.MockDamageLevelRepository)
		expectedError bool
	}{
		{
			name: "Success",
			damageLevel: func() *model.DamageLevel {
				description := "新しい被害程度の説明"
				return &model.DamageLevel{
					Name:        "新規被害程度",
					Description: &description,
				}
			}(),
			mockSetup: func(mockRepo *mockdatastore.MockDamageLevelRepository) {
				mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			damageLevel: func() *model.DamageLevel {
				description := "新しい被害程度の説明"
				return &model.DamageLevel{
					Name:        "新規被害程度",
					Description: &description,
				}
			}(),
			mockSetup: func(mockRepo *mockdatastore.MockDamageLevelRepository) {
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
			err := useCase.CreateDamageLevel(ctx, tt.damageLevel)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDamageLevelUseCase_UpdateDamageLevel(t *testing.T) {
	// Setup
	mockRepo, useCase := setupDamageLevelTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		damageLevel   *model.DamageLevel
		mockSetup     func(mockRepo *mockdatastore.MockDamageLevelRepository)
		expectedError bool
	}{
		{
			name: "Success",
			damageLevel: func() *model.DamageLevel {
				description := "更新された被害程度の説明"
				return &model.DamageLevel{
					ID:          1,
					Name:        "更新された被害程度",
					Description: &description,
				}
			}(),
			mockSetup: func(mockRepo *mockdatastore.MockDamageLevelRepository) {
				mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			damageLevel: func() *model.DamageLevel {
				description := "更新された被害程度の説明"
				return &model.DamageLevel{
					ID:          1,
					Name:        "更新された被害程度",
					Description: &description,
				}
			}(),
			mockSetup: func(mockRepo *mockdatastore.MockDamageLevelRepository) {
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
			err := useCase.UpdateDamageLevel(ctx, tt.damageLevel)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDamageLevelUseCase_DeleteDamageLevel(t *testing.T) {
	// Setup
	mockRepo, useCase := setupDamageLevelTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		id            int32
		mockSetup     func(mockRepo *mockdatastore.MockDamageLevelRepository)
		expectedError bool
	}{
		{
			name: "Success",
			id:   1,
			mockSetup: func(mockRepo *mockdatastore.MockDamageLevelRepository) {
				mockRepo.EXPECT().Delete(gomock.Any(), int32(1)).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			id:   1,
			mockSetup: func(mockRepo *mockdatastore.MockDamageLevelRepository) {
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
			err := useCase.DeleteDamageLevel(ctx, tt.id)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
