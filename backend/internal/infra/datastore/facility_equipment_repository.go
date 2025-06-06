//go:generate mockgen -source=facility_equipment_repository.go -destination=../../../tests/mock/datastore/facility_equipment_repository.mock.go
package datastore

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/domain/query"
	domain "github.com/AI1411/fullstack-react-go/internal/domain/repository"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
)

type facilityEquipmentRepository struct {
	client db.Client
	query  *query.Query
}

func NewFacilityEquipmentRepository(
	ctx context.Context,
	client db.Client,
) domain.FacilityEquipmentRepository {
	return &facilityEquipmentRepository{
		client: client,
		query:  query.Use(client.Conn(ctx)),
	}
}

func (r *facilityEquipmentRepository) Find(ctx context.Context) ([]*model.FacilityEquipment, error) {
	var facilityEquipments []*model.FacilityEquipment
	if err := r.client.Conn(ctx).Preload("FacilityType").Find(&facilityEquipments).Error; err != nil {
		return nil, err
	}

	return facilityEquipments, nil
}

func (r *facilityEquipmentRepository) FindByID(ctx context.Context, id int32) (*model.FacilityEquipment, error) {
	var facilityEquipment model.FacilityEquipment
	if err := r.client.Conn(ctx).Preload("FacilityType").Where("id = ?", id).First(&facilityEquipment).Error; err != nil {
		return nil, err
	}

	return &facilityEquipment, nil
}

func (r *facilityEquipmentRepository) Create(ctx context.Context, facilityEquipment *model.FacilityEquipment) error {
	return r.client.Conn(ctx).Create(facilityEquipment).Error
}

func (r *facilityEquipmentRepository) Update(ctx context.Context, facilityEquipment *model.FacilityEquipment) error {
	return r.client.Conn(ctx).Save(facilityEquipment).Error
}

func (r *facilityEquipmentRepository) Delete(ctx context.Context, id int32) error {
	return r.client.Conn(ctx).Delete(&model.FacilityEquipment{}, id).Error
}
