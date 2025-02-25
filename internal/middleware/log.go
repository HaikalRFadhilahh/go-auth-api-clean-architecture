package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Var Path and Method
		requestMethod := r.Method
		requestUrlPath := r.URL.Path
		requestTime := time.Now().Format("15:04 02-01-2006")

		// Print Some Log
		fmt.Printf("%v - %v (%v)\n", requestMethod, requestUrlPath, requestTime)

		// Next Request
		next.ServeHTTP(w, r)
	})
}
