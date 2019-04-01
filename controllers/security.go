package controllers

import qhttp "gitlab.com/beacon-software/quimby/http"

// BasicAuthValidator handles verification of Username/Password
type BasicAuthValidator interface {
	Validate(context *qhttp.Context, username, password string) bool
}

// BasicAuthenticatedController handles operations on the Registration Collection
type BasicAuthenticatedController struct {
	MethodNotAllowedController
	Validator BasicAuthValidator
}

// Authenticate verifies that a valid Registration header was provided
func (controller *BasicAuthenticatedController) Authenticate(context *qhttp.Context) bool {
	context.Response.Header().Set("WWW-Authenticate", `Basic realm="Basic Auth Required"`)
	username, password, ok := context.Request.BasicAuth()
	if !ok {
		return false
	}
	return controller.Validator.Validate(context, username, password)
}

// HeaderAuthValidator handles verification by checking header values
type HeaderAuthValidator interface {
	Validate(context *qhttp.Context) bool
}

// HeaderAuthenticatedController handles operations on the Registration Collection
type HeaderAuthenticatedController struct {
	MethodNotAllowedController
	validator HeaderAuthValidator
}

// Authenticate verifies that a valid Registration header was provided
func (controller *HeaderAuthenticatedController) Authenticate(context *qhttp.Context) bool {
	return controller.validator.Validate(context)
}
