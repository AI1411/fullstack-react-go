//go:generate mockgen -source=organization_repository.go -destination=../../../tests/mock/datastore/organization_repository.mock.go
package datastore

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/domain/query"
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
	query  *query.Query
}

func NewOrganizationRepository(
	ctx context.Context,
	client db.Client,
) OrganizationRepository {
	return &organizationRepository{
		client: client,
		query:  query.Use(client.Conn(ctx)),
	}
}

func (r *organizationRepository) Find(ctx context.Context) ([]*model.Organization, error) {
	organizations, err := r.query.WithContext(ctx).Organization.Find()
	if err != nil {
		return nil, err
	}

	return organizations, nil
}

func (r *organizationRepository) FindByID(ctx context.Context, id int32) (*model.Organization, error) {
	organization, err := r.query.WithContext(ctx).
		Organization.
		Where(r.query.Organization.ID.Eq(id)).
		Preload(r.query.Organization.Users).
		First()
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (r *organizationRepository) Create(ctx context.Context, organization *model.Organization) error {
	return r.query.WithContext(ctx).Organization.Create(organization)
}

func (r *organizationRepository) Update(ctx context.Context, organization *model.Organization) error {
	_, err := r.query.WithContext(ctx).Organization.Where(r.query.Organization.ID.Eq(organization.ID)).Updates(organization)
	return err
}

func (r *organizationRepository) Delete(ctx context.Context, id int32) error {
	_, err := r.query.WithContext(ctx).Organization.Where(r.query.Organization.ID.Eq(id)).Delete()
	return err
}
