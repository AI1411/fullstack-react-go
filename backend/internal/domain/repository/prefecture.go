//go:generate mockgen -source=prefecture_repository.go -destination=../../../tests/mock/datastore/prefecture_repository.mock.go
package domain

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
)

type PrefectureRepository interface {
	Find(ctx context.Context) ([]*model.Prefecture, error)
	FindByID(ctx context.Context, code string) (*model.Prefecture, error)
}
