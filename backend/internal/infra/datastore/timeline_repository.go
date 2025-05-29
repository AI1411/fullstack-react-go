package datastore

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
)

type TimelineRepository interface {
	FindByDisasterID(ctx context.Context, disasterID string) ([]*model.Timeline, error)
}

type timelineRepository struct {
	client db.Client
}

func NewTimelineRepository(
	ctx context.Context,
	client db.Client,
) TimelineRepository {
	return &timelineRepository{
		client: client,
	}
}

func (r *timelineRepository) FindByDisasterID(ctx context.Context, disasterID string) ([]*model.Timeline, error) {
	var timelines []*model.Timeline

	err := r.client.Conn(ctx).
		Where("disaster_id = ?", disasterID).
		Order("event_time").
		Find(&timelines).Error

	if err != nil {
		return nil, err
	}

	return timelines, nil
}
