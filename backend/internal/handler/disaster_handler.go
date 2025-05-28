package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/AI1411/fullstack-react-go/internal/usecase"
)

type Disaster interface {
	ListDisasters(c *gin.Context)
}

type disasterHandler struct {
	disasterUseCase usecase.DisasterUseCase
}

func NewDisasterHandler(
	disasterUseCase usecase.DisasterUseCase,
) Disaster {
	return &disasterHandler{
		disasterUseCase: disasterUseCase,
	}
}

type ListDisastersResponse struct {
	ID                    string   `json:"id"`
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
// @Success 200 {object} ListDisastersResponse
// @Router /disasters [get]
func (h *disasterHandler) ListDisasters(c *gin.Context) {
	disasters, err := h.disasterUseCase.ListDisasters(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	var response []*ListDisastersResponse
	for _, disaster := range disasters {
		response = append(response, &ListDisastersResponse{
			ID:                    disaster.ID,
			DisasterCode:          disaster.DisasterCode,
			Name:                  disaster.Name,
			PrefectureID:          disaster.PrefectureID,
			OccurredAt:            disaster.OccurredAt.Format(time.DateTime),
			Summary:               disaster.Summary,
			DisasterType:          disaster.DisasterType,
			Status:                disaster.Status,
			ImpactLevel:           disaster.ImpactLevel,
			AffectedAreaSize:      disaster.AffectedAreaSize,
			EstimatedDamageAmount: disaster.EstimatedDamageAmount,
		})
	}

	c.JSON(http.StatusOK, response)
}
