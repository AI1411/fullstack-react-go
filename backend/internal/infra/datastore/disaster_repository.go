package datastore

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/domain/query"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
)

type DisasterRepository interface {
	Find(ctx context.Context) ([]*model.Disaster, error)
	FindByID(ctx context.Context, id string) (*model.Disaster, error)
	Create(ctx context.Context, disaster *model.Disaster) error
	Update(ctx context.Context, disaster *model.Disaster) error
	Delete(ctx context.Context, id string) error
}

type disasterRepository struct {
	client db.Client
	query  *query.Query
}

func NewDisasterRepository(
	ctx context.Context,
	client db.Client,
) DisasterRepository {
	return &disasterRepository{
		client: client,
		query:  query.Use(client.Conn(ctx)),
	}
}

func (r *disasterRepository) Find(ctx context.Context) ([]*model.Disaster, error) {
	disasters, err := r.query.WithContext(ctx).Disaster.
		Preload(r.query.Disaster.Prefecture).
		Find()
	if err != nil {
		return nil, err
	}

	return disasters, nil
}

func (r *disasterRepository) FindByID(ctx context.Context, id string) (*model.Disaster, error) {
	disaster, err := r.query.WithContext(ctx).
		Disaster.
		Where(r.query.Disaster.ID.Eq(id)).
		Preload(r.query.Disaster.Prefecture).
		First()
	if err != nil {
		return nil, err
	}

	return disaster, nil
}

func (r *disasterRepository) Create(ctx context.Context, disaster *model.Disaster) error {
	return r.query.WithContext(ctx).Disaster.Create(disaster)
}

func (r *disasterRepository) Update(ctx context.Context, disaster *model.Disaster) error {
	_, err := r.query.WithContext(ctx).Disaster.Where(r.query.Disaster.ID.Eq(disaster.ID)).Updates(disaster)
	return err
}

func (r *disasterRepository) Delete(ctx context.Context, id string) error {
	_, err := r.query.WithContext(ctx).Disaster.Where(r.query.Disaster.ID.Eq(id)).Delete()

	return err
}
