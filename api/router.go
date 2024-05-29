package api

import "net/http"

type IRoutes interface {
	RegisterRoutes(r *http.ServeMux, routes ...IRoutes)
}
