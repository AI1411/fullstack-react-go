package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/logger"
	"github.com/AI1411/fullstack-react-go/internal/usecase"
)

type DamageLevel interface {
	ListDamageLevels(c *gin.Context)
	GetDamageLevel(c *gin.Context)
	CreateDamageLevel(c *gin.Context)
	UpdateDamageLevel(c *gin.Context)
	DeleteDamageLevel(c *gin.Context)
}

type damageLevelHandler struct {
	l                  *logger.Logger
	damageLevelUseCase usecase.DamageLevelUseCase
}

func NewDamageLevelHandler(
	l *logger.Logger,
	damageLevelUseCase usecase.DamageLevelUseCase,
) DamageLevel {
	return &damageLevelHandler{
		l:                  l,
		damageLevelUseCase: damageLevelUseCase,
	}
}

type DamageLevelResponse struct {
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type CreateDamageLevelRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
}

type UpdateDamageLevelRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

// ListDamageLevels @title 被害程度一覧取得
// @id ListDamageLevels
// @tags damage_levels
// @accept json
// @produce json
// @Summary 被害程度一覧取得
// @Success 200 {array} DamageLevelResponse
// @Router /damage-levels [get]
func (h *damageLevelHandler) ListDamageLevels(c *gin.Context) {
	ctx := c.Request.Context()
	damageLevels, err := h.damageLevelUseCase.ListDamageLevels(ctx)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to list damage levels")
		c.JSON(500, gin.H{"error": "Internal Server Error"})

		return
	}

	var response []*DamageLevelResponse
	for _, damageLevel := range damageLevels {
		response = append(response, &DamageLevelResponse{
			ID:          damageLevel.ID,
			Name:        damageLevel.Name,
			Description: damageLevel.Description,
		})
	}

	h.l.InfoContext(ctx, "Successfully listed damage levels", "count", len(response))
	c.JSON(http.StatusOK, response)
}

// GetDamageLevel @title 被害程度詳細取得
// @id GetDamageLevel
// @tags damage_levels
// @accept json
// @produce json
// @Param id path int true "被害程度ID"
// @Summary 被害程度詳細取得
// @Success 200 {object} DamageLevelResponse
// @Failure 404 {object} map[string]string
// @Router /damage-levels/{id} [get]
func (h *damageLevelHandler) GetDamageLevel(c *gin.Context) {
	idStr := c.Param("id")
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Invalid damage level ID", "damage_level_id_str", idStr)
		c.JSON(400, gin.H{"error": "Invalid damage level ID"})

		return
	}

	damageLevel, err := h.damageLevelUseCase.GetDamageLevelByID(ctx, int32(id))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Damage level not found", "damage_level_id", id)
		c.JSON(404, gin.H{"error": "Damage level not found"})

		return
	}

	response := &DamageLevelResponse{
		ID:          damageLevel.ID,
		Name:        damageLevel.Name,
		Description: damageLevel.Description,
	}

	h.l.InfoContext(ctx, "Successfully retrieved damage level", "damage_level_id", id)
	c.JSON(http.StatusOK, response)
}

// CreateDamageLevel @title 被害程度作成
// @id CreateDamageLevel
// @tags damage_levels
// @accept json
// @produce json
// @Param request body CreateDamageLevelRequest true "被害程度作成リクエスト"
// @Summary 被害程度作成
// @Success 201 {object} DamageLevelResponse
// @Failure 400 {object} map[string]string
// @Router /damage-levels [post]
func (h *damageLevelHandler) CreateDamageLevel(c *gin.Context) {
	var req CreateDamageLevelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	damageLevel := &model.DamageLevel{
		Name:        req.Name,
		Description: req.Description,
	}

	err := h.damageLevelUseCase.CreateDamageLevel(c.Request.Context(), damageLevel)
	if err != nil {
		h.l.ErrorContext(c.Request.Context(), err, "Failed to create damage level")
		c.JSON(500, gin.H{"error": "Failed to create damage level"})
		return
	}

	response := &DamageLevelResponse{
		ID:          damageLevel.ID,
		Name:        damageLevel.Name,
		Description: damageLevel.Description,
	}

	h.l.InfoContext(c.Request.Context(), "Successfully created damage level", "damage_level_id", damageLevel.ID)
	c.JSON(http.StatusCreated, response)
}

// UpdateDamageLevel @title 被害程度更新
// @id UpdateDamageLevel
// @tags damage_levels
// @accept json
// @produce json
// @Param id path int true "被害程度ID"
// @Param request body UpdateDamageLevelRequest true "被害程度更新リクエスト"
// @Summary 被害程度更新
// @Success 200 {object} DamageLevelResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /damage-levels/{id} [put]
func (h *damageLevelHandler) UpdateDamageLevel(c *gin.Context) {
	idStr := c.Param("id")
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Invalid damage level ID", "damage_level_id_str", idStr)
		c.JSON(400, gin.H{"error": "Invalid damage level ID"})
		return
	}

	var req UpdateDamageLevelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Get existing damage level
	damageLevel, err := h.damageLevelUseCase.GetDamageLevelByID(ctx, int32(id))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Damage level not found", "damage_level_id", id)
		c.JSON(404, gin.H{"error": "Damage level not found"})
		return
	}

	// Update fields if provided
	if req.Name != "" {
		damageLevel.Name = req.Name
	}

	if req.Description != nil {
		damageLevel.Description = req.Description
	}

	err = h.damageLevelUseCase.UpdateDamageLevel(ctx, damageLevel)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to update damage level", "damage_level_id", id)
		c.JSON(500, gin.H{"error": "Failed to update damage level"})
		return
	}

	response := &DamageLevelResponse{
		ID:          damageLevel.ID,
		Name:        damageLevel.Name,
		Description: damageLevel.Description,
	}

	h.l.InfoContext(ctx, "Successfully updated damage level", "damage_level_id", id)
	c.JSON(http.StatusOK, response)
}

// DeleteDamageLevel @title 被害程度削除
// @id DeleteDamageLevel
// @tags damage_levels
// @Param id path int true "被害程度ID"
// @Summary 被害程度削除
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /damage-levels/{id} [delete]
func (h *damageLevelHandler) DeleteDamageLevel(c *gin.Context) {
	idStr := c.Param("id")
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Invalid damage level ID", "damage_level_id_str", idStr)
		c.JSON(400, gin.H{"error": "Invalid damage level ID"})
		return
	}

	// Check if damage level exists
	_, err = h.damageLevelUseCase.GetDamageLevelByID(ctx, int32(id))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Damage level not found for deletion", "damage_level_id", id)
		c.JSON(404, gin.H{"error": "Damage level not found"})
		return
	}

	err = h.damageLevelUseCase.DeleteDamageLevel(ctx, int32(id))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to delete damage level", "damage_level_id", id)
		c.JSON(500, gin.H{"error": "Failed to delete damage level"})
		return
	}

	h.l.InfoContext(ctx, "Successfully deleted damage level", "damage_level_id", id)
	c.Status(http.StatusNoContent)
}
