package datastore

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
)

type DamageLevelRepository interface {
	Find(ctx context.Context) ([]*model.DamageLevel, error)
	FindByID(ctx context.Context, id int32) (*model.DamageLevel, error)
}

type damageLevelRepository struct {
	client db.Client
}

func NewDamageLevelRepository(
	ctx context.Context,
	client db.Client,
) DamageLevelRepository {
	return &damageLevelRepository{
		client: client,
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
