package usecase

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	domain "github.com/AI1411/fullstack-react-go/internal/domain/repository"
)

type EmailVarificationTokenUsecase interface {
	CreateEmailVarificationToken(ctx context.Context, token *model.EmailVerificationToken) error
	FindEmailVarificationTokenByTokenAndUserID(ctx context.Context, token string) (*model.EmailVerificationToken, error)
	MarkEmailVarificationTokenAsUsed(ctx context.Context, tokenID string) error
}

type emailVarificationTokenUsecase struct {
	emailVarificationTokenRepo domain.EmailVarificationTokenRepository
}

func NewEmailVarificationTokenRepository(
	emailVarificationTokenRepo domain.EmailVarificationTokenRepository,
) EmailVarificationTokenUsecase {
	return &emailVarificationTokenUsecase{
		emailVarificationTokenRepo: emailVarificationTokenRepo,
	}
}

func (e emailVarificationTokenUsecase) CreateEmailVarificationToken(ctx context.Context, token *model.EmailVerificationToken) error {
	if err := e.emailVarificationTokenRepo.Save(ctx, token); err != nil {
		return err
	}

	return nil
}

func (e emailVarificationTokenUsecase) FindEmailVarificationTokenByTokenAndUserID(ctx context.Context, token string) (*model.EmailVerificationToken, error) {
	verificationToken, err := e.emailVarificationTokenRepo.FindByTokenAndUserID(ctx, token)
	if err != nil {
		return nil, err
	}

	return verificationToken, nil
}

func (e emailVarificationTokenUsecase) MarkEmailVarificationTokenAsUsed(ctx context.Context, tokenID string) error {
	if err := e.emailVarificationTokenRepo.MarkAsUsed(ctx, tokenID); err != nil {
		return err
	}

	return nil
}
