package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/AI1411/fullstack-react-go/internal/usecase"
)

type Prefecture interface {
	ListPrefectures(c *gin.Context)
	GetPrefecture(c *gin.Context)
}

type prefectureHandler struct {
	prefectureUseCase usecase.PrefectureUseCase
}

func NewPrefectureHandler(
	prefectureUseCase usecase.PrefectureUseCase,
) Prefecture {
	return &prefectureHandler{
		prefectureUseCase: prefectureUseCase,
	}
}

type PrefectureResponse struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	RegionID int32  `json:"region_id"`
}

// ListPrefectures @title 都道府県一覧取得
// @id ListPrefectures
// @tags prefectures
// @accept json
// @produce json
// @Summary 都道府県一覧取得
// @Success 200 {array} PrefectureResponse
// @Router /prefectures [get]
func (h *prefectureHandler) ListPrefectures(c *gin.Context) {
	prefectures, err := h.prefectureUseCase.ListPrefectures(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	var response []*PrefectureResponse
	for _, prefecture := range prefectures {
		response = append(response, &PrefectureResponse{
			ID:       prefecture.ID,
			Name:     prefecture.Name,
			RegionID: prefecture.RegionID,
		})
	}

	c.JSON(http.StatusOK, response)
}

// GetPrefecture @title 都道府県詳細取得
// @id GetPrefecture
// @tags prefectures
// @accept json
// @produce json
// @Param id path int true "都道府県ID"
// @Summary 都道府県詳細取得
// @Success 200 {object} PrefectureResponse
// @Failure 404 {object} map[string]string
// @Router /prefectures/{id} [get]
func (h *prefectureHandler) GetPrefecture(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid prefecture ID"})
		return
	}

	prefecture, err := h.prefectureUseCase.GetPrefectureByID(c.Request.Context(), int32(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Prefecture not found"})
		return
	}

	response := &PrefectureResponse{
		ID:       prefecture.ID,
		Name:     prefecture.Name,
		RegionID: prefecture.RegionID,
	}

	c.JSON(http.StatusOK, response)
}
