package datastore

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/domain/query"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
)

type PrefectureRepository interface {
	Find(ctx context.Context) ([]*model.Prefecture, error)
	FindByID(ctx context.Context, id int32) (*model.Prefecture, error)
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

func (r *prefectureRepository) FindByID(ctx context.Context, id int32) (*model.Prefecture, error) {
	prefecture, err := r.query.WithContext(ctx).Prefecture.Where(r.query.Prefecture.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}

	return prefecture, nil
}
