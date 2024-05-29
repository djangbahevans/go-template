package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/djangbahevans/go-template/api"
	v1 "github.com/djangbahevans/go-template/api/v1"
	"github.com/djangbahevans/go-template/api/v1/handlers"
	"github.com/djangbahevans/go-template/config"
	"github.com/djangbahevans/go-template/databases/inmemory"
	"github.com/djangbahevans/go-template/middleware"
	"github.com/djangbahevans/go-template/repositories"
	"github.com/djangbahevans/go-template/services"
	"go.uber.org/fx"
)

func main() {
	config.LoadConfig()

	fx.New(
		fx.Provide(
			fx.Annotate(
				config.GetConfig,
				fx.ResultTags(`name:"config"`),
			),
			fx.Annotate(
				NewHttpServer,
				fx.ParamTags(`name:"config"`, `group:"http-server"`),
			),
			fx.Annotate(
				api.NewApi,
				fx.As(new(api.IRoute)),
				fx.ResultTags(`group:"http-server"`),
				fx.ParamTags(`group:"v1"`),
			),
			fx.Annotate(
				v1.NewV1Router,
				fx.As(new(api.IRoute)),
				fx.ResultTags(`group:"v1"`),
				fx.ParamTags(`group:"handlers"`),
			),
			fx.Annotate(
				handlers.NewUserRoutes,
				fx.As(new(api.IRoute)),
				fx.ResultTags(`group:"handlers"`),
			),
			fx.Annotate(
				services.NewUserService,
				fx.As(new(services.IUserService)),
			),
			fx.Annotate(
				inmemory.NewInMemoryDB,
				fx.As(new(repositories.IUserRepository)),
			),
		),
		fx.Invoke(func(lc fx.Lifecycle, srv *http.Server) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go func() {
						fmt.Printf("Starting server on %s\n", srv.Addr)
						if err := srv.ListenAndServe(); err != nil {
							fmt.Println(err)
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return srv.Shutdown(ctx)
				},
			})
		}),
	).Run()
}

func NewHttpServer(cfg *config.Config, routes ...api.IRoute) *http.Server {
	api := http.NewServeMux()

	for _, route := range routes {
		route.RegisterRoutes(api)
	}

	server := &http.Server{
		Addr:    cfg.ServerAddr,
		Handler: middleware.ApplyMiddleware(api, middleware.Logger),
	}

	return server
}
