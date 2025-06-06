package datastore

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/domain/query"
	domain "github.com/AI1411/fullstack-react-go/internal/domain/repository"
	myerrors "github.com/AI1411/fullstack-react-go/internal/errors"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
)

type emailVarificationTokenRepository struct {
	client db.Client
	query  *query.Query
}

func NewEmailVarificationTokenRepository(
	ctx context.Context,
	client db.Client,
) domain.EmailVarificationTokenRepository {
	return &emailVarificationTokenRepository{
		client: client,
		query:  query.Use(client.Conn(ctx)),
	}
}

func (e *emailVarificationTokenRepository) Save(ctx context.Context, token *model.EmailVerificationToken) error {
	if err := e.query.WithContext(ctx).EmailVerificationToken.Create(token); err != nil {
		return err
	}
	return nil
}

func (e *emailVarificationTokenRepository) FindByTokenAndUserID(ctx context.Context, token string) (*model.EmailVerificationToken, error) {
	verificationToken, err := e.query.WithContext(ctx).
		EmailVerificationToken.
		Where(
			e.query.EmailVerificationToken.Token.Eq(token),
		).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerrors.APIError{
				Code:    myerrors.EmailVarificationTokenNotFound,
				Message: myerrors.EmailVarificationTokenNotFoundErrorMessage,
			}
		}

		return nil, err
	}

	return verificationToken, nil
}

func (e *emailVarificationTokenRepository) MarkAsUsed(ctx context.Context, tokenID string) error {
	result, err := e.query.WithContext(ctx).
		EmailVerificationToken.
		Where(
			e.query.EmailVerificationToken.ID.Eq(tokenID),
		).
		Update(e.query.EmailVerificationToken.IsUsed, true)
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return myerrors.APIError{
			Code:    myerrors.EmailVarificationTokenNotFound,
			Message: myerrors.EmailVarificationTokenNotFoundErrorMessage,
		}
	}

	return nil
}
