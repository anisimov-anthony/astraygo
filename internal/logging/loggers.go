package logging

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// TODO: logger setup from env
func InitLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	return logger
}

func ZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Before request
		start := time.Now()

		// Process request
		ctx.Next()

		// After request
		latency := time.Since(start)

		logger.Info(
			"HTTP Request",

			// Writer data
			zap.Int("status", ctx.Writer.Status()),
			zap.Int("size", ctx.Writer.Size()),

			// Request data
			zap.String("method", ctx.Request.Method),
			zap.String("host", ctx.Request.URL.Host),
			zap.String("path", ctx.Request.URL.Path),
			zap.String("query", ctx.Request.URL.RawQuery),
			zap.String("proto", ctx.Request.Proto),
			zap.String("user-agent", ctx.Request.UserAgent()),

			// Custom
			zap.Duration("latency", latency),
		)
	}
}
