package health

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nerdwave-nick/nerdlocke/internal/api/common"
)

type HealthzBody struct {
	Message string `json:"message"`
}

type Controller struct{}

func (c *Controller) RegisterRoutes(rctx common.RouteCreationContext) {
	defaultTags := []string{"Health"}
	// basic health/liveness check routes
	common.AddHumaRoute(rctx, c.Healthz, huma.Operation{
		Method: http.MethodGet,
		Path:   "/healthz",
		Tags:   defaultTags,
	})
}

// Live is a health check handler for checking that the server is running and serving requests
// It's probed by the Cloud run service's own health check: https://console.cloud.google.com/run/detail/europe-west3/questionnaire-backend-v2
func (c *Controller) Healthz(_ context.Context, _ *struct{}) (*common.Response[HealthzBody], huma.StatusError) {
	return &common.Response[HealthzBody]{Body: &HealthzBody{Message: "OK"}, Status: http.StatusOK}, nil
}

func MakeController() *Controller {
	return &Controller{}
}
