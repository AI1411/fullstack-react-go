//go:generate mockgen -source=organization.go -destination=../../../tests/mock/domain/organization.mock.go
package domain

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
)

type OrganizationRepository interface {
	Find(ctx context.Context) ([]*model.Organization, error)
	FindByID(ctx context.Context, id int64) (*model.Organization, error)
	Create(ctx context.Context, organization *model.Organization) error
	Update(ctx context.Context, organization *model.Organization) error
	Delete(ctx context.Context, id int64) error
}
