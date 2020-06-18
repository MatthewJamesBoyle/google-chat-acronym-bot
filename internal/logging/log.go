package logging

import (
	"context"
	"go.uber.org/zap"
)

type loggerKey struct{}

var defaultLogger *zap.SugaredLogger

func init() {

	c := zap.NewProductionConfig()
	c.EncoderConfig.MessageKey = "message"
	c.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	if l, err := c.Build(zap.WithCaller(true)); err != nil {
		defaultLogger = &zap.SugaredLogger{}
	} else {
		defaultLogger = l.Named("default").Sugar()
	}
	defer defaultLogger.Sync()
}

func WithLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func FromContext(ctx context.Context) *zap.SugaredLogger {
	if logger, ok := ctx.Value(loggerKey{}).(*zap.SugaredLogger); ok {
		return logger
	}
	return defaultLogger
}
