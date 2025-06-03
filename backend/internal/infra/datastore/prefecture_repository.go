//go:generate mockgen -source=prefecture_repository.go -destination=../../../tests/mock/datastore/prefecture_repository.mock.go
package datastore

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/domain/query"
	myerrors "github.com/AI1411/fullstack-react-go/internal/errors"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
)

type PrefectureRepository interface {
	Find(ctx context.Context) ([]*model.Prefecture, error)
	FindByID(ctx context.Context, code string) (*model.Prefecture, error)
}

type prefectureRepository struct {
	client db.Client
	query  *query.Query
}

func NewPrefectureRepository(
	ctx context.Context,
	client db.Client,
) PrefectureRepository {
	return &prefectureRepository{
		client: client,
		query:  query.Use(client.Conn(ctx)),
	}
}

func (r *prefectureRepository) Find(ctx context.Context) ([]*model.Prefecture, error) {
	prefectures, err := r.query.WithContext(ctx).Prefecture.Find()
	if err != nil {
		return nil, err
	}

	return prefectures, nil
}

func (r *prefectureRepository) FindByID(ctx context.Context, code string) (*model.Prefecture, error) {
	prefecture, err := r.query.WithContext(ctx).
		Prefecture.
		Where(r.query.Prefecture.Code.Eq(code)).
		Preload(r.query.Prefecture.Municipalities.On(r.query.Municipality.IsActive.Is(true))).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &myerrors.APIError{
				Code:    myerrors.PrefectureNotFoundError,
				Message: myerrors.PrefectureNotFoundErrorMessage,
			}
		}
		return nil, err
	}

	return prefecture, nil
}
