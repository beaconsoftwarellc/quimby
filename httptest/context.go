package httptest

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/beaconsoftwarellc/quimby/v2/controllers"
	qhttp "github.com/beaconsoftwarellc/quimby/v2/http"
)

/******************************************************
 *          Supporting code for tests                 *
 ******************************************************/

// CreateTestContext creates a Context that is appropriate for testing
func CreateTestContext(c qhttp.Controller[qhttp.ContextBase], r *http.Request) *qhttp.Context[qhttp.ContextBase] {
	w := httptest.NewRecorder()
	router := qhttp.CreateRouter[qhttp.ContextBase](c)
	return qhttp.CreateContext[qhttp.ContextBase](qhttp.NewDefaultBase(context.Background()), w, r, router, nil)
}

// TestController implements the Controller interface and nothing else
type TestController[T qhttp.ContextBase] struct {
	controllers.MethodNotAllowedController[T]
	controllers.NoAuthenticationController[T]
}
