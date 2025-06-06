//go:generate mockgen -source=support_application.go -destination=../../../tests/mock/domain/support_application.mock.go
package domain

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
)

type SupportApplicationRepository interface {
	Find(ctx context.Context) ([]*model.SupportApplication, error)
	FindByID(ctx context.Context, id string) (*model.SupportApplication, error)
	Create(ctx context.Context, supportApplication *model.SupportApplication) error
}
