//go:generate mockgen -source=email_history_repository.go -destination=../../tests/mock/datastore/email_history_repository.mock.go
package datastore

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/domain/query"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
)

type EmailRepository interface {
	SaveEmailHistory(ctx context.Context, email *model.EmailHistory) error
	ListEmailHistoriesByUserID(ctx context.Context, userID string) ([]*model.EmailHistory, error)
}

type emailRepository struct {
	dbClient db.Client
	query    *query.Query
}

func NewEmailRepository(ctx context.Context, dbClient db.Client) EmailRepository {
	return &emailRepository{
		dbClient: dbClient,
		query:    query.Use(dbClient.Conn(ctx)),
	}
}

func (r *emailRepository) SaveEmailHistory(ctx context.Context, email *model.EmailHistory) error {
	err := r.query.WithContext(ctx).EmailHistory.Create(email)
	if err != nil {
		return err
	}

	return nil
}

func (r *emailRepository) ListEmailHistoriesByUserID(ctx context.Context, userID string) ([]*model.EmailHistory, error) {
	q := r.query.WithContext(ctx).EmailHistory.
		Where(r.query.EmailHistory.UserID.Eq(userID)).
		Order(r.query.EmailHistory.SentAt.Desc())

	emailHistories, err := q.Find()
	if err != nil {
		return nil, err
	}

	return emailHistories, nil
}
