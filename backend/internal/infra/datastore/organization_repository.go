package datastore

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
)

type OrganizationRepository interface {
	Find(ctx context.Context) ([]*model.Organization, error)
	FindByID(ctx context.Context, id int32) (*model.Organization, error)
	Create(ctx context.Context, organization *model.Organization) error
	Update(ctx context.Context, organization *model.Organization) error
	Delete(ctx context.Context, id int32) error
}

type organizationRepository struct {
	client db.Client
}

func NewOrganizationRepository(
	ctx context.Context,
	client db.Client,
) OrganizationRepository {
	return &organizationRepository{
		client: client,
	}
}

func (r *organizationRepository) Find(ctx context.Context) ([]*model.Organization, error) {
	var organizations []*model.Organization
	if err := r.client.Conn(ctx).Find(&organizations).Error; err != nil {
		return nil, err
	}

	return organizations, nil
}

func (r *organizationRepository) FindByID(ctx context.Context, id int32) (*model.Organization, error) {
	var organization model.Organization
	if err := r.client.Conn(ctx).Where("id = ?", id).First(&organization).Error; err != nil {
		return nil, err
	}

	return &organization, nil
}

func (r *organizationRepository) Create(ctx context.Context, organization *model.Organization) error {
	return r.client.Conn(ctx).Create(organization).Error
}

func (r *organizationRepository) Update(ctx context.Context, organization *model.Organization) error {
	return r.client.Conn(ctx).Save(organization).Error
}

func (r *organizationRepository) Delete(ctx context.Context, id int32) error {
	return r.client.Conn(ctx).Delete(&model.Organization{}, id).Error
}
