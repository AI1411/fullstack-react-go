//go:generate mockgen -source=token.go -destination=../../tests/mock/utils/token.mock.go
package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type TokenGenerator interface {
	GenerateEmailVerificationToken() (string, error)
}

type tokenGenerator struct{}

func NewTokenGenerator() TokenGenerator {
	return &tokenGenerator{}
}

func (t *tokenGenerator) GenerateEmailVerificationToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}
