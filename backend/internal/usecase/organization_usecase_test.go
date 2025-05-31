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

func setupOrganizationTest(t *testing.T) (*mockdatastore.MockOrganizationRepository, usecase.OrganizationUseCase) {
	ctrl := gomock.NewController(t)
	mockRepo := mockdatastore.NewMockOrganizationRepository(ctrl)
	useCase := usecase.NewOrganizationUseCase(mockRepo)
	return mockRepo, useCase
}

func TestOrganizationUseCase_ListOrganizations(t *testing.T) {
	// Setup
	mockRepo, useCase := setupOrganizationTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		mockSetup     func(mockRepo *mockdatastore.MockOrganizationRepository)
		expectedError bool
		expectedLen   int
	}{
		{
			name: "Success",
			mockSetup: func(mockRepo *mockdatastore.MockOrganizationRepository) {
				now := time.Now()
				prefectureID1 := int32(13)
				description1 := "東京都の行政組織"

				prefectureID2 := int32(27)
				description2 := "大阪府の行政組織"
				parentID2 := int32(1)

				organizations := []*model.Organization{
					{
						ID:           1,
						Name:         "東京都庁",
						Type:         "都道府県",
						PrefectureID: &prefectureID1,
						ParentID:     nil,
						Description:  &description1,
						CreatedAt:    now,
						UpdatedAt:    now,
					},
					{
						ID:           2,
						Name:         "大阪府庁",
						Type:         "都道府県",
						PrefectureID: &prefectureID2,
						ParentID:     &parentID2,
						Description:  &description2,
						CreatedAt:    now,
						UpdatedAt:    now,
					},
				}
				mockRepo.EXPECT().Find(gomock.Any()).Return(organizations, nil)
			},
			expectedError: false,
			expectedLen:   2,
		},
		{
			name: "Error",
			mockSetup: func(mockRepo *mockdatastore.MockOrganizationRepository) {
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
			organizations, err := useCase.ListOrganizations(ctx)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, organizations)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, organizations)
				assert.Equal(t, tt.expectedLen, len(organizations))
			}
		})
	}
}

func TestOrganizationUseCase_GetOrganizationByID(t *testing.T) {
	// Setup
	mockRepo, useCase := setupOrganizationTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		id            int32
		mockSetup     func(mockRepo *mockdatastore.MockOrganizationRepository)
		expectedError bool
	}{
		{
			name: "Success",
			id:   1,
			mockSetup: func(mockRepo *mockdatastore.MockOrganizationRepository) {
				now := time.Now()
				prefectureID := int32(13)
				description := "東京都の行政組織"

				organization := &model.Organization{
					ID:           1,
					Name:         "東京都庁",
					Type:         "都道府県",
					PrefectureID: &prefectureID,
					ParentID:     nil,
					Description:  &description,
					CreatedAt:    now,
					UpdatedAt:    now,
				}
				mockRepo.EXPECT().FindByID(gomock.Any(), int32(1)).Return(organization, nil)
			},
			expectedError: false,
		},
		{
			name: "Not Found",
			id:   999,
			mockSetup: func(mockRepo *mockdatastore.MockOrganizationRepository) {
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
			organization, err := useCase.GetOrganizationByID(ctx, tt.id)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, organization)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, organization)
				assert.Equal(t, tt.id, organization.ID)
			}
		})
	}
}

func TestOrganizationUseCase_CreateOrganization(t *testing.T) {
	// Setup
	mockRepo, useCase := setupOrganizationTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		organization  *model.Organization
		mockSetup     func(mockRepo *mockdatastore.MockOrganizationRepository)
		expectedError bool
	}{
		{
			name: "Success",
			organization: func() *model.Organization {
				prefectureID := int32(14)
				description := "神奈川県の行政組織"

				return &model.Organization{
					Name:         "神奈川県庁",
					Type:         "都道府県",
					PrefectureID: &prefectureID,
					ParentID:     nil,
					Description:  &description,
				}
			}(),
			mockSetup: func(mockRepo *mockdatastore.MockOrganizationRepository) {
				mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			organization: func() *model.Organization {
				return &model.Organization{
					Name: "神奈川県庁",
					Type: "都道府県",
				}
			}(),
			mockSetup: func(mockRepo *mockdatastore.MockOrganizationRepository) {
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
			err := useCase.CreateOrganization(ctx, tt.organization)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestOrganizationUseCase_UpdateOrganization(t *testing.T) {
	// Setup
	mockRepo, useCase := setupOrganizationTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		organization  *model.Organization
		mockSetup     func(mockRepo *mockdatastore.MockOrganizationRepository)
		expectedError bool
	}{
		{
			name: "Success",
			organization: func() *model.Organization {
				prefectureID := int32(13)
				description := "東京都の行政組織（更新）"

				return &model.Organization{
					ID:           1,
					Name:         "東京都庁（更新）",
					Type:         "都道府県",
					PrefectureID: &prefectureID,
					ParentID:     nil,
					Description:  &description,
				}
			}(),
			mockSetup: func(mockRepo *mockdatastore.MockOrganizationRepository) {
				mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			organization: func() *model.Organization {
				return &model.Organization{
					ID:   1,
					Name: "東京都庁（更新）",
					Type: "都道府県",
				}
			}(),
			mockSetup: func(mockRepo *mockdatastore.MockOrganizationRepository) {
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
			err := useCase.UpdateOrganization(ctx, tt.organization)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestOrganizationUseCase_DeleteOrganization(t *testing.T) {
	// Setup
	mockRepo, useCase := setupOrganizationTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		id            int32
		mockSetup     func(mockRepo *mockdatastore.MockOrganizationRepository)
		expectedError bool
	}{
		{
			name: "Success",
			id:   1,
			mockSetup: func(mockRepo *mockdatastore.MockOrganizationRepository) {
				mockRepo.EXPECT().Delete(gomock.Any(), int32(1)).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			id:   1,
			mockSetup: func(mockRepo *mockdatastore.MockOrganizationRepository) {
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
			err := useCase.DeleteOrganization(ctx, tt.id)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
