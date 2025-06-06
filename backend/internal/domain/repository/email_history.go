package domain

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
)

type EmailHistoryRepository interface {
	SaveEmailHistory(ctx context.Context, email *model.EmailHistory) error
	ListEmailHistoriesByUserID(ctx context.Context, userID string) ([]*model.EmailHistory, error)
}
