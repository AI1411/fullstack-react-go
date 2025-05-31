package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/logger"
	"github.com/AI1411/fullstack-react-go/internal/usecase"
)

type SupportApplication interface {
	ListSupportApplications(c *gin.Context)
	GetSupportApplication(c *gin.Context)
	CreateSupportApplication(c *gin.Context)
}

type supportApplicationHandler struct {
	l                         *logger.Logger
	supportApplicationUseCase usecase.SupportApplicationUseCase
}

func NewSupportApplicationHandler(
	l *logger.Logger,
	supportApplicationUseCase usecase.SupportApplicationUseCase,
) SupportApplication {
	return &supportApplicationHandler{
		l:                         l,
		supportApplicationUseCase: supportApplicationUseCase,
	}
}

type SupportApplicationResponse struct {
	ApplicationID   string  `json:"application_id"`
	ApplicationDate string  `json:"application_date"`
	ApplicantName   string  `json:"applicant_name"`
	DisasterName    string  `json:"disaster_name"`
	RequestedAmount int64   `json:"requested_amount"`
	Status          string  `json:"status"`
	ReviewedAt      *string `json:"reviewed_at,omitempty"`
	ApprovedAt      *string `json:"approved_at,omitempty"`
	CompletedAt     *string `json:"completed_at,omitempty"`
	Notes           *string `json:"notes,omitempty"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

type ListSupportApplicationsResponse struct {
	SupportApplications []*SupportApplicationResponse `json:"support_applications"`
	Total               int64                         `json:"total"`
}

type CreateSupportApplicationRequest struct {
	ApplicationID   string  `json:"application_id" binding:"required"`
	ApplicationDate string  `json:"application_date" binding:"required"`
	ApplicantName   string  `json:"applicant_name" binding:"required"`
	DisasterName    string  `json:"disaster_name" binding:"required"`
	RequestedAmount int64   `json:"requested_amount" binding:"required"`
	Status          string  `json:"status"`
	Notes           *string `json:"notes"`
}

// ListSupportApplications @title 支援申請一覧取得
// @id ListSupportApplications
// @tags support-applications
// @accept json
// @produce json
// @version バージョン(1.0)
// @description
// @Summary 支援申請一覧取得
// @Success 200 {array} ListSupportApplicationsResponse
// @Router /support-applications [get]
func (h *supportApplicationHandler) ListSupportApplications(c *gin.Context) {
	ctx := c.Request.Context()
	supportApplications, err := h.supportApplicationUseCase.ListSupportApplications(ctx)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to list support applications")
		c.JSON(500, gin.H{"error": "Internal Server Error"})

		return
	}

	var sas []*SupportApplicationResponse

	for _, sa := range supportApplications {
		var reviewedAt, approvedAt, completedAt *string

		if sa.ReviewedAt != nil {
			formatted := sa.ReviewedAt.Format(time.DateTime)
			reviewedAt = &formatted
		}

		if sa.ApprovedAt != nil {
			formatted := sa.ApprovedAt.Format(time.DateTime)
			approvedAt = &formatted
		}

		if sa.CompletedAt != nil {
			formatted := sa.CompletedAt.Format(time.DateTime)
			completedAt = &formatted
		}

		sas = append(sas, &SupportApplicationResponse{
			ApplicationID:   sa.ApplicationID,
			ApplicationDate: sa.ApplicationDate.Format("2006-01-02"),
			ApplicantName:   sa.ApplicantName,
			DisasterName:    sa.DisasterName,
			RequestedAmount: sa.RequestedAmount,
			Status:          sa.Status,
			ReviewedAt:      reviewedAt,
			ApprovedAt:      approvedAt,
			CompletedAt:     completedAt,
			Notes:           sa.Notes,
			CreatedAt:       sa.CreatedAt.Format(time.DateTime),
			UpdatedAt:       sa.UpdatedAt.Format(time.DateTime),
		})
	}

	res := &ListSupportApplicationsResponse{
		SupportApplications: sas,
		Total:               int64(len(sas)),
	}

	c.JSON(http.StatusOK, res)
}

// GetSupportApplication @title 支援申請詳細取得
// @id GetSupportApplication
// @tags support-applications
// @accept json
// @produce json
// @Param id path string true "申請ID"
// @Summary 支援申請詳細取得
// @Success 200 {object} SupportApplicationResponse
// @Failure 404 {object} map[string]string
// @Router /support-applications/{id} [get]
func (h *supportApplicationHandler) GetSupportApplication(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	supportApplication, err := h.supportApplicationUseCase.GetSupportApplicationByID(ctx, id)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to get support application", "application_id", id)
		c.JSON(404, gin.H{"error": "Support application not found"})

		return
	}

	var reviewedAt, approvedAt, completedAt *string

	if supportApplication.ReviewedAt != nil {
		formatted := supportApplication.ReviewedAt.Format(time.DateTime)
		reviewedAt = &formatted
	}

	if supportApplication.ApprovedAt != nil {
		formatted := supportApplication.ApprovedAt.Format(time.DateTime)
		approvedAt = &formatted
	}

	if supportApplication.CompletedAt != nil {
		formatted := supportApplication.CompletedAt.Format(time.DateTime)
		completedAt = &formatted
	}

	response := &SupportApplicationResponse{
		ApplicationID:   supportApplication.ApplicationID,
		ApplicationDate: supportApplication.ApplicationDate.Format("2006-01-02"),
		ApplicantName:   supportApplication.ApplicantName,
		DisasterName:    supportApplication.DisasterName,
		RequestedAmount: supportApplication.RequestedAmount,
		Status:          supportApplication.Status,
		ReviewedAt:      reviewedAt,
		ApprovedAt:      approvedAt,
		CompletedAt:     completedAt,
		Notes:           supportApplication.Notes,
		CreatedAt:       supportApplication.CreatedAt.Format(time.DateTime),
		UpdatedAt:       supportApplication.UpdatedAt.Format(time.DateTime),
	}

	h.l.InfoContext(ctx, "Successfully retrieved support application", "application_id", id)
	c.JSON(http.StatusOK, response)
}

// CreateSupportApplication @title 支援申請作成
// @id CreateSupportApplication
// @tags support-applications
// @accept json
// @produce json
// @Param request body CreateSupportApplicationRequest true "支援申請作成リクエスト"
// @Summary 支援申請作成
// @Success 201 {object} SupportApplicationResponse
// @Failure 400 {object} map[string]string
// @Router /support-applications [post]
func (h *supportApplicationHandler) CreateSupportApplication(c *gin.Context) {
	var req CreateSupportApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Parse application_date
	applicationDate, err := time.Parse("2006-01-02", req.ApplicationDate)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid application_date format. Use YYYY-MM-DD format"})
		return
	}

	// Set default status if not provided
	status := req.Status
	if status == "" {
		status = "審査中"
	}

	supportApplication := &model.SupportApplication{
		ApplicationID:   req.ApplicationID,
		ApplicationDate: applicationDate,
		ApplicantName:   req.ApplicantName,
		DisasterName:    req.DisasterName,
		RequestedAmount: req.RequestedAmount,
		Status:          status,
		Notes:           req.Notes,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	err = h.supportApplicationUseCase.CreateSupportApplication(c.Request.Context(), supportApplication)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create support application"})
		return
	}

	response := &SupportApplicationResponse{
		ApplicationID:   supportApplication.ApplicationID,
		ApplicationDate: supportApplication.ApplicationDate.Format("2006-01-02"),
		ApplicantName:   supportApplication.ApplicantName,
		DisasterName:    supportApplication.DisasterName,
		RequestedAmount: supportApplication.RequestedAmount,
		Status:          supportApplication.Status,
		Notes:           supportApplication.Notes,
		CreatedAt:       supportApplication.CreatedAt.Format(time.DateTime),
		UpdatedAt:       supportApplication.UpdatedAt.Format(time.DateTime),
	}

	c.JSON(http.StatusCreated, response)
}
