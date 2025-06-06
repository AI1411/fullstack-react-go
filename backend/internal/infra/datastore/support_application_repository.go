//go:generate mockgen -source=support_application_repository.go -destination=../../../tests/mock/datastore/support_application_repository.mock.go
package datastore

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/domain/query"
	domain "github.com/AI1411/fullstack-react-go/internal/domain/repository"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
)

type supportApplicationRepository struct {
	client db.Client
	query  *query.Query
}

func NewSupportApplicationRepository(
	ctx context.Context,
	client db.Client,
) domain.SupportApplicationRepository {
	return &supportApplicationRepository{
		client: client,
		query:  query.Use(client.Conn(ctx)),
	}
}

func (r *supportApplicationRepository) Find(ctx context.Context) ([]*model.SupportApplication, error) {
	supportApplications, err := r.query.WithContext(ctx).SupportApplication.
		Order(r.query.SupportApplication.ApplicationDate.Desc()).
		Find()
	if err != nil {
		return nil, err
	}

	return supportApplications, nil
}

func (r *supportApplicationRepository) FindByID(ctx context.Context, id string) (*model.SupportApplication, error) {
	supportApplication, err := r.query.WithContext(ctx).
		SupportApplication.
		Where(r.query.SupportApplication.ApplicationID.Eq(id)).
		First()
	if err != nil {
		return nil, err
	}

	return supportApplication, nil
}

func (r *supportApplicationRepository) Create(ctx context.Context, supportApplication *model.SupportApplication) error {
	return r.query.WithContext(ctx).SupportApplication.Create(supportApplication)
}
