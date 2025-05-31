package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/logger"
	"github.com/AI1411/fullstack-react-go/internal/usecase"
)

type User interface {
	ListUsers(c *gin.Context)
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userHandler struct {
	logger      *logger.Logger
	userUseCase usecase.UserUseCase
}

func NewUserHandler(l *logger.Logger, userUseCase usecase.UserUseCase) User {
	return &userHandler{
		logger:      l,
		userUseCase: userUseCase,
	}
}

type UserResponse struct {
	ID        int32      `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"omitempty,min=6"`
}

func (h *userHandler) ListUsers(c *gin.Context) {
	ctx := c.Request.Context()

	users, err := h.userUseCase.ListUsers(ctx)
	if err != nil {
		h.logger.Error("Failed to get users", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})

		return
	}

	var response []UserResponse
	for _, user := range users {
		response = append(response, UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) GetUser(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid user ID", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})

		return
	}

	user, err := h.userUseCase.GetUserByID(ctx, int32(id))
	if err != nil {
		h.logger.Error("Failed to get user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})

		return
	}

	response := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	now := time.Now()
	user := &model.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	if err := h.userUseCase.CreateUser(ctx, user); err != nil {
		h.logger.Error("Failed to create user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})

		return
	}

	response := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid user ID", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})

		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	user, err := h.userUseCase.GetUserByID(ctx, int32(id))
	if err != nil {
		h.logger.Error("Failed to get user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})

		return
	}

	now := time.Now()
	user.Name = req.Name
	user.Email = req.Email

	if req.Password != "" {
		user.Password = req.Password
	}

	user.UpdatedAt = &now

	if err := h.userUseCase.UpdateUser(ctx, user); err != nil {
		h.logger.Error("Failed to update user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})

		return
	}

	response := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid user ID", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})

		return
	}

	if err := h.userUseCase.DeleteUser(ctx, int32(id)); err != nil {
		h.logger.Error("Failed to delete user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})

		return
	}

	c.JSON(http.StatusNoContent, nil)
}
