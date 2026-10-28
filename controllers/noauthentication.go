package controllers

import (
	"github.com/beaconsoftwarellc/quimby/v2/http"
	"github.com/beaconsoftwarellc/quimby/v2/http/authentication"
)

// NoAuthenticationController serves as a base for controllers that do not
// implement authentication.
type NoAuthenticationController[T http.ContextBase] struct {
	authenticator http.Authenticator[T]
}

// Authenticate always returns true
func (controller NoAuthenticationController[T]) Authenticate(context *http.Context[T]) (http.Authentication, bool) {
	if nil == controller.authenticator {
		controller.authenticator = authentication.NewAcceptAll[T]()
	}
	return controller.authenticator.Authenticate(context)
}
