package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/logger"
	"github.com/AI1411/fullstack-react-go/internal/usecase"
)

type Disaster interface {
	ListDisasters(c *gin.Context)
	GetDisaster(c *gin.Context)
	CreateDisaster(c *gin.Context)
	UpdateDisaster(c *gin.Context)
	DeleteDisaster(c *gin.Context)
}

type disasterHandler struct {
	l               *logger.Logger
	disasterUseCase usecase.DisasterUseCase
}

func NewDisasterHandler(
	l *logger.Logger,
	disasterUseCase usecase.DisasterUseCase,
) Disaster {
	return &disasterHandler{
		l:               l,
		disasterUseCase: disasterUseCase,
	}
}

type PrefectureItem struct {
	Name   string     `json:"name" binding:"required"`
	Region RegionItem `json:"region" binding:"required"`
}

type RegionItem struct {
	Name string `json:"name" binding:"required"`
}

type DisasterResponse struct {
	ID                    string         `json:"id"`
	DisasterCode          string         `json:"disaster_code"`
	Name                  string         `json:"name"`
	Prefecture            PrefectureItem `json:"prefecture"`
	OccurredAt            string         `json:"occurred_at"`
	Summary               string         `json:"summary"`
	DisasterType          string         `json:"disaster_type"`
	Status                string         `json:"status"`
	ImpactLevel           string         `json:"impact_level"`
	AffectedAreaSize      *float64       `json:"affected_area_size"`
	EstimatedDamageAmount *float64       `json:"estimated_damage_amount"`
}

type ListDisastersResponse struct {
	Disasters []*DisasterResponse `json:"disasters"`
	Total     int64               `json:"total"`
}

type CreateDisasterRequest struct {
	DisasterCode          string   `json:"disaster_code" binding:"required"`
	Name                  string   `json:"name" binding:"required"`
	PrefectureID          int32    `json:"prefecture_id" binding:"required"`
	OccurredAt            string   `json:"occurred_at" binding:"required"`
	Summary               string   `json:"summary" binding:"required"`
	DisasterType          string   `json:"disaster_type" binding:"required"`
	Status                string   `json:"status"`
	ImpactLevel           string   `json:"impact_level" binding:"required"`
	AffectedAreaSize      *float64 `json:"affected_area_size"`
	EstimatedDamageAmount *float64 `json:"estimated_damage_amount"`
}

type UpdateDisasterRequest struct {
	DisasterCode          string   `json:"disaster_code"`
	Name                  string   `json:"name"`
	PrefectureID          int32    `json:"prefecture_id"`
	OccurredAt            string   `json:"occurred_at"`
	Summary               string   `json:"summary"`
	DisasterType          string   `json:"disaster_type"`
	Status                string   `json:"status"`
	ImpactLevel           string   `json:"impact_level"`
	AffectedAreaSize      *float64 `json:"affected_area_size"`
	EstimatedDamageAmount *float64 `json:"estimated_damage_amount"`
}

// ListDisasters @title 災害マスタ一覧取得
// @id ListDisasters
// @tags disasters
// @accept json
// @produce json
// @version バージョン(1.0)
// @description
// @Summary 災害マスタ一覧取得
// @Success 200 {array} ListDisastersResponse
// @Router /disasters [get]
func (h *disasterHandler) ListDisasters(c *gin.Context) {
	ctx := c.Request.Context()
	disasters, err := h.disasterUseCase.ListDisasters(ctx)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to list disasters")
		c.JSON(500, gin.H{"error": "Internal Server Error"})

		return
	}

	var ds []*DisasterResponse
	for _, disaster := range disasters {
		ds = append(ds, &DisasterResponse{
			ID:           disaster.ID,
			DisasterCode: disaster.DisasterCode,
			Name:         disaster.Name,
			Prefecture: PrefectureItem{
				Name: disaster.Prefecture.Name,
			},
			OccurredAt:            disaster.OccurredAt.Format(time.DateTime),
			Summary:               disaster.Summary,
			DisasterType:          disaster.DisasterType,
			Status:                disaster.Status,
			ImpactLevel:           disaster.ImpactLevel,
			AffectedAreaSize:      disaster.AffectedAreaSize,
			EstimatedDamageAmount: disaster.EstimatedDamageAmount,
		})
	}

	res := &ListDisastersResponse{
		Disasters: ds,
		Total:     int64(len(ds)),
	}

	c.JSON(http.StatusOK, res)
}

// GetDisaster @title 災害詳細取得
// @id GetDisaster
// @tags disasters
// @accept json
// @produce json
// @Param id path string true "災害ID"
// @Summary 災害詳細取得
// @Success 200 {object} ListDisastersResponse
// @Failure 404 {object} map[string]string
// @Router /disasters/{id} [get]
func (h *disasterHandler) GetDisaster(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	disaster, err := h.disasterUseCase.GetDisasterByID(ctx, id)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to get disaster", "disaster_id", id)
		c.JSON(404, gin.H{"error": "Disaster not found"})

		return
	}

	response := &DisasterResponse{
		ID:                    disaster.ID,
		DisasterCode:          disaster.DisasterCode,
		Name:                  disaster.Name,
		OccurredAt:            disaster.OccurredAt.Format(time.DateTime),
		Summary:               disaster.Summary,
		DisasterType:          disaster.DisasterType,
		Status:                disaster.Status,
		ImpactLevel:           disaster.ImpactLevel,
		AffectedAreaSize:      disaster.AffectedAreaSize,
		EstimatedDamageAmount: disaster.EstimatedDamageAmount,
	}

	h.l.InfoContext(ctx, "Successfully retrieved disaster", "disaster_id", id)
	c.JSON(http.StatusOK, response)
}

// CreateDisaster @title 災害作成
// @id CreateDisaster
// @tags disasters
// @accept json
// @produce json
// @Param request body CreateDisasterRequest true "災害作成リクエスト"
// @Summary 災害作成
// @Success 201 {object} ListDisastersResponse
// @Failure 400 {object} map[string]string
// @Router /disasters [post]
func (h *disasterHandler) CreateDisaster(c *gin.Context) {
	var req CreateDisasterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Parse occurred_at
	occurredAt, err := time.Parse(time.RFC3339, req.OccurredAt)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid occurred_at format. Use RFC3339 format"})
		return
	}

	// Set default status if not provided
	status := req.Status
	if status == "" {
		status = "pending"
	}

	disaster := &model.Disaster{
		DisasterCode:          req.DisasterCode,
		Name:                  req.Name,
		PrefectureID:          req.PrefectureID,
		OccurredAt:            occurredAt,
		Summary:               req.Summary,
		DisasterType:          req.DisasterType,
		Status:                status,
		ImpactLevel:           req.ImpactLevel,
		AffectedAreaSize:      req.AffectedAreaSize,
		EstimatedDamageAmount: req.EstimatedDamageAmount,
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
	}

	err = h.disasterUseCase.CreateDisaster(c.Request.Context(), disaster)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create disaster"})
		return
	}

	response := &DisasterResponse{
		ID:                    disaster.ID,
		DisasterCode:          disaster.DisasterCode,
		Name:                  disaster.Name,
		OccurredAt:            disaster.OccurredAt.Format(time.DateTime),
		Summary:               disaster.Summary,
		DisasterType:          disaster.DisasterType,
		Status:                disaster.Status,
		ImpactLevel:           disaster.ImpactLevel,
		AffectedAreaSize:      disaster.AffectedAreaSize,
		EstimatedDamageAmount: disaster.EstimatedDamageAmount,
	}

	c.JSON(http.StatusCreated, response)
}

// UpdateDisaster @title 災害更新
// @id UpdateDisaster
// @tags disasters
// @accept json
// @produce json
// @Param id path string true "災害ID"
// @Param request body UpdateDisasterRequest true "災害更新リクエスト"
// @Summary 災害更新
// @Success 200 {object} ListDisastersResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /disasters/{id} [put]
func (h *disasterHandler) UpdateDisaster(c *gin.Context) {
	id := c.Param("id")

	var req UpdateDisasterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Get existing disaster
	disaster, err := h.disasterUseCase.GetDisasterByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Disaster not found"})
		return
	}

	// Update fields if provided
	if req.DisasterCode != "" {
		disaster.DisasterCode = req.DisasterCode
	}

	if req.Name != "" {
		disaster.Name = req.Name
	}

	if req.PrefectureID != 0 {
		disaster.PrefectureID = req.PrefectureID
	}

	if req.OccurredAt != "" {
		occurredAt, err := time.Parse(time.RFC3339, req.OccurredAt)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid occurred_at format. Use RFC3339 format"})
			return
		}

		disaster.OccurredAt = occurredAt
	}

	if req.Summary != "" {
		disaster.Summary = req.Summary
	}

	if req.DisasterType != "" {
		disaster.DisasterType = req.DisasterType
	}

	if req.Status != "" {
		disaster.Status = req.Status
	}

	if req.ImpactLevel != "" {
		disaster.ImpactLevel = req.ImpactLevel
	}

	if req.AffectedAreaSize != nil {
		disaster.AffectedAreaSize = req.AffectedAreaSize
	}

	if req.EstimatedDamageAmount != nil {
		disaster.EstimatedDamageAmount = req.EstimatedDamageAmount
	}

	disaster.UpdatedAt = time.Now()

	err = h.disasterUseCase.UpdateDisaster(c.Request.Context(), disaster)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update disaster"})
		return
	}

	response := &DisasterResponse{
		ID:                    disaster.ID,
		DisasterCode:          disaster.DisasterCode,
		Name:                  disaster.Name,
		OccurredAt:            disaster.OccurredAt.Format(time.DateTime),
		Summary:               disaster.Summary,
		DisasterType:          disaster.DisasterType,
		Status:                disaster.Status,
		ImpactLevel:           disaster.ImpactLevel,
		AffectedAreaSize:      disaster.AffectedAreaSize,
		EstimatedDamageAmount: disaster.EstimatedDamageAmount,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteDisaster @title 災害削除
// @id DeleteDisaster
// @tags disasters
// @Param id path string true "災害ID"
// @Summary 災害削除
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /disasters/{id} [delete]
func (h *disasterHandler) DeleteDisaster(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	// Check if disaster exists
	_, err := h.disasterUseCase.GetDisasterByID(ctx, id)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Disaster not found for deletion", "disaster_id", id)
		c.JSON(404, gin.H{"error": "Disaster not found"})

		return
	}

	err = h.disasterUseCase.DeleteDisaster(ctx, id)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to delete disaster", "disaster_id", id)
		c.JSON(500, gin.H{"error": "Failed to delete disaster"})

		return
	}

	c.Status(http.StatusNoContent)
}
