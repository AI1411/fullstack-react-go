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

func setupFacilityEquipmentTest(t *testing.T) (*mockdatastore.MockFacilityEquipmentRepository, usecase.FacilityEquipmentUseCase) {
	ctrl := gomock.NewController(t)
	mockRepo := mockdatastore.NewMockFacilityEquipmentRepository(ctrl)
	useCase := usecase.NewFacilityEquipmentUseCase(mockRepo)
	return mockRepo, useCase
}

func TestFacilityEquipmentUseCase_ListFacilityEquipments(t *testing.T) {
	// Setup
	mockRepo, useCase := setupFacilityEquipmentTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		mockSetup     func(mockRepo *mockdatastore.MockFacilityEquipmentRepository)
		expectedError bool
		expectedLen   int
	}{
		{
			name: "Success",
			mockSetup: func(mockRepo *mockdatastore.MockFacilityEquipmentRepository) {
				now := time.Now()
				modelNumber1 := "ABC-123"
				manufacturer1 := "メーカーA"
				installationDate1 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
				locationDesc1 := "1階北側"
				notes1 := "定期点検済み"

				modelNumber2 := "XYZ-789"
				manufacturer2 := "メーカーB"
				locationDesc2 := "2階南側"

				facilityEquipments := []*model.FacilityEquipment{
					{
						ID:                  1,
						Name:                "空調設備",
						FacilityTypeID:      1,
						ModelNumber:         &modelNumber1,
						Manufacturer:        &manufacturer1,
						InstallationDate:    &installationDate1,
						Status:              "稼働中",
						LocationDescription: &locationDesc1,
						LocationLatitude:    nil,
						LocationLongitude:   nil,
						Notes:               &notes1,
						CreatedAt:           now,
						UpdatedAt:           now,
					},
					{
						ID:                  2,
						Name:                "照明設備",
						FacilityTypeID:      2,
						ModelNumber:         &modelNumber2,
						Manufacturer:        &manufacturer2,
						InstallationDate:    nil,
						Status:              "メンテナンス中",
						LocationDescription: &locationDesc2,
						LocationLatitude:    nil,
						LocationLongitude:   nil,
						Notes:               nil,
						CreatedAt:           now,
						UpdatedAt:           now,
					},
				}
				mockRepo.EXPECT().Find(gomock.Any()).Return(facilityEquipments, nil)
			},
			expectedError: false,
			expectedLen:   2,
		},
		{
			name: "Error",
			mockSetup: func(mockRepo *mockdatastore.MockFacilityEquipmentRepository) {
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
			facilityEquipments, err := useCase.ListFacilityEquipments(ctx)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, facilityEquipments)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, facilityEquipments)
				assert.Equal(t, tt.expectedLen, len(facilityEquipments))
			}
		})
	}
}

func TestFacilityEquipmentUseCase_GetFacilityEquipmentByID(t *testing.T) {
	// Setup
	mockRepo, useCase := setupFacilityEquipmentTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		id            int32
		mockSetup     func(mockRepo *mockdatastore.MockFacilityEquipmentRepository)
		expectedError bool
	}{
		{
			name: "Success",
			id:   1,
			mockSetup: func(mockRepo *mockdatastore.MockFacilityEquipmentRepository) {
				now := time.Now()
				modelNumber := "ABC-123"
				manufacturer := "メーカーA"
				installationDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
				locationDesc := "1階北側"
				notes := "定期点検済み"

				facilityEquipment := &model.FacilityEquipment{
					ID:                  1,
					Name:                "空調設備",
					FacilityTypeID:      1,
					ModelNumber:         &modelNumber,
					Manufacturer:        &manufacturer,
					InstallationDate:    &installationDate,
					Status:              "稼働中",
					LocationDescription: &locationDesc,
					LocationLatitude:    nil,
					LocationLongitude:   nil,
					Notes:               &notes,
					CreatedAt:           now,
					UpdatedAt:           now,
				}
				mockRepo.EXPECT().FindByID(gomock.Any(), int32(1)).Return(facilityEquipment, nil)
			},
			expectedError: false,
		},
		{
			name: "Not Found",
			id:   999,
			mockSetup: func(mockRepo *mockdatastore.MockFacilityEquipmentRepository) {
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
			facilityEquipment, err := useCase.GetFacilityEquipmentByID(ctx, tt.id)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, facilityEquipment)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, facilityEquipment)
				assert.Equal(t, tt.id, facilityEquipment.ID)
			}
		})
	}
}

func TestFacilityEquipmentUseCase_CreateFacilityEquipment(t *testing.T) {
	// Setup
	mockRepo, useCase := setupFacilityEquipmentTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name              string
		facilityEquipment *model.FacilityEquipment
		mockSetup         func(mockRepo *mockdatastore.MockFacilityEquipmentRepository)
		expectedError     bool
	}{
		{
			name: "Success",
			facilityEquipment: func() *model.FacilityEquipment {
				modelNumber := "DEF-456"
				manufacturer := "メーカーC"
				installationDate := time.Now()
				locationDesc := "3階東側"
				notes := "新規設置"

				return &model.FacilityEquipment{
					Name:                "給水設備",
					FacilityTypeID:      3,
					ModelNumber:         &modelNumber,
					Manufacturer:        &manufacturer,
					InstallationDate:    &installationDate,
					Status:              "稼働中",
					LocationDescription: &locationDesc,
					Notes:               &notes,
				}
			}(),
			mockSetup: func(mockRepo *mockdatastore.MockFacilityEquipmentRepository) {
				mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			facilityEquipment: func() *model.FacilityEquipment {
				return &model.FacilityEquipment{
					Name:           "給水設備",
					FacilityTypeID: 3,
					Status:         "稼働中",
				}
			}(),
			mockSetup: func(mockRepo *mockdatastore.MockFacilityEquipmentRepository) {
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
			err := useCase.CreateFacilityEquipment(ctx, tt.facilityEquipment)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestFacilityEquipmentUseCase_UpdateFacilityEquipment(t *testing.T) {
	// Setup
	mockRepo, useCase := setupFacilityEquipmentTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name              string
		facilityEquipment *model.FacilityEquipment
		mockSetup         func(mockRepo *mockdatastore.MockFacilityEquipmentRepository)
		expectedError     bool
	}{
		{
			name: "Success",
			facilityEquipment: func() *model.FacilityEquipment {
				modelNumber := "ABC-123-Updated"
				manufacturer := "メーカーA"
				locationDesc := "1階北側（更新）"
				notes := "定期点検済み（更新）"

				return &model.FacilityEquipment{
					ID:                  1,
					Name:                "空調設備（更新）",
					FacilityTypeID:      1,
					ModelNumber:         &modelNumber,
					Manufacturer:        &manufacturer,
					Status:              "メンテナンス中",
					LocationDescription: &locationDesc,
					Notes:               &notes,
				}
			}(),
			mockSetup: func(mockRepo *mockdatastore.MockFacilityEquipmentRepository) {
				mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			facilityEquipment: func() *model.FacilityEquipment {
				return &model.FacilityEquipment{
					ID:             1,
					Name:           "空調設備（更新）",
					FacilityTypeID: 1,
					Status:         "メンテナンス中",
				}
			}(),
			mockSetup: func(mockRepo *mockdatastore.MockFacilityEquipmentRepository) {
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
			err := useCase.UpdateFacilityEquipment(ctx, tt.facilityEquipment)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestFacilityEquipmentUseCase_DeleteFacilityEquipment(t *testing.T) {
	// Setup
	mockRepo, useCase := setupFacilityEquipmentTest(t)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name          string
		id            int32
		mockSetup     func(mockRepo *mockdatastore.MockFacilityEquipmentRepository)
		expectedError bool
	}{
		{
			name: "Success",
			id:   1,
			mockSetup: func(mockRepo *mockdatastore.MockFacilityEquipmentRepository) {
				mockRepo.EXPECT().Delete(gomock.Any(), int32(1)).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Error",
			id:   1,
			mockSetup: func(mockRepo *mockdatastore.MockFacilityEquipmentRepository) {
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
			err := useCase.DeleteFacilityEquipment(ctx, tt.id)

			// Check results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
