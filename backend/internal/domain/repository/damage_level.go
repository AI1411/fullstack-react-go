//go:generate mockgen -source=damage_level.go -destination=../../../tests/mock/datastore/damage_level.mock.go
package domain

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
)

type DamageLevelRepository interface {
	Find(ctx context.Context) ([]*model.DamageLevel, error)
	FindByID(ctx context.Context, id int32) (*model.DamageLevel, error)
	Create(ctx context.Context, damageLevel *model.DamageLevel) error
	Update(ctx context.Context, damageLevel *model.DamageLevel) error
	Delete(ctx context.Context, id int32) error
}
