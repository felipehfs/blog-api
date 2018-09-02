// Package controller contains the actions and middlewares
// for the api works well
package controller

import (
	"net/http"
)

// Cors setup the headers correctly
func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Add("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		next.ServeHTTP(w, r)
	})
}
