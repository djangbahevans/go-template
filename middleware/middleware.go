package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func ApplyMiddleware(h http.Handler, m ...Middleware) http.Handler {
	for _, middleware := range m {
		h = middleware(h)
	}
	return h
}

