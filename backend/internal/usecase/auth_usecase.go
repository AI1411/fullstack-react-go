//go:generate mockgen -source=auth_usecase.go -destination=../../tests/mock/usecase/auth_usecase.mock.go
package usecase

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	domain "github.com/AI1411/fullstack-react-go/internal/domain/repository"
	myerrors "github.com/AI1411/fullstack-react-go/internal/errors"
)

type AuthUsecase interface {
	ValidateEmailVarificationToken(ctx context.Context, token string) error
	GenerateToken(user *model.User) (string, error)
}

type authUsecase struct {
	jwtClient                  domain.JWT
	emailVarificationTokenRepo domain.EmailVarificationTokenRepository
}

func NewAuthUsecase(
	jwtClient domain.JWT,
	emailVarificationTokenRepo domain.EmailVarificationTokenRepository,
) AuthUsecase {
	return &authUsecase{
		jwtClient:                  jwtClient,
		emailVarificationTokenRepo: emailVarificationTokenRepo,
	}
}

func (a authUsecase) ValidateEmailVarificationToken(ctx context.Context, token string) error {
	emailToken, err := a.emailVarificationTokenRepo.FindByToken(ctx, token)
	if err != nil {
		return err
	}

	if emailToken == nil {
		return myerrors.APIError{
			Code:    myerrors.EmailVarificationTokenNotFound,
			Message: myerrors.EmailVarificationTokenNotFoundErrorMessage,
		}
	}

	if emailToken.IsUsed {
		return myerrors.APIError{
			Code:    myerrors.EmailVarificationTokenUsedError,
			Message: myerrors.EmailVarificationTokenUsedErrorMessage,
		}
	}

	// Mark the token as used
	if err := a.emailVarificationTokenRepo.MarkAsUsed(ctx, emailToken.ID); err != nil {
		return myerrors.APIError{
			Code:    myerrors.SystemError,
			Message: myerrors.SystemErrorMessage,
		}
	}

	return nil
}

func (a authUsecase) GenerateToken(user *model.User) (string, error) {
	// Use the JWT client to generate a token
	return a.jwtClient.GenerateToken(context.Background(), user)
}
