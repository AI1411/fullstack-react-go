package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AI1411/fullstack-react-go/internal/infra/logger"
	"github.com/AI1411/fullstack-react-go/internal/usecase"
)

type Timeline interface {
	GetTimelinesByDisasterID(c *gin.Context)
}

type timelineHandler struct {
	logger          *logger.Logger
	timelineUseCase usecase.TimelineUseCase
}

func NewTimelineHandler(
	logger *logger.Logger,
	timelineUseCase usecase.TimelineUseCase,
) Timeline {
	return &timelineHandler{
		logger:          logger,
		timelineUseCase: timelineUseCase,
	}
}

type TimelineResponse struct {
	ID          int32  `json:"id"`
	DisasterID  string `json:"disaster_id"`
	EventName   string `json:"event_name"`
	EventTime   string `json:"event_time"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
}

type ListTimelinesResponse struct {
	Timelines []*TimelineResponse `json:"timelines"`
}

// GetTimelinesByDisasterID godoc
// @Summary Get timelines by disaster ID
// @Description Get timelines by disaster ID
// @Tags timelines
// @Accept json
// @Produce json
// @Param id path string true "Disaster ID"
// @Success 200 {object} ListTimelinesResponse
// @Router /disasters/{id}/timelines [get]
func (h *timelineHandler) GetTimelinesByDisasterID(c *gin.Context) {
	disasterID := c.Param("id")
	if disasterID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "disaster_id is required",
		})
		return
	}

	timelines, err := h.timelineUseCase.GetTimelinesByDisasterID(c.Request.Context(), disasterID)
	if err != nil {
		h.logger.Error("failed to get timelines", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get timelines",
		})
		return
	}

	if len(timelines) == 0 {
		c.JSON(http.StatusOK, ListTimelinesResponse{
			Timelines: []*TimelineResponse{},
		})
		return
	}

	var response []*TimelineResponse
	for _, timeline := range timelines {
		var severity string
		if timeline.Severity != nil {
			severity = *timeline.Severity
		}

		response = append(response, &TimelineResponse{
			ID:          timeline.ID,
			DisasterID:  timeline.DisasterID,
			EventName:   timeline.EventName,
			EventTime:   timeline.EventTime.Format("2006-01-02 15:04:05"),
			Description: timeline.Description,
			Severity:    severity,
		})
	}

	c.JSON(http.StatusOK, ListTimelinesResponse{
		Timelines: response,
	})
}
