package api

import (
	"net/http"
)

type Api struct{}

func NewApi() *Api {
	return &Api{}
}

func (*Api) RegisterRoutes(r *http.ServeMux, routes ...IRoutes) {
	api := http.NewServeMux()
	for _, route := range routes {
		route.RegisterRoutes(r)
	}
	r.Handle("/api/", http.StripPrefix("/api", api))
}
