package v1

import (
	"net/http"

	"github.com/djangbahevans/go-template/api"
)

type V1Router struct {
	routes []api.IRoute
}

func NewV1Router(routes ...api.IRoute) *V1Router {
	return &V1Router{routes: routes}
}

func (router *V1Router) RegisterRoutes(r *http.ServeMux) {
	v1 := http.NewServeMux()
	for _, route := range router.routes {
		route.RegisterRoutes(v1)
	}

	r.Handle("/v1/", http.StripPrefix("/v1", v1))
}
