package api

import (
	"net/http"
)

type Api struct {
	routes []IRoute
}

func NewApi(routes ...IRoute) *Api {
	return &Api{routes}
}

func (router *Api) RegisterRoutes(r *http.ServeMux) {
	mux := http.NewServeMux()
	for _, route := range router.routes {
		route.RegisterRoutes(mux)
	}

	r.Handle("/api/", http.StripPrefix("/api", mux))
}
