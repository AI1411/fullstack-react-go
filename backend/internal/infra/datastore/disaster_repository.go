package datastore

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/domain/query"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
)

type DisasterRepository interface {
	Find(ctx context.Context) ([]*model.Disaster, error)
}

type disasterRepository struct {
	client db.Client
	query  *query.Query
}

func NewDisasterRepository(
	ctx context.Context,
	client db.Client,
) DisasterRepository {
	return &disasterRepository{
		client: client,
		query:  query.Use(client.Conn(ctx)),
	}
}

func (r *disasterRepository) Find(ctx context.Context) ([]*model.Disaster, error) {
	disasters, err := r.query.Disaster.Find()
	if err != nil {
		return nil, err
	}

	return disasters, nil
}
