//go:generate mockgen -source=support_application_usecase.go -destination=../../tests/mock/usecase/support_application_usecase.mock.go
package usecase

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/datastore"
)

type SupportApplicationUseCase interface {
	ListSupportApplications(ctx context.Context) ([]*model.SupportApplication, error)
	GetSupportApplicationByID(ctx context.Context, id string) (*model.SupportApplication, error)
	CreateSupportApplication(ctx context.Context, supportApplication *model.SupportApplication) error
}

type supportApplicationUseCase struct {
	supportApplicationRepository datastore.SupportApplicationRepository
}

func NewSupportApplicationUseCase(
	supportApplicationRepository datastore.SupportApplicationRepository,
) SupportApplicationUseCase {
	return &supportApplicationUseCase{
		supportApplicationRepository: supportApplicationRepository,
	}
}

func (u *supportApplicationUseCase) ListSupportApplications(ctx context.Context) ([]*model.SupportApplication, error) {
	supportApplications, err := u.supportApplicationRepository.Find(ctx)
	if err != nil {
		return nil, err
	}

	return supportApplications, nil
}

func (u *supportApplicationUseCase) GetSupportApplicationByID(ctx context.Context, id string) (*model.SupportApplication, error) {
	supportApplication, err := u.supportApplicationRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return supportApplication, nil
}

func (u *supportApplicationUseCase) CreateSupportApplication(ctx context.Context, supportApplication *model.SupportApplication) error {
	return u.supportApplicationRepository.Create(ctx, supportApplication)
}
