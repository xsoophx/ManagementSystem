package util

import (
	"context"
	"go.uber.org/zap"
)

type ctxKey string

const loggerKey = ctxKey("logger")

func GetLogger(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(loggerKey).(*zap.Logger)
	if !ok || logger == nil {
		fallbackLogger, _ := zap.NewDevelopment()
		return fallbackLogger
	}
	return logger
}
