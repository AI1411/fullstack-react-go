package handler

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/AI1411/fullstack-react-go/internal/infra/logger"
)

// File is the interface for file operations
type File interface {
	UploadDisasterImage(c *gin.Context)
}

type fileHandler struct {
	logger *logger.Logger
}

// NewFileHandler creates a new file handler
func NewFileHandler(l *logger.Logger) File {
	return &fileHandler{
		logger: l,
	}
}

// UploadDisasterImage handles the upload of disaster images
func (h *fileHandler) UploadDisasterImage(c *gin.Context) {
	// Get the disaster ID from the URL parameter
	disasterID := c.Param("id")
	if disasterID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Disaster ID is required",
		})
		return
	}

	// Get the file from the form
	file, err := c.FormFile("image")
	if err != nil {
		h.logger.Error("Failed to get file from form", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get file from form",
		})
		return
	}

	// Generate a unique filename
	ext := filepath.Ext(file.Filename)
	filename := uuid.New().String() + ext

	// Create the upload directory if it doesn't exist
	uploadDir := "uploads/disasters/" + disasterID
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		h.logger.Error("Failed to create upload directory", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create upload directory",
		})
		return
	}

	// Save the file
	dst := filepath.Join(uploadDir, filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		h.logger.Error("Failed to save file", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save file",
		})
		return
	}

	// Return the file information
	c.JSON(http.StatusOK, gin.H{
		"filename": filename,
		"path":     dst,
		"size":     file.Size,
		"uploaded": time.Now(),
	})
}
