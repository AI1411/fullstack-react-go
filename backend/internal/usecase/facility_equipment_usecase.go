//go:generate mockgen -source=facility_equipment_usecase.go -destination=../../tests/mock/usecase/facility_equipment_usecase.mock.go
package usecase

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	domain "github.com/AI1411/fullstack-react-go/internal/domain/repository"
)

type FacilityEquipmentUseCase interface {
	ListFacilityEquipments(ctx context.Context) ([]*model.FacilityEquipment, error)
	GetFacilityEquipmentByID(ctx context.Context, id int32) (*model.FacilityEquipment, error)
	CreateFacilityEquipment(ctx context.Context, facilityEquipment *model.FacilityEquipment) error
	UpdateFacilityEquipment(ctx context.Context, facilityEquipment *model.FacilityEquipment) error
	DeleteFacilityEquipment(ctx context.Context, id int32) error
}

type facilityEquipmentUseCase struct {
	facilityEquipmentRepository domain.FacilityEquipmentRepository
}

func NewFacilityEquipmentUseCase(
	facilityEquipmentRepository domain.FacilityEquipmentRepository,
) FacilityEquipmentUseCase {
	return &facilityEquipmentUseCase{
		facilityEquipmentRepository: facilityEquipmentRepository,
	}
}

func (u *facilityEquipmentUseCase) ListFacilityEquipments(ctx context.Context) ([]*model.FacilityEquipment, error) {
	facilityEquipments, err := u.facilityEquipmentRepository.Find(ctx)
	if err != nil {
		return nil, err
	}

	return facilityEquipments, nil
}

func (u *facilityEquipmentUseCase) GetFacilityEquipmentByID(ctx context.Context, id int32) (*model.FacilityEquipment, error) {
	facilityEquipment, err := u.facilityEquipmentRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return facilityEquipment, nil
}

func (u *facilityEquipmentUseCase) CreateFacilityEquipment(ctx context.Context, facilityEquipment *model.FacilityEquipment) error {
	return u.facilityEquipmentRepository.Create(ctx, facilityEquipment)
}

func (u *facilityEquipmentUseCase) UpdateFacilityEquipment(ctx context.Context, facilityEquipment *model.FacilityEquipment) error {
	return u.facilityEquipmentRepository.Update(ctx, facilityEquipment)
}

func (u *facilityEquipmentUseCase) DeleteFacilityEquipment(ctx context.Context, id int32) error {
	return u.facilityEquipmentRepository.Delete(ctx, id)
}
