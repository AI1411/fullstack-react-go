package datastore

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/domain/query"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
)

type UserRepository interface {
	Find(ctx context.Context) ([]*model.User, error)
	FindByID(ctx context.Context, id int32) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int32) error
}

type userRepository struct {
	client db.Client
	query  *query.Query
}

func NewUserRepository(
	ctx context.Context,
	client db.Client,
) UserRepository {
	return &userRepository{
		client: client,
		query:  query.Use(client.Conn(ctx)),
	}
}

func (r *userRepository) Find(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	if err := r.client.Conn(ctx).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) FindByID(ctx context.Context, id int32) (*model.User, error) {
	user, err := r.query.User.
		Where(r.query.User.ID.Eq(id)).
		Preload(r.query.User.Organizations).
		First()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	return r.client.Conn(ctx).Create(user).Error
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	return r.client.Conn(ctx).Save(user).Error
}

func (r *userRepository) Delete(ctx context.Context, id int32) error {
	return r.client.Conn(ctx).Delete(&model.User{}, id).Error
}
