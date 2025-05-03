package middleware

import (
    "time"
    "github.com/gin-gonic/gin"
    "vibegogo/pkg/logger"
)

func LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        duration := time.Since(start)
        status := c.Writer.Status()
        logger.RequestLog(
            c.FullPath(),
            duration,
            status,
            c.Request.Method,
            c.ClientIP(),
        )
    }
}