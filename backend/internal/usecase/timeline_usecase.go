//go:generate mockgen -source=timeline_usecase.go -destination=../../tests/mock/usecase/timeline_usecase.mock.go
package usecase

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	domain "github.com/AI1411/fullstack-react-go/internal/domain/repository"
)

type TimelineUseCase interface {
	GetTimelinesByDisasterID(ctx context.Context, disasterID string) ([]*model.Timeline, error)
}

type timelineUseCase struct {
	timelineRepository domain.TimelineRepository
}

func NewTimelineUseCase(
	timelineRepository domain.TimelineRepository,
) TimelineUseCase {
	return &timelineUseCase{
		timelineRepository: timelineRepository,
	}
}

func (u *timelineUseCase) GetTimelinesByDisasterID(ctx context.Context, disasterID string) ([]*model.Timeline, error) {
	return u.timelineRepository.FindByDisasterID(ctx, disasterID)
}
