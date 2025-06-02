package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/env"
	"github.com/AI1411/fullstack-react-go/internal/infra/logger"
	"github.com/AI1411/fullstack-react-go/internal/usecase"
)

// Auth interface defines the methods for authentication
type Auth interface {
	Login(c *gin.Context)
	Callback(c *gin.Context)
	Logout(c *gin.Context)
}

type authHandler struct {
	logger       *logger.Logger
	userUseCase  usecase.UserUseCase
	env          *env.Values
	provider     *oidc.Provider
	oauth2Config oauth2.Config
	verifier     *oidc.IDTokenVerifier
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(
	ctx context.Context,
	l *logger.Logger,
	userUseCase usecase.UserUseCase,
	env *env.Values,
) (Auth, error) {
	// If OIDC issuer is not set, return a mock auth handler
	if env.OIDCIssuer == "" {
		l.Warn("OIDC issuer not set, using mock auth handler")
		return NewMockAuthHandler(l, userUseCase, env), nil
	}

	// Initialize OIDC provider
	provider, err := oidc.NewProvider(ctx, env.OIDCIssuer)
	if err != nil {
		l.Warn("Failed to initialize OIDC provider, using mock auth handler", "error", err)
		return NewMockAuthHandler(l, userUseCase, env), nil
	}

	// Configure OAuth2
	oauth2Config := oauth2.Config{
		ClientID:     env.OIDCClientID,
		ClientSecret: env.OIDCClientSecret,
		RedirectURL:  env.OIDCRedirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	// Configure ID token verifier
	verifier := provider.Verifier(&oidc.Config{
		ClientID: env.OIDCClientID,
	})

	return &authHandler{
		logger:       l,
		userUseCase:  userUseCase,
		env:          env,
		provider:     provider,
		oauth2Config: oauth2Config,
		verifier:     verifier,
	}, nil
}

// mockAuthHandler is a mock implementation of the Auth interface
type mockAuthHandler struct {
	logger      *logger.Logger
	userUseCase usecase.UserUseCase
	env         *env.Values
}

// NewMockAuthHandler creates a new mock auth handler
func NewMockAuthHandler(l *logger.Logger, userUseCase usecase.UserUseCase, env *env.Values) Auth {
	return &mockAuthHandler{
		logger:      l,
		userUseCase: userUseCase,
		env:         env,
	}
}

// Login handles the login request for the mock auth handler
func (h *mockAuthHandler) Login(c *gin.Context) {
	h.logger.Warn("Mock auth handler: Login called")
	c.JSON(http.StatusOK, gin.H{"message": "Mock login endpoint. OIDC not configured."})
}

// Callback handles the callback request for the mock auth handler
func (h *mockAuthHandler) Callback(c *gin.Context) {
	h.logger.Warn("Mock auth handler: Callback called")
	c.JSON(http.StatusOK, gin.H{"message": "Mock callback endpoint. OIDC not configured."})
}

// Logout handles the logout request for the mock auth handler
func (h *mockAuthHandler) Logout(c *gin.Context) {
	h.logger.Warn("Mock auth handler: Logout called")
	c.JSON(http.StatusOK, gin.H{"message": "Mock logout endpoint. OIDC not configured."})
}

// generateRandomState generates a random state string for CSRF protection
func generateRandomState() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// Login initiates the OIDC authentication flow
// @title ログイン
// @id Login
// @tags auth
// @accept json
// @produce json
// @version 1.0
// @description OIDCプロバイダーへのログインを開始します
// @Summary OIDCログイン
// @Success 302
// @Router /auth/login [get]
func (h *authHandler) Login(c *gin.Context) {
	// Generate random state for CSRF protection
	state, err := generateRandomState()
	if err != nil {
		h.logger.Error("Failed to generate random state", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initiate login"})
		return
	}

	// Store state in cookie
	c.SetCookie("auth_state", state, 3600, "/", "", false, true)

	// Redirect to OIDC provider
	authURL := h.oauth2Config.AuthCodeURL(state)
	c.Redirect(http.StatusFound, authURL)
}

// Callback handles the OIDC callback
// @title OIDCコールバック
// @id Callback
// @tags auth
// @accept json
// @produce json
// @version 1.0
// @description OIDCプロバイダーからのコールバックを処理します
// @Summary OIDCコールバック
// @Param code query string true "認証コード"
// @Param state query string true "状態"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/callback [get]
func (h *authHandler) Callback(c *gin.Context) {
	ctx := c.Request.Context()

	// Verify state for CSRF protection
	state := c.Query("state")
	storedState, err := c.Cookie("auth_state")
	if err != nil || state != storedState {
		h.logger.Error("Invalid state", "error", err, "state", state, "storedState", storedState)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state"})
		return
	}

	// Clear state cookie
	c.SetCookie("auth_state", "", -1, "/", "", false, true)

	// Exchange code for token
	code := c.Query("code")
	oauth2Token, err := h.oauth2Config.Exchange(ctx, code)
	if err != nil {
		h.logger.Error("Failed to exchange code for token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange code for token"})
		return
	}

	// Extract ID token
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		h.logger.Error("No ID token in token response")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No ID token in token response"})
		return
	}

	// Verify ID token
	idToken, err := h.verifier.Verify(ctx, rawIDToken)
	if err != nil {
		h.logger.Error("Failed to verify ID token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify ID token"})
		return
	}

	// Extract claims
	var claims struct {
		Email   string `json:"email"`
		Name    string `json:"name"`
		Sub     string `json:"sub"`
		Picture string `json:"picture"`
	}
	if err := idToken.Claims(&claims); err != nil {
		h.logger.Error("Failed to extract claims", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract claims"})
		return
	}

	// Find or create user
	user, err := h.userUseCase.GetUserByEmail(ctx, claims.Email)
	if err != nil {
		// Create new user if not found
		user = &model.User{
			Name:     claims.Name,
			Email:    claims.Email,
			Password: "", // No password for OIDC users
		}
		if err := h.userUseCase.CreateUser(ctx, user); err != nil {
			h.logger.Error("Failed to create user", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Duration(h.env.JWTExpiration) * time.Second).Unix(),
	})

	// Sign token
	tokenString, err := token.SignedString([]byte(h.env.JWTSecret))
	if err != nil {
		h.logger.Error("Failed to sign token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign token"})
		return
	}

	// Set token in cookie
	c.SetCookie("auth_token", tokenString, h.env.JWTExpiration, "/", "", false, true)

	// Return token
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

// Logout logs out the user
// @title ログアウト
// @id Logout
// @tags auth
// @accept json
// @produce json
// @version 1.0
// @description ユーザーをログアウトします
// @Summary ログアウト
// @Success 200 {object} map[string]string
// @Router /auth/logout [post]
func (h *authHandler) Logout(c *gin.Context) {
	// Clear token cookie
	c.SetCookie("auth_token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
