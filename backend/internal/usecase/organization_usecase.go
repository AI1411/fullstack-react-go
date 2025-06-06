//go:generate mockgen -source=organization_usecase.go -destination=../../tests/mock/usecase/organization_usecase.mock.go
package usecase

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	domain "github.com/AI1411/fullstack-react-go/internal/domain/repository"
)

type OrganizationUseCase interface {
	ListOrganizations(ctx context.Context) ([]*model.Organization, error)
	GetOrganizationByID(ctx context.Context, id int64) (*model.Organization, error)
	CreateOrganization(ctx context.Context, organization *model.Organization) error
	UpdateOrganization(ctx context.Context, organization *model.Organization) error
	DeleteOrganization(ctx context.Context, id int64) error
}

type organizationUseCase struct {
	organizationRepository domain.OrganizationRepository
}

func NewOrganizationUseCase(
	organizationRepository domain.OrganizationRepository,
) OrganizationUseCase {
	return &organizationUseCase{
		organizationRepository: organizationRepository,
	}
}

func (u *organizationUseCase) ListOrganizations(ctx context.Context) ([]*model.Organization, error) {
	organizations, err := u.organizationRepository.Find(ctx)
	if err != nil {
		return nil, err
	}

	return organizations, nil
}

func (u *organizationUseCase) GetOrganizationByID(ctx context.Context, id int64) (*model.Organization, error) {
	organization, err := u.organizationRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (u *organizationUseCase) CreateOrganization(ctx context.Context, organization *model.Organization) error {
	return u.organizationRepository.Create(ctx, organization)
}

func (u *organizationUseCase) UpdateOrganization(ctx context.Context, organization *model.Organization) error {
	return u.organizationRepository.Update(ctx, organization)
}

func (u *organizationUseCase) DeleteOrganization(ctx context.Context, id int64) error {
	return u.organizationRepository.Delete(ctx, id)
}
