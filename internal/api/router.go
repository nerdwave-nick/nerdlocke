package api

import (
	"net/http"

	// import the pprof package to expose the profiling routes
	_ "net/http/pprof"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/nerdwave-nick/nerdlocke/internal/api/common"
	"github.com/rs/cors"
)

type Controller interface {
	RegisterRoutes(common.RouteCreationContext)
}
type Persistence struct{}

func (p *Persistence) DoStuff() {
}

func MakeRouter(router *http.ServeMux, controllers []Controller) http.Handler {
	corsMiddleware := cors.New(cors.Options{
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "OPTIONS", "DELETE"},
		AllowCredentials: true,
	})

	humaConfig := huma.DefaultConfig("Event Bus", "v1")
	humaConfig.DocsPath = "/api/openapi"

	api := humago.New(router, humaConfig)
	rctx := common.RouteCreationContext{
		HTTP: router,
		API:  api,
	}

	// set up controllers
	// don't list them all here as before, since there will bee maaaaany of them
	// so for ease of reading, check the individual controllers for the routes
	// and the respective middleware setup
	for _, controller := range controllers {
		controller.RegisterRoutes(rctx)
	}

	handler := corsMiddleware.Handler(router)
	return handler
}
