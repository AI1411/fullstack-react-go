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

type Organization interface {
	ListOrganizations(c *gin.Context)
	GetOrganization(c *gin.Context)
	CreateOrganization(c *gin.Context)
	UpdateOrganization(c *gin.Context)
	DeleteOrganization(c *gin.Context)
}

type organizationHandler struct {
	l                   *logger.Logger
	organizationUseCase usecase.OrganizationUseCase
}

func NewOrganizationHandler(
	l *logger.Logger,
	organizationUseCase usecase.OrganizationUseCase,
) Organization {
	return &organizationHandler{
		l:                   l,
		organizationUseCase: organizationUseCase,
	}
}

type OrganizationResponse struct {
	ID           int32          `json:"id"`
	Name         string         `json:"name"`
	Type         string         `json:"type"`
	PrefectureID *int32         `json:"prefecture_id,omitempty"`
	ParentID     *int32         `json:"parent_id,omitempty"`
	Description  *string        `json:"description,omitempty"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	Users        []UserResponse `json:"users,omitempty"`
}

type CreateOrganizationRequest struct {
	Name         string  `json:"name" binding:"required"`
	Type         string  `json:"type" binding:"required"`
	PrefectureID *int32  `json:"prefecture_id"`
	ParentID     *int32  `json:"parent_id"`
	Description  *string `json:"description"`
}

type UpdateOrganizationRequest struct {
	Name         string  `json:"name" binding:"required"`
	Type         string  `json:"type" binding:"required"`
	PrefectureID *int32  `json:"prefecture_id"`
	ParentID     *int32  `json:"parent_id"`
	Description  *string `json:"description"`
}

// ListOrganizations @title 組織一覧取得
// @id ListOrganizations
// @tags organization
// @accept json
// @produce json
// @Summary 組織一覧取得
// @Success 200 {array} OrganizationResponse
// @Router /organizations [get]
func (h *organizationHandler) ListOrganizations(c *gin.Context) {
	ctx := c.Request.Context()
	organizations, err := h.organizationUseCase.ListOrganizations(ctx)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to list organizations")
		c.JSON(500, gin.H{"error": "Internal Server Error"})

		return
	}

	var response []*OrganizationResponse

	for _, organization := range organizations {
		resp := &OrganizationResponse{
			ID:           organization.ID,
			Name:         organization.Name,
			Type:         organization.Type,
			PrefectureID: organization.PrefectureID,
			ParentID:     organization.ParentID,
			Description:  organization.Description,
			CreatedAt:    organization.CreatedAt,
			UpdatedAt:    organization.UpdatedAt,
		}

		response = append(response, resp)
	}

	h.l.InfoContext(ctx, "Successfully listed organizations", "count", len(response))
	c.JSON(http.StatusOK, response)
}

// GetOrganization @title 組織詳細取得
// @id GetOrganization
// @tags organization
// @accept json
// @produce json
// @Param id path int true "組織ID"
// @Summary 組織詳細取得
// @Success 200 {object} OrganizationResponse
// @Failure 404 {object} map[string]string
// @Router /organizations/{id} [get]
func (h *organizationHandler) GetOrganization(c *gin.Context) {
	idStr := c.Param("id")
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Invalid organization ID", "organization_id_str", idStr)
		c.JSON(400, gin.H{"error": "Invalid organization ID"})

		return
	}

	organization, err := h.organizationUseCase.GetOrganizationByID(ctx, int32(id))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Organization not found", "organization_id", id)
		c.JSON(404, gin.H{"error": "Organization not found"})

		return
	}

	// Map users to UserResponse
	var users []UserResponse
	for _, user := range organization.Users {
		users = append(users, UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	response := &OrganizationResponse{
		ID:           organization.ID,
		Name:         organization.Name,
		Type:         organization.Type,
		PrefectureID: organization.PrefectureID,
		ParentID:     organization.ParentID,
		Description:  organization.Description,
		CreatedAt:    organization.CreatedAt,
		UpdatedAt:    organization.UpdatedAt,
		Users:        users,
	}

	h.l.InfoContext(ctx, "Successfully retrieved organization", "organization_id", id)
	c.JSON(http.StatusOK, response)
}

// CreateOrganization @title 組織作成
// @id CreateOrganization
// @tags organization
// @accept json
// @produce json
// @Param request body CreateOrganizationRequest true "組織作成リクエスト"
// @Summary 組織作成
// @Success 201 {object} OrganizationResponse
// @Failure 400 {object} map[string]string
// @Router /organizations [post]
func (h *organizationHandler) CreateOrganization(c *gin.Context) {
	var req CreateOrganizationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	organization := &model.Organization{
		Name:         req.Name,
		Type:         req.Type,
		PrefectureID: req.PrefectureID,
		ParentID:     req.ParentID,
		Description:  req.Description,
	}

	err := h.organizationUseCase.CreateOrganization(c.Request.Context(), organization)
	if err != nil {
		h.l.ErrorContext(c.Request.Context(), err, "Failed to create organization")
		c.JSON(500, gin.H{"error": "Failed to create organization"})

		return
	}

	response := &OrganizationResponse{
		ID:           organization.ID,
		Name:         organization.Name,
		Type:         organization.Type,
		PrefectureID: organization.PrefectureID,
		ParentID:     organization.ParentID,
		Description:  organization.Description,
		CreatedAt:    organization.CreatedAt,
		UpdatedAt:    organization.UpdatedAt,
	}

	h.l.InfoContext(c.Request.Context(), "Successfully created organization", "organization_id", organization.ID)
	c.JSON(http.StatusCreated, response)
}

// UpdateOrganization @title 組織更新
// @id UpdateOrganization
// @tags organization
// @accept json
// @produce json
// @Param id path int true "組織ID"
// @Param request body UpdateOrganizationRequest true "組織更新リクエスト"
// @Summary 組織更新
// @Success 200 {object} OrganizationResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /organizations/{id} [put]
func (h *organizationHandler) UpdateOrganization(c *gin.Context) {
	idStr := c.Param("id")
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Invalid organization ID", "organization_id_str", idStr)
		c.JSON(400, gin.H{"error": "Invalid organization ID"})

		return
	}

	var req UpdateOrganizationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.ErrorContext(ctx, err, "Invalid request body")
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	// Check if the organization exists
	existingOrganization, err := h.organizationUseCase.GetOrganizationByID(ctx, int32(id))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Organization not found", "organization_id", id)
		c.JSON(404, gin.H{"error": "Organization not found"})

		return
	}

	// Update the organization
	existingOrganization.Name = req.Name
	existingOrganization.Type = req.Type
	existingOrganization.PrefectureID = req.PrefectureID
	existingOrganization.ParentID = req.ParentID
	existingOrganization.Description = req.Description

	err = h.organizationUseCase.UpdateOrganization(ctx, existingOrganization)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to update organization", "organization_id", id)
		c.JSON(500, gin.H{"error": "Failed to update organization"})

		return
	}

	response := &OrganizationResponse{
		ID:           existingOrganization.ID,
		Name:         existingOrganization.Name,
		Type:         existingOrganization.Type,
		PrefectureID: existingOrganization.PrefectureID,
		ParentID:     existingOrganization.ParentID,
		Description:  existingOrganization.Description,
		CreatedAt:    existingOrganization.CreatedAt,
		UpdatedAt:    existingOrganization.UpdatedAt,
	}

	h.l.InfoContext(ctx, "Successfully updated organization", "organization_id", id)
	c.JSON(http.StatusOK, response)
}

// DeleteOrganization @title 組織削除
// @id DeleteOrganization
// @tags organization
// @accept json
// @produce json
// @Param id path int true "組織ID"
// @Summary 組織削除
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /organizations/{id} [delete]
func (h *organizationHandler) DeleteOrganization(c *gin.Context) {
	idStr := c.Param("id")
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Invalid organization ID", "organization_id_str", idStr)
		c.JSON(400, gin.H{"error": "Invalid organization ID"})

		return
	}

	// Check if the organization exists
	_, err = h.organizationUseCase.GetOrganizationByID(ctx, int32(id))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Organization not found", "organization_id", id)
		c.JSON(404, gin.H{"error": "Organization not found"})

		return
	}

	err = h.organizationUseCase.DeleteOrganization(ctx, int32(id))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to delete organization", "organization_id", id)
		c.JSON(500, gin.H{"error": "Failed to delete organization"})

		return
	}

	h.l.InfoContext(ctx, "Successfully deleted organization", "organization_id", id)
	c.Status(http.StatusNoContent)
}
