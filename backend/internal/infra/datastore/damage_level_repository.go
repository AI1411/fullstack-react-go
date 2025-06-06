//go:generate mockgen -source=damage_level_repository.go -destination=../../../tests/mock/datastore/damage_level_repository.mock.go
package datastore

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/domain/query"
	domain "github.com/AI1411/fullstack-react-go/internal/domain/repository"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
)

type damageLevelRepository struct {
	client db.Client
	query  *query.Query
}

func NewDamageLevelRepository(
	ctx context.Context,
	client db.Client,
) domain.DamageLevelRepository {
	return &damageLevelRepository{
		client: client,
		query:  query.Use(client.Conn(ctx)),
	}
}

func (r *damageLevelRepository) Find(ctx context.Context) ([]*model.DamageLevel, error) {
	var damageLevels []*model.DamageLevel
	if err := r.client.Conn(ctx).Find(&damageLevels).Error; err != nil {
		return nil, err
	}

	return damageLevels, nil
}

func (r *damageLevelRepository) FindByID(ctx context.Context, id int32) (*model.DamageLevel, error) {
	var damageLevel model.DamageLevel
	if err := r.client.Conn(ctx).Where("id = ?", id).First(&damageLevel).Error; err != nil {
		return nil, err
	}

	return &damageLevel, nil
}

func (r *damageLevelRepository) Create(ctx context.Context, damageLevel *model.DamageLevel) error {
	return r.client.Conn(ctx).Create(damageLevel).Error
}

func (r *damageLevelRepository) Update(ctx context.Context, damageLevel *model.DamageLevel) error {
	return r.client.Conn(ctx).Save(damageLevel).Error
}

func (r *damageLevelRepository) Delete(ctx context.Context, id int32) error {
	return r.client.Conn(ctx).Delete(&model.DamageLevel{}, id).Error
}
