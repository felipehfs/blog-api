// Package controller contains the actions and some
// middlewares to api works well
package controller

import (
	"log"
	"net/http"
)

// Logging is the middleware for display all notifications
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[ %s ] %-40s - %10s\n", r.Method, r.URL.Path, r.UserAgent())

		next.ServeHTTP(w, r)
	})
}
