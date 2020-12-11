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

//Cors gin框架的http接口支持跨域请求
//http://www.5bug.wang/post/91.html
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("origin")
		if len(origin) == 0 {
			origin = c.Request.Header.Get("Origin")
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
