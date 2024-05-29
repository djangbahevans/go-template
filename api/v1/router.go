package v1

import (
	"net/http"

	"github.com/djangbahevans/go-template/api"
)

type V1Router struct{}

func (*V1Router) RegisterRoutes(r *http.ServeMux, routes ...api.IRoutes) {
	v1 := http.NewServeMux()
	for _, route := range routes {
		route.RegisterRoutes(v1)
	}

	r.Handle("/v1/", http.StripPrefix("/v1", v1))
}
