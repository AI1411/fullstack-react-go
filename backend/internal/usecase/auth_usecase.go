package usecase

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	domain "github.com/AI1411/fullstack-react-go/internal/domain/repository"
)

type AuthUsecase interface {
	Register(email string, password string) (string, error)
	VerifyEmail(token string) error
	ValidateToken(token string) (*model.Claims, error)
	GenerateToken(user *model.User) (string, error)
}

type authUsecase struct {
	jwtClient domain.JWT
}

func NewAuthUsecase(jwtClient domain.JWT) AuthUsecase {
	return &authUsecase{
		jwtClient: jwtClient,
	}
}

func (a authUsecase) Register(email string, password string) (string, error) {
	// This method would typically create a user and generate a token
	// But since user creation is handled in the handler, we'll just return an empty string
	// In a real implementation, this would create the user and generate a token
	return "", nil
}

func (a authUsecase) VerifyEmail(token string) error {
	// Validate the token
	_, err := a.ValidateToken(token)
	return err
}

func (a authUsecase) ValidateToken(token string) (*model.Claims, error) {
	// Use the JWT client to validate the token
	return a.jwtClient.ValidateToken(context.Background(), token)
}

func (a authUsecase) GenerateToken(user *model.User) (string, error) {
	// Use the JWT client to generate a token
	return a.jwtClient.GenerateToken(context.Background(), user)
}
