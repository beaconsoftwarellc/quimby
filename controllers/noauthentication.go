package controllers

import qhttp "gitlab.com/beacon-software/quimby/http"

// NoAuthenticationController serves as a base for controllers that do not
// implement authentication.
type NoAuthenticationController struct{}

// Authenticate always returns true
func (controller NoAuthenticationController) Authenticate(context *qhttp.Context) bool {
	return true
}
