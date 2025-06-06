//go:generate mockgen -source=notification.go -destination=../../../tests/mock/domain/notification.mock.go
package domain

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
)

type NotificationRepository interface {
	Find(ctx context.Context) ([]*model.Notification, error)
	FindByID(ctx context.Context, id int32) (*model.Notification, error)
	FindByUserID(ctx context.Context, userID int32) ([]*model.Notification, error)
	Create(ctx context.Context, notification *model.Notification) error
	Update(ctx context.Context, notification *model.Notification) error
	Delete(ctx context.Context, id int32) error
	MarkAsRead(ctx context.Context, id int32) error
}
