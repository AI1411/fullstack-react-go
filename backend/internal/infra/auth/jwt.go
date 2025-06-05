// internal/infra/auth/jwt.go
package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	domain "github.com/AI1411/fullstack-react-go/internal/domain/repository"
)

type JWTConfig struct {
	SecretKey  string
	Expiration time.Duration
	Issuer     string
}

type jwtClient struct {
	config JWTConfig
}

func NewJWTClient(config JWTConfig) domain.JWT {
	return &jwtClient{config: config}
}

func (j *jwtClient) GenerateToken(ctx context.Context, user *model.User) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"iss":     j.config.Issuer,
		"iat":     now.Unix(),
		"exp":     now.Add(j.config.Expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.config.SecretKey))
}

func (j *jwtClient) ValidateToken(ctx context.Context, tokenString string) (*model.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.config.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	userID, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		return nil, errors.New("invalid user_id in token")
	}

	return &model.Claims{
		UserID:    userID.String(),
		Email:     claims["email"].(string),
		Role:      claims["role"].(model.Role),
		IssuedAt:  time.Unix(int64(claims["iat"].(float64)), 0),
		ExpiresAt: time.Unix(int64(claims["exp"].(float64)), 0),
	}, nil
}

func (j *jwtClient) RefreshToken(ctx context.Context, tokenString string) (string, error) {
	claims, err := j.ValidateToken(ctx, tokenString)
	if err != nil {
		return "", err
	}

	// リフレッシュ可能期間内かチェック
	if time.Until(claims.ExpiresAt) > j.config.Expiration/2 {
		return "", errors.New("token not eligible for refresh")
	}

	user := &model.User{
		ID:    claims.UserID,
		Email: claims.Email,
		Role:  claims.Role,
	}

	return j.GenerateToken(ctx, user)
}
