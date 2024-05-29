package api

import "net/http"

type IRoute interface {
	RegisterRoutes(r *http.ServeMux)
}
