package controllers

import (
	nhttp "net/http"
	"time"

	"github.com/beaconsoftwarellc/quimby/v2/http"
)

// HealthCheckController responds to Get requests with a JSON object representing
// the status of the server.
type HealthCheckController[T http.ContextBase] struct {
	MethodNotAllowedController[T]
	NoAuthenticationController[T]
}

// HealthCheckResource is returned by the HealthCheckController for conveying
// the status of the server.
type HealthCheckResource struct {
	Timestamp string
	Status    string
}

// GetRoutes only responds on the 'health' resource endpoint.
func (controller *HealthCheckController[T]) GetRoutes() []string {
	return []string{http.HealthCheckRoute}
}

// Get returns a new instance of the HealthCheckResource
func (controller *HealthCheckController[T]) Get(context *http.Context[T]) {
	model := HealthCheckResource{Timestamp: time.Now().UTC().Format(time.RFC822),
		Status: "OK"}
	context.SetResponse(model, nhttp.StatusOK)
}
