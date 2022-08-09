package logger

import (
	"github.com/samsoft00/golang-starter/service/lib/utils"
	"go.uber.org/zap"
)

// WithLogger can be used to compose a struct with a logger.
//
// Example:
//
//	type MyStruct struct {
//		*utils.WithLogger
//	}
type WithLogger struct {
	Logger *zap.SugaredLogger
}

func NewWithLogger(logger *zap.SugaredLogger) *WithLogger {
	return &WithLogger{Logger: logger}
}

func NewTestWithLogger() *WithLogger {
	logger := NewLogger(&utils.DefaultConfig{IsDebug: true})
	return &WithLogger{Logger: logger.Sugar()}
}

// NewSugaredLogger returns a new sugared logger
func NewSugaredLogger(config *utils.DefaultConfig) *zap.SugaredLogger {
	return NewLogger(config).Sugar()
}

// NewLogger returns a standard logger
func NewLogger(config *utils.DefaultConfig) *zap.Logger {
	var logger *zap.Logger

	// we use the config here so that we can use the same logger for tests
	if config.IsDebug {
		logger, _ = zap.NewDevelopment()
	} else {
		c := zap.NewProductionConfig()
		c.Encoding = "console"
		c.EncoderConfig = zap.NewDevelopmentEncoderConfig()
		logger, _ = c.Build()
	}

	return logger
}
