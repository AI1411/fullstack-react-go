package datastore

import (
	"context"
	"time"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
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

type notificationRepository struct {
	client db.Client
}

func NewNotificationRepository(
	ctx context.Context,
	client db.Client,
) NotificationRepository {
	return &notificationRepository{
		client: client,
	}
}

func (r *notificationRepository) Find(ctx context.Context) ([]*model.Notification, error) {
	var notifications []*model.Notification
	if err := r.client.Conn(ctx).Find(&notifications).Error; err != nil {
		return nil, err
	}

	return notifications, nil
}

func (r *notificationRepository) FindByID(ctx context.Context, id int32) (*model.Notification, error) {
	var notification model.Notification
	if err := r.client.Conn(ctx).Where("id = ?", id).First(&notification).Error; err != nil {
		return nil, err
	}

	return &notification, nil
}

func (r *notificationRepository) FindByUserID(ctx context.Context, userID int32) ([]*model.Notification, error) {
	var notifications []*model.Notification
	if err := r.client.Conn(ctx).Where("user_id = ?", userID).Find(&notifications).Error; err != nil {
		return nil, err
	}

	return notifications, nil
}

func (r *notificationRepository) Create(ctx context.Context, notification *model.Notification) error {
	return r.client.Conn(ctx).Create(notification).Error
}

func (r *notificationRepository) Update(ctx context.Context, notification *model.Notification) error {
	return r.client.Conn(ctx).Save(notification).Error
}

func (r *notificationRepository) Delete(ctx context.Context, id int32) error {
	return r.client.Conn(ctx).Delete(&model.Notification{}, id).Error
}

func (r *notificationRepository) MarkAsRead(ctx context.Context, id int32) error {
	now := time.Now()
	return r.client.Conn(ctx).Model(&model.Notification{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_read": true,
			"read_at": now,
		}).Error
}
