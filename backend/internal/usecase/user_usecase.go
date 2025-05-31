//go:generate mockgen -source=user_usecase.go -destination=../../tests/mock/usecase/user_usecase.mock.go
package usecase

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/datastore"
)

type UserUseCase interface {
	ListUsers(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, id int32) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id int32) error
}

type userUseCase struct {
	userRepository datastore.UserRepository
}

func NewUserUseCase(
	userRepository datastore.UserRepository,
) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (u *userUseCase) ListUsers(ctx context.Context) ([]*model.User, error) {
	users, err := u.userRepository.Find(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userUseCase) GetUserByID(ctx context.Context, id int32) (*model.User, error) {
	user, err := u.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUseCase) CreateUser(ctx context.Context, user *model.User) error {
	return u.userRepository.Create(ctx, user)
}

func (u *userUseCase) UpdateUser(ctx context.Context, user *model.User) error {
	return u.userRepository.Update(ctx, user)
}

func (u *userUseCase) DeleteUser(ctx context.Context, id int32) error {
	return u.userRepository.Delete(ctx, id)
}
