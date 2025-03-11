package middleware

import (
	"net/http"
	"time"

	"github.com/shohratd15/todolist-api/internal/logger"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logger.Log.Infof("Received request: %s %s", r.Method, r.URL.Path)

		start := time.Now()
		next.ServeHTTP(w,r)
		duration := time.Since(start)

		logger.Log.Infof("Request processed in %s", duration)
	})
}