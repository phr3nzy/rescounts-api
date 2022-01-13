// package logger is used to create a performant JSON logger with extra configuration.
// It will be mainly used for HTTP request logging.
package logger

import (
	"strings"

	"github.com/phr3nzy/rescounts-api/internals/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logger JSON based STDOUT logger.
var logger *zap.Logger

// createLoggerInstance creates an instance of our JSON logger
func createLoggerInstance() *zap.Logger {
	cfg := config.GetConfig()
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.OutputPaths = []string{"stdout"}
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.MessageKey = "message"
	loggerConfig.DisableStacktrace = true

	switch strings.ToLower(cfg.LogLevel) {
	case "silent":
	case "trace":
	case "debug":
		{
			loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
			break
		}
	case "info":
		{
			loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
			break
		}
	case "warn":
		{
			loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
			break
		}
	case "error":
		{
			loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
			break
		}
	case "fatal":
		{
			loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.FatalLevel)
			break
		}
	case "default":
		{
			loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
			break
		}
	}
	loggerConfig.InitialFields = map[string]interface{}{
		"service": cfg.ServiceName,
	}

	log, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}

	return log
}

// GetLoggerInstance returns the active instance of the JSON logger
func GetLoggerInstance() *zap.Logger {
	if logger == nil {
		logger = createLoggerInstance()
	}

	return logger
}
