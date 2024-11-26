package infrastructurelogging

import (
	"fmt"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger(logger *zap.Logger) *Logger {
	return &Logger{logger}
}

func (l *Logger) Info(process, correlationID, step, key string, val any) {
	l.logger.Info(process, zap.String("correlationID", correlationID), zap.String("step", step), zap.String("cid", correlationID), zap.String("step", step), zap.String(key, fmt.Sprint(val)))
}

func (l *Logger) Error(process, correlationID, step, key string, val any) {
	l.logger.Error(process, zap.String("correlationID", correlationID), zap.String("step", step), zap.String("cid", correlationID), zap.String("step", step), zap.String(key, fmt.Sprint(val)))
}
