package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "go.uber.org/zap"
    "vibegogo/pkg/logger"
    "net/http"
    "strings"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        bearerToken := strings.Split(authHeader, " ")
        if len(bearerToken) != 2 {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
            c.Abort()
            return
        }

        token := bearerToken[1]
        claims := jwt.MapClaims{}

        parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
            return []byte(secretKey), nil
        })

        if err != nil || !parsedToken.Valid {
            logger.Error("Invalid token", zap.Error(err))
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        userID := claims["user_id"].(string)
        c.Set("user_id", userID)
        c.Next()
    }
}