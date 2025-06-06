//go:generate mockgen -source=user_repository.go -destination=../../../tests/mock/datastore/user_repository.mock.go
package datastore

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/domain/query"
	domain "github.com/AI1411/fullstack-react-go/internal/domain/repository"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
)

type userRepository struct {
	client db.Client
	query  *query.Query
}

func NewUserRepository(
	ctx context.Context,
	client db.Client,
) domain.UserRepository {
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

func (r *userRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
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
	return r.query.User.Create(user)
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	return r.client.Conn(ctx).Save(user).Error
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := r.query.User.
		Where(r.query.User.Email.Eq(email)).
		Preload(r.query.User.Organizations).
		First()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Delete(ctx context.Context, id int32) error {
	return r.client.Conn(ctx).Delete(&model.User{}, id).Error
}
