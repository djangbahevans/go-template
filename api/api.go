package api

import (
	"log"
	"net/http"

	v1 "github.com/djangbahevans/go-template/api/v1"
	"github.com/djangbahevans/go-template/middleware"
	"github.com/djangbahevans/go-template/services"
	"github.com/djangbahevans/go-template/utils"
)

type ApiServer struct {
	addr        string
	userService services.UserService
}

func NewApiServer(addr string, userService services.UserService) *ApiServer {
	return &ApiServer{addr: addr, userService: userService}
}

func (s *ApiServer) Start() error {
	r := http.NewServeMux()
	v1.RegisterV1Routes(r, s.userService)
	
	api := http.NewServeMux()
	api.Handle("/api/", http.StripPrefix("/api", r))
	api.HandleFunc("/health", HealthCheck)

	server := &http.Server{
		Addr:    s.addr,
		Handler: middleware.ApplyMiddleware(api, middleware.LoggingMiddleware),
	}

	log.Printf("server listening on %s", s.addr)

	return server.ListenAndServe()
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.RespondJSON(w, map[string]string{"status": "ok"})
}
