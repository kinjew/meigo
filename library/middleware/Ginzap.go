package middleware

import (
	"meigo/library/log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Ginzap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
func Ginzap() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middleware modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		latency := time.Since(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		switch {
		case statusCode >= 400 && statusCode <= 499:
			{
				log.Warn("[GIN]",
					/*
					   zap.Int("statusCode", statusCode),
					   zap.String("latency", latency.String()),
					   zap.String("clientIP", clientIP),
					   zap.String("method", method),
					   zap.String("path", path),
					   zap.String("query", query),
					   zap.String("user-agent", c.Request.UserAgent()),
					   zap.String("error", c.Errors.String()), */
					"statusCode:"+strconv.Itoa(statusCode),
					"latency:"+latency.String(),
					"clientIP:"+clientIP,
					"method:"+method,
					"path:"+path,
					"query:"+query,
					"user-agent:"+c.Request.UserAgent(),
					"error:"+c.Errors.String(),
				)
			}
		case statusCode >= 500:
			{
				log.Error("[GIN]",
					"statusCode:"+strconv.Itoa(statusCode),
					"latency:"+latency.String(),
					"clientIP:"+clientIP,
					"method:"+method,
					"path:"+path,
					"query:"+query,
					"user-agent:"+c.Request.UserAgent(),
					"error:"+c.Errors.String(),
				)
			}
		default:
			log.Info("[GIN]",
				"statusCode:"+strconv.Itoa(statusCode),
				"latency:"+latency.String(),
				"clientIP:"+clientIP,
				"method:"+method,
				"path:"+path,
				"query:"+query,
				"user-agent:"+c.Request.UserAgent(),
				"error:"+c.Errors.String(),
			)
		}
	}
}
