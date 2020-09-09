package controllers

import (
    "github.com/beaconsoftwarellc/quimby/http"
    "github.com/beaconsoftwarellc/quimby/http/authentication"
)

// NoAuthenticationController serves as a base for controllers that do not
// implement authentication.
type NoAuthenticationController struct{
    authenticator http.Authenticator
}

// Authenticate always returns true
func (controller NoAuthenticationController) Authenticate(context *http.Context) (http.Authentication, bool) {
    if nil == controller.authenticator {
        controller.authenticator = authentication.NewAcceptAll()
    }
	return controller.authenticator.Authenticate(context)
}
