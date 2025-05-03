package middleware

import (
    "net/http"
    "time"
    "vibegogo/logger"
)

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        duration := time.Since(start)
        logger.LogRequest(r.URL.Path, duration, 200, map[string]interface{}{
            "method": r.Method,
            "remote": r.RemoteAddr,
        })
    })
}