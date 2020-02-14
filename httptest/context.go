package httptest

import (
	"net/http"
	"net/http/httptest"

	"github.com/beaconsoftwarellc/quimby/controllers"
	qhttp "github.com/beaconsoftwarellc/quimby/http"
)

/******************************************************
 *          Supporting code for tests                 *
 ******************************************************/

// CreateTestContext creates a Context that is appropriate for testing
func CreateTestContext(c qhttp.Controller, r *http.Request) (context *qhttp.Context) {
	w := httptest.NewRecorder()
	router := qhttp.CreateRouter(c)
	context = qhttp.CreateContext(w, r, router)
	return context
}

// TestController implements the Controller interface and nothing else
type TestController struct {
	controllers.MethodNotAllowedController
	controllers.NoAuthenticationController
}
