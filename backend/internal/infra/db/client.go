package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	applogger "github.com/AI1411/fullstack-react-go/internal/infra/logger"
)

type SqlHandler struct {
	Conn *gorm.DB
}

// JSONLogger is a custom GORM logger that uses our application's JSON logger
type JSONLogger struct {
	Logger        *applogger.Logger
	SlowThreshold time.Duration
	LogLevel      logger.LogLevel
}

// NewJSONLogger creates a new JSONLogger
func NewJSONLogger(appLogger *applogger.Logger) *JSONLogger {
	return &JSONLogger{
		Logger:        appLogger,
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
	}
}

// LogMode sets the log level
func (l *JSONLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info logs info messages
func (l *JSONLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.Logger.Info(msg, "data", data)
	}
}

// Warn logs warn messages
func (l *JSONLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.Logger.Info(msg, "data", data)
	}
}

// Error logs error messages
func (l *JSONLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.Logger.LogError(errors.New(msg), msg, "data", data)
	}
}

// Trace logs SQL statements with execution time
func (l *JSONLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	if err != nil {
		l.Logger.LogError(err, "SQL error",
			"sql", sql,
			"rows", rows,
			"elapsed", elapsed)
		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		l.Logger.Info("SLOW SQL",
			"sql", sql,
			"rows", rows,
			"elapsed", elapsed)
		return
	}

	if l.LogLevel >= logger.Info {
		l.Logger.Info("SQL",
			"sql", sql,
			"rows", rows,
			"elapsed", elapsed)
	}
}

func NewSqlHandler(userID, password, host, port, dbName string, appLogger *applogger.Logger) (*SqlHandler, error) {
	// PostgreSQL形式の接続文字列
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		host, userID, password, dbName, port)

	// PostgreSQLに直接接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: NewJSONLogger(appLogger),
	})
	if err != nil {
		return nil, err
	}

	// 接続プールの設定
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 接続プールの設定
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &SqlHandler{
		Conn: db,
	}, nil
}
