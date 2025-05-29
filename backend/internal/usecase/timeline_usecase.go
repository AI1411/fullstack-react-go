package usecase

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/datastore"
)

type TimelineUseCase interface {
	GetTimelinesByDisasterID(ctx context.Context, disasterID string) ([]*model.Timeline, error)
}

type timelineUseCase struct {
	timelineRepository datastore.TimelineRepository
}

func NewTimelineUseCase(
	timelineRepository datastore.TimelineRepository,
) TimelineUseCase {
	return &timelineUseCase{
		timelineRepository: timelineRepository,
	}
}

func (u *timelineUseCase) GetTimelinesByDisasterID(ctx context.Context, disasterID string) ([]*model.Timeline, error) {
	return u.timelineRepository.FindByDisasterID(ctx, disasterID)
}
