package domain

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
)

type JWT interface {
	GenerateToken(ctx context.Context, user *model.User) (string, error)
	ValidateToken(ctx context.Context, tokenString string) (*model.Claims, error)
	RefreshToken(ctx context.Context, tokenString string) (string, error)
}
