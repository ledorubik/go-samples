package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func loggingMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Starting time of request
		startTime := time.Now()

		// Processing request
		ctx.Next()

		// End time of request
		endTime := time.Now()

		// Execution time
		latencyTime := endTime.Sub(startTime).String()

		// Request method
		reqMethod := ctx.Request.Method

		// Request route
		reqUri := ctx.Request.RequestURI

		// Status code
		statusCode := ctx.Writer.Status()

		// Request IP
		clientIp := ctx.ClientIP()

		// Body size
		bodySize := ctx.Writer.Size()

		logger.WithFields(logrus.Fields{
			"METHOD":    reqMethod,
			"URI":       reqUri,
			"STATUS":    statusCode,
			"LATENCY":   latencyTime,
			"CLIENT_IP": clientIp,
			"BODY_SIZE": bodySize,
		}).Info("HTTP REQUEST")
	}
}
