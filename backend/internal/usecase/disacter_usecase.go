package usecase

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/datastore"
)

type DisasterUseCase interface {
	ListDisasters(ctx context.Context) ([]*model.Disaster, error)
	GetDisasterByID(ctx context.Context, id string) (*model.Disaster, error)
	CreateDisaster(ctx context.Context, disaster *model.Disaster) error
	UpdateDisaster(ctx context.Context, disaster *model.Disaster) error
	DeleteDisaster(ctx context.Context, id string) error
}

type disasterUseCase struct {
	disasterRepository datastore.DisasterRepository
}

func NewDisasterUseCase(
	disasterRepository datastore.DisasterRepository,
) DisasterUseCase {
	return &disasterUseCase{
		disasterRepository: disasterRepository,
	}
}

func (u *disasterUseCase) ListDisasters(ctx context.Context) ([]*model.Disaster, error) {
	disasters, err := u.disasterRepository.Find(ctx)
	if err != nil {
		return nil, err
	}

	return disasters, nil
}

func (u *disasterUseCase) GetDisasterByID(ctx context.Context, id string) (*model.Disaster, error) {
	disaster, err := u.disasterRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return disaster, nil
}

func (u *disasterUseCase) CreateDisaster(ctx context.Context, disaster *model.Disaster) error {
	return u.disasterRepository.Create(ctx, disaster)
}

func (u *disasterUseCase) UpdateDisaster(ctx context.Context, disaster *model.Disaster) error {
	return u.disasterRepository.Update(ctx, disaster)
}

func (u *disasterUseCase) DeleteDisaster(ctx context.Context, id string) error {
	return u.disasterRepository.Delete(ctx, id)
}
