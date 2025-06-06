//go:generate mockgen -source=email_varification__token.go -destination=../../../tests/mock/domain/email_varification__token.mock.go
package domain

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
)

type EmailVarificationTokenRepository interface {
	Save(ctx context.Context, token *model.EmailVerificationToken) error
	FindByTokenAndUserID(ctx context.Context, token string) (*model.EmailVerificationToken, error)
	MarkAsUsed(ctx context.Context, tokenID string) error
}
