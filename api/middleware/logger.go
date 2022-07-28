package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func Logger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()

		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		requestLogger := logger.WithFields(logrus.Fields{
			"requestID": requestID,
			"path":      path,
		})

		c.Set("logger", requestLogger)

		start := time.Now()

		c.Next()

		requestLogger = requestLogger.WithFields(logrus.Fields{
			"latency":       time.Since(start).String(),
			"clientIP":      c.ClientIP(),
			"method":        c.Request.Method,
			"statusCode":    c.Writer.Status(),
			"internalError": c.Errors.ByType(gin.ErrorTypePrivate).String(),
			"responseSize":  c.Writer.Size(),
		})

		if c.Writer.Status() != 200 && c.Request.Body != nil {
			data, err := c.GetRawData()
			if err != nil {
				requestLogger = requestLogger.WithField("getRawDataError", err.Error())
			} else {
				requestLogger = requestLogger.WithField("body", string(data))
			}
		}

		requestLogger.Debug("Request info")
	}
}
