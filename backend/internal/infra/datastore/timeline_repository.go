//go:generate mockgen -source=timeline_repository.go -destination=../../../tests/mock/datastore/timeline_repository.mock.go
package datastore

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/domain/query"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
)

type TimelineRepository interface {
	FindByDisasterID(ctx context.Context, disasterID string) ([]*model.Timeline, error)
}

type timelineRepository struct {
	client db.Client
	query  *query.Query
}

func NewTimelineRepository(
	ctx context.Context,
	client db.Client,
) TimelineRepository {
	return &timelineRepository{
		client: client,
		query:  query.Use(client.Conn(ctx)),
	}
}

func (r *timelineRepository) FindByDisasterID(ctx context.Context, disasterID string) ([]*model.Timeline, error) {
	timelines, err := r.query.WithContext(ctx).
		Timeline.
		Where(r.query.Timeline.DisasterID.Eq(disasterID)).
		Order(r.query.Timeline.EventTime.Desc()).
		Find()
	if err != nil {
		return nil, err
	}

	return timelines, nil
}
