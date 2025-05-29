package datastore

import (
	"context"
	"time"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/domain/query"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
)

// DisasterSearchParams contains the search parameters for disasters
type DisasterSearchParams struct {
	Name         string
	DisasterType string
	Status       string
	PrefectureID int32
	StartDate    time.Time
	EndDate      time.Time
}

type DisasterRepository interface {
	Find(ctx context.Context, params *DisasterSearchParams) ([]*model.Disaster, error)
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

func (r *disasterRepository) Find(ctx context.Context, params *DisasterSearchParams) ([]*model.Disaster, error) {
	q := r.query.WithContext(ctx).Disaster.
		Preload(r.query.Disaster.Prefecture)

	// Apply filters if provided
	if params != nil {
		if params.Name != "" {
			q = q.Where(r.query.Disaster.Name.Like("%" + params.Name + "%"))
		}
		if params.DisasterType != "" {
			q = q.Where(r.query.Disaster.DisasterType.Eq(params.DisasterType))
		}
		if params.Status != "" {
			q = q.Where(r.query.Disaster.Status.Eq(params.Status))
		}
		if params.PrefectureID != 0 {
			q = q.Where(r.query.Disaster.PrefectureID.Eq(params.PrefectureID))
		}
		// Apply date range filter if both start and end dates are provided
		if !params.StartDate.IsZero() && !params.EndDate.IsZero() {
			q = q.Where(r.query.Disaster.OccurredAt.Between(params.StartDate, params.EndDate))
		} else if !params.StartDate.IsZero() {
			q = q.Where(r.query.Disaster.OccurredAt.Gte(params.StartDate))
		} else if !params.EndDate.IsZero() {
			q = q.Where(r.query.Disaster.OccurredAt.Lte(params.EndDate))
		}
	}

	disasters, err := q.Find()
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
