package v1

import (
	"net/http"

	"github.com/djangbahevans/go-template/api/v1/handlers"
	"github.com/djangbahevans/go-template/services"
)

func RegisterV1Routes(r *http.ServeMux, userService services.UserService) {
	v1 := http.NewServeMux()
	handlers.RegisterUserRoutes(v1, userService)

	r.Handle("/v1/", http.StripPrefix("/v1", v1))
}
