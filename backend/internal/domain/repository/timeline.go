//go:generate mockgen -source=timeline.go -destination=../../../tests/mock/domain/timeline.mock.go
package domain

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
)

type TimelineRepository interface {
	FindByDisasterID(ctx context.Context, disasterID string) ([]*model.Timeline, error)
}
