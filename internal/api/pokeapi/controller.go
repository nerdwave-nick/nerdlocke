package intapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nerdwave-nick/nerdlocke/internal/api/common"
)

type PokeApiBody struct {
	Body   json.RawMessage
	Status int
}

type Controller struct{}

func (c *Controller) RegisterRoutes(rctx common.RouteCreationContext) {
	defaultTags := []string{"Pokeapi"}
	// basic health/liveness check routes
	common.AddHumaRoute(rctx, c.Healthz, huma.Operation{
		Method: http.MethodGet,
		Path:   "/api/pokeapi",
		Tags:   defaultTags,
	})
}

// Live is a health check handler for checking that the server is running and serving requests
// It's probed by the Cloud run service's own health check: https://console.cloud.google.com/run/detail/europe-west3/questionnaire-backend-v2
func (c *Controller) Healthz(_ context.Context, _ *struct{}) (*PokeApiBody, huma.StatusError) {
	return &PokeApiBody{Body: []byte(`{"heh":"Awoo"}`), Status: http.StatusOK}, nil
}

func MakeController() *Controller {
	return &Controller{}
}
