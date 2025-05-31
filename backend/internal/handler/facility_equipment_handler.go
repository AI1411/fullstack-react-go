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

type FacilityEquipment interface {
	ListFacilityEquipments(c *gin.Context)
	GetFacilityEquipment(c *gin.Context)
	CreateFacilityEquipment(c *gin.Context)
	UpdateFacilityEquipment(c *gin.Context)
	DeleteFacilityEquipment(c *gin.Context)
}

type facilityEquipmentHandler struct {
	l                        *logger.Logger
	facilityEquipmentUseCase usecase.FacilityEquipmentUseCase
}

func NewFacilityEquipmentHandler(
	l *logger.Logger,
	facilityEquipmentUseCase usecase.FacilityEquipmentUseCase,
) FacilityEquipment {
	return &facilityEquipmentHandler{
		l:                        l,
		facilityEquipmentUseCase: facilityEquipmentUseCase,
	}
}

type FacilityEquipmentResponse struct {
	ID                  int32      `json:"id"`
	Name                string     `json:"name"`
	FacilityTypeID      int32      `json:"facility_type_id"`
	FacilityTypeName    string     `json:"facility_type_name,omitempty"`
	ModelNumber         *string    `json:"model_number,omitempty"`
	Manufacturer        *string    `json:"manufacturer,omitempty"`
	InstallationDate    *time.Time `json:"installation_date,omitempty"`
	Status              string     `json:"status"`
	LocationDescription *string    `json:"location_description,omitempty"`
	LocationLatitude    *float64   `json:"location_latitude,omitempty"`
	LocationLongitude   *float64   `json:"location_longitude,omitempty"`
	Notes               *string    `json:"notes,omitempty"`
}

type CreateFacilityEquipmentRequest struct {
	Name                string     `json:"name" binding:"required"`
	FacilityTypeID      int32      `json:"facility_type_id" binding:"required"`
	ModelNumber         *string    `json:"model_number"`
	Manufacturer        *string    `json:"manufacturer"`
	InstallationDate    *time.Time `json:"installation_date"`
	Status              string     `json:"status" binding:"required"`
	LocationDescription *string    `json:"location_description"`
	LocationLatitude    *float64   `json:"location_latitude"`
	LocationLongitude   *float64   `json:"location_longitude"`
	Notes               *string    `json:"notes"`
}

type UpdateFacilityEquipmentRequest struct {
	Name                string     `json:"name" binding:"required"`
	FacilityTypeID      int32      `json:"facility_type_id" binding:"required"`
	ModelNumber         *string    `json:"model_number"`
	Manufacturer        *string    `json:"manufacturer"`
	InstallationDate    *time.Time `json:"installation_date"`
	Status              string     `json:"status" binding:"required"`
	LocationDescription *string    `json:"location_description"`
	LocationLatitude    *float64   `json:"location_latitude"`
	LocationLongitude   *float64   `json:"location_longitude"`
	Notes               *string    `json:"notes"`
}

// ListFacilityEquipments @title 施設設備一覧取得
// @id ListFacilityEquipments
// @tags facility_equipment
// @accept json
// @produce json
// @Summary 施設設備一覧取得
// @Success 200 {array} FacilityEquipmentResponse
// @Router /facility-equipment [get]
func (h *facilityEquipmentHandler) ListFacilityEquipments(c *gin.Context) {
	ctx := c.Request.Context()
	facilityEquipments, err := h.facilityEquipmentUseCase.ListFacilityEquipments(ctx)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to list facility equipment")
		c.JSON(500, gin.H{"error": "Internal Server Error"})

		return
	}

	var response []*FacilityEquipmentResponse

	for _, facilityEquipment := range facilityEquipments {
		resp := &FacilityEquipmentResponse{
			ID:                  facilityEquipment.ID,
			Name:                facilityEquipment.Name,
			FacilityTypeID:      facilityEquipment.FacilityTypeID,
			ModelNumber:         facilityEquipment.ModelNumber,
			Manufacturer:        facilityEquipment.Manufacturer,
			InstallationDate:    facilityEquipment.InstallationDate,
			Status:              facilityEquipment.Status,
			LocationDescription: facilityEquipment.LocationDescription,
			LocationLatitude:    facilityEquipment.LocationLatitude,
			LocationLongitude:   facilityEquipment.LocationLongitude,
			Notes:               facilityEquipment.Notes,
		}

		response = append(response, resp)
	}

	h.l.InfoContext(ctx, "Successfully listed facility equipment", "count", len(response))
	c.JSON(http.StatusOK, response)
}

// GetFacilityEquipment @title 施設設備詳細取得
// @id GetFacilityEquipment
// @tags facility_equipment
// @accept json
// @produce json
// @Param id path int true "施設設備ID"
// @Summary 施設設備詳細取得
// @Success 200 {object} FacilityEquipmentResponse
// @Failure 404 {object} map[string]string
// @Router /facility-equipment/{id} [get]
func (h *facilityEquipmentHandler) GetFacilityEquipment(c *gin.Context) {
	idStr := c.Param("id")
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Invalid facility equipment ID", "facility_equipment_id_str", idStr)
		c.JSON(400, gin.H{"error": "Invalid facility equipment ID"})

		return
	}

	facilityEquipment, err := h.facilityEquipmentUseCase.GetFacilityEquipmentByID(ctx, int32(id))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Facility equipment not found", "facility_equipment_id", id)
		c.JSON(404, gin.H{"error": "Facility equipment not found"})

		return
	}

	response := &FacilityEquipmentResponse{
		ID:                  facilityEquipment.ID,
		Name:                facilityEquipment.Name,
		FacilityTypeID:      facilityEquipment.FacilityTypeID,
		ModelNumber:         facilityEquipment.ModelNumber,
		Manufacturer:        facilityEquipment.Manufacturer,
		InstallationDate:    facilityEquipment.InstallationDate,
		Status:              facilityEquipment.Status,
		LocationDescription: facilityEquipment.LocationDescription,
		LocationLatitude:    facilityEquipment.LocationLatitude,
		LocationLongitude:   facilityEquipment.LocationLongitude,
		Notes:               facilityEquipment.Notes,
	}

	h.l.InfoContext(ctx, "Successfully retrieved facility equipment", "facility_equipment_id", id)
	c.JSON(http.StatusOK, response)
}

// CreateFacilityEquipment @title 施設設備作成
// @id CreateFacilityEquipment
// @tags facility_equipment
// @accept json
// @produce json
// @Param request body CreateFacilityEquipmentRequest true "施設設備作成リクエスト"
// @Summary 施設設備作成
// @Success 201 {object} FacilityEquipmentResponse
// @Failure 400 {object} map[string]string
// @Router /facility-equipment [post]
func (h *facilityEquipmentHandler) CreateFacilityEquipment(c *gin.Context) {
	var req CreateFacilityEquipmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	facilityEquipment := &model.FacilityEquipment{
		Name:                req.Name,
		FacilityTypeID:      req.FacilityTypeID,
		ModelNumber:         req.ModelNumber,
		Manufacturer:        req.Manufacturer,
		InstallationDate:    req.InstallationDate,
		Status:              req.Status,
		LocationDescription: req.LocationDescription,
		LocationLatitude:    req.LocationLatitude,
		LocationLongitude:   req.LocationLongitude,
		Notes:               req.Notes,
	}

	err := h.facilityEquipmentUseCase.CreateFacilityEquipment(c.Request.Context(), facilityEquipment)
	if err != nil {
		h.l.ErrorContext(c.Request.Context(), err, "Failed to create facility equipment")
		c.JSON(500, gin.H{"error": "Failed to create facility equipment"})

		return
	}

	response := &FacilityEquipmentResponse{
		ID:                  facilityEquipment.ID,
		Name:                facilityEquipment.Name,
		FacilityTypeID:      facilityEquipment.FacilityTypeID,
		ModelNumber:         facilityEquipment.ModelNumber,
		Manufacturer:        facilityEquipment.Manufacturer,
		InstallationDate:    facilityEquipment.InstallationDate,
		Status:              facilityEquipment.Status,
		LocationDescription: facilityEquipment.LocationDescription,
		LocationLatitude:    facilityEquipment.LocationLatitude,
		LocationLongitude:   facilityEquipment.LocationLongitude,
		Notes:               facilityEquipment.Notes,
	}

	h.l.InfoContext(c.Request.Context(), "Successfully created facility equipment", "facility_equipment_id", facilityEquipment.ID)
	c.JSON(http.StatusCreated, response)
}

// UpdateFacilityEquipment @title 施設設備更新
// @id UpdateFacilityEquipment
// @tags facility_equipment
// @accept json
// @produce json
// @Param id path int true "施設設備ID"
// @Param request body UpdateFacilityEquipmentRequest true "施設設備更新リクエスト"
// @Summary 施設設備更新
// @Success 200 {object} FacilityEquipmentResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /facility-equipment/{id} [put]
func (h *facilityEquipmentHandler) UpdateFacilityEquipment(c *gin.Context) {
	idStr := c.Param("id")
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Invalid facility equipment ID", "facility_equipment_id_str", idStr)
		c.JSON(400, gin.H{"error": "Invalid facility equipment ID"})

		return
	}

	var req UpdateFacilityEquipmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.ErrorContext(ctx, err, "Invalid request body")
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	// Check if the facility equipment exists
	existingFacilityEquipment, err := h.facilityEquipmentUseCase.GetFacilityEquipmentByID(ctx, int32(id))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Facility equipment not found", "facility_equipment_id", id)
		c.JSON(404, gin.H{"error": "Facility equipment not found"})

		return
	}

	// Update the facility equipment
	existingFacilityEquipment.Name = req.Name
	existingFacilityEquipment.FacilityTypeID = req.FacilityTypeID
	existingFacilityEquipment.ModelNumber = req.ModelNumber
	existingFacilityEquipment.Manufacturer = req.Manufacturer
	existingFacilityEquipment.InstallationDate = req.InstallationDate
	existingFacilityEquipment.Status = req.Status
	existingFacilityEquipment.LocationDescription = req.LocationDescription
	existingFacilityEquipment.LocationLatitude = req.LocationLatitude
	existingFacilityEquipment.LocationLongitude = req.LocationLongitude
	existingFacilityEquipment.Notes = req.Notes

	err = h.facilityEquipmentUseCase.UpdateFacilityEquipment(ctx, existingFacilityEquipment)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to update facility equipment", "facility_equipment_id", id)
		c.JSON(500, gin.H{"error": "Failed to update facility equipment"})

		return
	}

	response := &FacilityEquipmentResponse{
		ID:                  existingFacilityEquipment.ID,
		Name:                existingFacilityEquipment.Name,
		FacilityTypeID:      existingFacilityEquipment.FacilityTypeID,
		ModelNumber:         existingFacilityEquipment.ModelNumber,
		Manufacturer:        existingFacilityEquipment.Manufacturer,
		InstallationDate:    existingFacilityEquipment.InstallationDate,
		Status:              existingFacilityEquipment.Status,
		LocationDescription: existingFacilityEquipment.LocationDescription,
		LocationLatitude:    existingFacilityEquipment.LocationLatitude,
		LocationLongitude:   existingFacilityEquipment.LocationLongitude,
		Notes:               existingFacilityEquipment.Notes,
	}

	h.l.InfoContext(ctx, "Successfully updated facility equipment", "facility_equipment_id", id)
	c.JSON(http.StatusOK, response)
}

// DeleteFacilityEquipment @title 施設設備削除
// @id DeleteFacilityEquipment
// @tags facility_equipment
// @accept json
// @produce json
// @Param id path int true "施設設備ID"
// @Summary 施設設備削除
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /facility-equipment/{id} [delete]
func (h *facilityEquipmentHandler) DeleteFacilityEquipment(c *gin.Context) {
	idStr := c.Param("id")
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Invalid facility equipment ID", "facility_equipment_id_str", idStr)
		c.JSON(400, gin.H{"error": "Invalid facility equipment ID"})

		return
	}

	// Check if the facility equipment exists
	_, err = h.facilityEquipmentUseCase.GetFacilityEquipmentByID(ctx, int32(id))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Facility equipment not found", "facility_equipment_id", id)
		c.JSON(404, gin.H{"error": "Facility equipment not found"})

		return
	}

	err = h.facilityEquipmentUseCase.DeleteFacilityEquipment(ctx, int32(id))
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to delete facility equipment", "facility_equipment_id", id)
		c.JSON(500, gin.H{"error": "Failed to delete facility equipment"})

		return
	}

	h.l.InfoContext(ctx, "Successfully deleted facility equipment", "facility_equipment_id", id)
	c.Status(http.StatusNoContent)
}
