package usecase

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/datastore"
)

type DisasterUseCase interface {
	ListDisasters(ctx context.Context) ([]*model.Disaster, error)
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
