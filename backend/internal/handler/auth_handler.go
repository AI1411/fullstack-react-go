package handler

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	Register(c *gin.Context)
	VerifyEmail(c *gin.Context)
}

type authHandler struct {
	logger                        *logger.Logger
	userUseCase                   usecase.UserUseCase
	authUsecase                   usecase.AuthUsecase
	emailVarificationTokenUsecase usecase.EmailVarificationTokenUsecase
	env                           *env.Values
	provider                      *oidc.Provider
	oauth2Config                  oauth2.Config
	verifier                      *oidc.IDTokenVerifier
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(
	l *logger.Logger,
	env *env.Values,
	userUseCase usecase.UserUseCase,
	authUsecase usecase.AuthUsecase,
	emailVarificationTokenUsecase usecase.EmailVarificationTokenUsecase,
) (Auth, error) {
	return &authHandler{
		logger:                        l,
		env:                           env,
		userUseCase:                   userUseCase,
		authUsecase:                   authUsecase,
		emailVarificationTokenUsecase: emailVarificationTokenUsecase,
	}, nil
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

	// Generate JWT token using auth usecase
	tokenString, err := h.authUsecase.GenerateToken(user)
	if err != nil {
		h.logger.Error("Failed to generate token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
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

// RegisterRequest defines the request body for user registration
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

// Register registers a new user
// @title ユーザー登録
// @id Register
// @tags auth
// @accept json
// @produce json
// @version 1.0
// @description 新規ユーザーを登録します
// @Summary ユーザー登録
// @Param request body RegisterRequest true "ユーザー登録リクエスト"
// @Success 201 {object} RegisterResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/register [post]
func (h *authHandler) Register(c *gin.Context) {
	ctx := c.Request.Context()

	// Parse request body
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	existingUser, err := h.userUseCase.GetUserByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		h.logger.Error("User already exists", "email", req.Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with this email already exists"})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		h.logger.Error("Failed to hash password", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// Create user
	now := time.Now()
	user := &model.User{
		Name:          req.Name,
		Email:         req.Email,
		Password:      string(hashedPassword),
		IsActive:      false,
		EmailVerified: false,
		RoleID:        1, // Default role
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}

	if err := h.userUseCase.CreateUser(ctx, user); err != nil {
		h.logger.Error("Failed to create user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// Generate JWT token using auth usecase
	tokenString, err := h.authUsecase.GenerateToken(user)
	if err != nil {
		h.logger.Error("Failed to generate token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate authentication token"})
		return
	}

	// Set token in cookie
	c.SetCookie("auth_token", tokenString, h.env.JWTExpiration, "/", "", false, true)

	// Return token and user info
	c.JSON(http.StatusCreated, gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

type VerifyEmailRequest struct {
	Token string `json:"token" binding:"required"`
}

func (h *authHandler) VerifyEmail(c *gin.Context) {
	var req VerifyEmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なリクエストです"})
		return
	}

	// Use auth usecase to verify email
	err := h.authUsecase.VerifyEmail(req.Token)
	if err != nil {
		h.logger.Error("Failed to verify email", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "メールアドレスの確認に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "メールアドレスが確認されました"})
}
