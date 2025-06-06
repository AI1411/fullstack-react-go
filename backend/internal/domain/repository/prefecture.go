//go:generate mockgen -source=prefecture.go -destination=../../../tests/mock/domain/prefecture.mock.go
package domain

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
)

type PrefectureRepository interface {
	Find(ctx context.Context) ([]*model.Prefecture, error)
	FindByID(ctx context.Context, code string) (*model.Prefecture, error)
}
