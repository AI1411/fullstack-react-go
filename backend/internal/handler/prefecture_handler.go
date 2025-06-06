package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AI1411/fullstack-react-go/internal/infra/logger"
	"github.com/AI1411/fullstack-react-go/internal/usecase"
)

type Prefecture interface {
	ListPrefectures(c *gin.Context)
	GetPrefecture(c *gin.Context)
}

type prefectureHandler struct {
	l                 *logger.Logger
	prefectureUseCase usecase.PrefectureUseCase
}

func NewPrefectureHandler(
	l *logger.Logger,
	prefectureUseCase usecase.PrefectureUseCase,
) Prefecture {
	return &prefectureHandler{
		l:                 l,
		prefectureUseCase: prefectureUseCase,
	}
}

type PrefectureResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
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
	ctx := c.Request.Context()
	prefectures, err := h.prefectureUseCase.ListPrefectures(ctx)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Failed to list prefectures")
		c.JSON(500, gin.H{"error": "Internal Server Error"})

		return
	}

	var response []*PrefectureResponse
	for _, prefecture := range prefectures {
		response = append(response, &PrefectureResponse{
			ID:   prefecture.ID,
			Name: prefecture.Name,
		})
	}

	h.l.InfoContext(ctx, "Successfully listed prefectures", "count", len(response))
	c.JSON(http.StatusOK, response)
}

type GetPrefectureResponse struct {
	ID             int32           `json:"id"`
	Name           string          `json:"name"`
	Municipalities []*Municipality `json:"municipalities"`
}

type Municipality struct {
	ID                    int32  `json:"id"`
	PrefectureCode        string `json:"prefecture_code"`
	OrganizationCode      string `json:"organization_code"`
	PrefectureNameKanji   string `json:"prefecture_name_kanji"`
	MunicipalityNameKanji string `json:"municipality_name_kanji"`
	PrefectureNameKana    string `json:"prefecture_name_kana"`
	MunicipalityNameKana  string `json:"municipality_name_kana"`
	IsActive              bool   `json:"is_active"`
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
	ctx := c.Request.Context()
	code := c.Param("code")

	prefecture, err := h.prefectureUseCase.GetPrefectureByID(ctx, code)
	if err != nil {
		h.l.ErrorContext(ctx, err, "Prefecture not found", "prefecture_code", code)
		c.JSON(404, gin.H{"error": "Prefecture not found"})

		return
	}

	response := &GetPrefectureResponse{
		ID:             prefecture.ID,
		Name:           prefecture.Name,
		Municipalities: make([]*Municipality, len(prefecture.Municipalities)),
	}

	for i, municipality := range prefecture.Municipalities {
		response.Municipalities[i] = &Municipality{
			ID:                    municipality.ID,
			PrefectureCode:        municipality.PrefectureCode,
			OrganizationCode:      municipality.OrganizationCode,
			PrefectureNameKanji:   municipality.PrefectureNameKanji,
			MunicipalityNameKanji: municipality.MunicipalityNameKanji,
			PrefectureNameKana:    municipality.PrefectureNameKana,
			MunicipalityNameKana:  municipality.MunicipalityNameKana,
			IsActive:              municipality.IsActive,
		}
	}

	c.JSON(http.StatusOK, response)
}
