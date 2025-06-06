//go:generate mockgen -source=facility_equipment.go -destination=../../../tests/mock/domain/facility_equipment.mock.go
package domain

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
)

type FacilityEquipmentRepository interface {
	Find(ctx context.Context) ([]*model.FacilityEquipment, error)
	FindByID(ctx context.Context, id int32) (*model.FacilityEquipment, error)
	Create(ctx context.Context, facilityEquipment *model.FacilityEquipment) error
	Update(ctx context.Context, facilityEquipment *model.FacilityEquipment) error
	Delete(ctx context.Context, id int32) error
}
