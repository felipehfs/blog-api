// Package controller contains the api actions
package controller

import (
	"context"
	"net/http"

	"github.com/felipehfs/blog/model"
)

// WithDB setup the database middleware
func WithDB(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		database, err := model.NewDatabase()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}
		defer database.Close()

		ctx := context.WithValue(r.Context(), BlogContext("database"), database)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
