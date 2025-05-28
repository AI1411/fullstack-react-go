package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/AI1411/fullstack-react-go/internal/infra/logger"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func NewLogging(logger *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		endpoint := c.Request.RequestURI
		if endpoint == "/metrics" {
			c.Next()
			return
		}

		var traceID string
		traceID = c.GetHeader("X-Trace-ID")

		if traceID == "" {
			traceID = getTraceID(c.Request.Context())
		}

		ctx := context.WithValue(c.Request.Context(), "trace_id", traceID)
		c.Request = c.Request.WithContext(ctx)

		c.Set("trace_id", traceID)
		bufBody, _ := io.ReadAll(c.Request.Body)
		reqBody := map[string]any{}
		_ = json.Unmarshal(bufBody, &reqBody)

		c.Request.Body = io.NopCloser(bytes.NewBuffer(bufBody))

		logger.Info("request",
			slog.String("trace_id", traceID),
			slog.String("http_method", c.Request.Method),
			slog.String("endpoint", endpoint),
			slog.Any("header", c.Request.Header),
			slog.Any("body", reqBody),
		)

		writer := &responseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = writer
		defer func() {
			logger.Info("response",
				slog.String("trace_id", traceID),
				slog.String("endpoint", endpoint),
				slog.Any("header", writer.Header()),
				slog.Int("http_status", writer.Status()),
				slog.Any("body", string(writer.body.Bytes())),
			)
		}()

		c.Next()
	}
}

func getTraceID(ctx context.Context) string {
	traceID, ok := ctx.Value("trace_id").(string)
	if !ok {
		tid, _ := uuid.NewRandom()
		traceID = tid.String()
		return traceID
	}
	return traceID
}
