//go:generate mockgen -source=damage_level_usecase.go -destination=../../tests/mock/usecase/damage_level_usecase.mock.go
package usecase

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/datastore"
)

type DamageLevelUseCase interface {
	ListDamageLevels(ctx context.Context) ([]*model.DamageLevel, error)
	GetDamageLevelByID(ctx context.Context, id int32) (*model.DamageLevel, error)
	CreateDamageLevel(ctx context.Context, damageLevel *model.DamageLevel) error
	UpdateDamageLevel(ctx context.Context, damageLevel *model.DamageLevel) error
	DeleteDamageLevel(ctx context.Context, id int32) error
}

type damageLevelUseCase struct {
	damageLevelRepository datastore.DamageLevelRepository
}

func NewDamageLevelUseCase(
	damageLevelRepository datastore.DamageLevelRepository,
) DamageLevelUseCase {
	return &damageLevelUseCase{
		damageLevelRepository: damageLevelRepository,
	}
}

func (u *damageLevelUseCase) ListDamageLevels(ctx context.Context) ([]*model.DamageLevel, error) {
	damageLevels, err := u.damageLevelRepository.Find(ctx)
	if err != nil {
		return nil, err
	}

	return damageLevels, nil
}

func (u *damageLevelUseCase) GetDamageLevelByID(ctx context.Context, id int32) (*model.DamageLevel, error) {
	damageLevel, err := u.damageLevelRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return damageLevel, nil
}

func (u *damageLevelUseCase) CreateDamageLevel(ctx context.Context, damageLevel *model.DamageLevel) error {
	return u.damageLevelRepository.Create(ctx, damageLevel)
}

func (u *damageLevelUseCase) UpdateDamageLevel(ctx context.Context, damageLevel *model.DamageLevel) error {
	return u.damageLevelRepository.Update(ctx, damageLevel)
}

func (u *damageLevelUseCase) DeleteDamageLevel(ctx context.Context, id int32) error {
	return u.damageLevelRepository.Delete(ctx, id)
}
