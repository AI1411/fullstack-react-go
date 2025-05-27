// Package logger provides a centralized way to configure and use slog throughout the application.
package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"time"
)

// LogLevel represents the logging level.
type LogLevel string

const (
	// DebugLevel logs debug messages.
	DebugLevel LogLevel = "debug"
	// InfoLevel logs informational messages.
	InfoLevel LogLevel = "info"
	// WarnLevel logs warning messages.
	WarnLevel LogLevel = "warn"
	// ErrorLevel logs error messages.
	ErrorLevel LogLevel = "error"
)

// Logger is a wrapper around slog.Logger that provides additional functionality.
type Logger struct {
	*slog.Logger
}

// Config holds the configuration for the logger.
type Config struct {
	// Level is the minimum log level that will be logged.
	Level LogLevel
	// Output is where the logs will be written to.
	Output io.Writer
	// AddSource adds the source file and line number to the log.
	AddSource bool
	// JSON determines whether the logs should be formatted as JSON.
	JSON bool
}

// DefaultConfig returns a default configuration for the logger.
func DefaultConfig() Config {
	return Config{
		Level:     InfoLevel,
		Output:    os.Stdout,
		AddSource: true,
		JSON:      true,
	}
}

// New creates a new Logger with the given configuration.
func New(cfg Config) *Logger {
	var level slog.Level
	switch cfg.Level {
	case DebugLevel:
		level = slog.LevelDebug
	case InfoLevel:
		level = slog.LevelInfo
	case WarnLevel:
		level = slog.LevelWarn
	case ErrorLevel:
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: cfg.AddSource,
	}

	if cfg.JSON {
		handler = slog.NewJSONHandler(cfg.Output, opts)
	} else {
		handler = slog.NewTextHandler(cfg.Output, opts)
	}

	return &Logger{
		Logger: slog.New(handler),
	}
}

// With returns a new Logger with the given attributes added to the context.
func (l *Logger) With(args ...any) *Logger {
	return &Logger{
		Logger: l.Logger.With(args...),
	}
}

// WithContext returns a new context with the logger attached.
func (l *Logger) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, loggerKey{}, l)
}

// FromContext returns the logger from the context.
// If no logger is found, it returns a default logger.
func FromContext(ctx context.Context) *Logger {
	if logger, ok := ctx.Value(loggerKey{}).(*Logger); ok {
		return logger
	}
	return &Logger{
		Logger: slog.Default(),
	}
}

// loggerKey is the key used to store the logger in the context.
type loggerKey struct{}

// LogRequest logs information about an HTTP request.
func (l *Logger) LogRequest(method, path string, statusCode int, latency time.Duration) {
	l.Info("request",
		slog.String("method", method),
		slog.String("path", path),
		slog.Int("status", statusCode),
		slog.Duration("latency", latency),
	)
}

// LogError logs an error with additional context.
func (l *Logger) LogError(err error, msg string, args ...any) {
	if err != nil {
		l.Error(msg, append(args, slog.Any("error", err))...)
	}
}
