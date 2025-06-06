//go:generate mockgen -source=email_varification_token.go -destination=../../../tests/mock/domain/email_varification_token.mock.go
package domain

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
)

type EmailVarificationTokenRepository interface {
	Save(ctx context.Context, token *model.EmailVerificationToken) error
	FindByToken(ctx context.Context, token string) (*model.EmailVerificationToken, error)
	MarkAsUsed(ctx context.Context, tokenID string) error
}
