//go:generate mockgen -source=notification_usecase.go -destination=../../tests/mock/usecase/notification_usecase.mock.go
package usecase

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/datastore"
)

type NotificationUseCase interface {
	ListNotifications(ctx context.Context) ([]*model.Notification, error)
	GetNotificationByID(ctx context.Context, id int32) (*model.Notification, error)
	GetNotificationsByUserID(ctx context.Context, userID int32) ([]*model.Notification, error)
	CreateNotification(ctx context.Context, notification *model.Notification) error
	UpdateNotification(ctx context.Context, notification *model.Notification) error
	DeleteNotification(ctx context.Context, id int32) error
	MarkAsRead(ctx context.Context, id int32) error
}

type notificationUseCase struct {
	notificationRepository datastore.NotificationRepository
}

func NewNotificationUseCase(
	notificationRepository datastore.NotificationRepository,
) NotificationUseCase {
	return &notificationUseCase{
		notificationRepository: notificationRepository,
	}
}

func (u *notificationUseCase) ListNotifications(ctx context.Context) ([]*model.Notification, error) {
	notifications, err := u.notificationRepository.Find(ctx)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func (u *notificationUseCase) GetNotificationByID(ctx context.Context, id int32) (*model.Notification, error) {
	notification, err := u.notificationRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return notification, nil
}

func (u *notificationUseCase) GetNotificationsByUserID(ctx context.Context, userID int32) ([]*model.Notification, error) {
	notifications, err := u.notificationRepository.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func (u *notificationUseCase) CreateNotification(ctx context.Context, notification *model.Notification) error {
	return u.notificationRepository.Create(ctx, notification)
}

func (u *notificationUseCase) UpdateNotification(ctx context.Context, notification *model.Notification) error {
	return u.notificationRepository.Update(ctx, notification)
}

func (u *notificationUseCase) DeleteNotification(ctx context.Context, id int32) error {
	return u.notificationRepository.Delete(ctx, id)
}

func (u *notificationUseCase) MarkAsRead(ctx context.Context, id int32) error {
	return u.notificationRepository.MarkAsRead(ctx, id)
}
