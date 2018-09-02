package controller

import "net/http"

type middleware func(http.Handler) http.Handler

// Adapt chain the handlers to one
func Adapt(handler http.Handler, middlewares ...middleware) http.Handler {
	for _, mid := range middlewares {
		handler = mid(handler)
	}
	return handler
}
