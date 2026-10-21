package controllers

import (
	"github.com/beaconsoftwarellc/quimby/v2/http"
)

// BasicAuthValidator handles verification of Username/Password
type BasicAuthValidator[T http.ContextBase] interface {
	Validate(context *http.Context[T], username, password string) bool
}

// BasicAuthenticatedController handles operations on the Registration Collection
type BasicAuthenticatedController[T http.ContextBase] struct {
	MethodNotAllowedController[T]
	Validator BasicAuthValidator[T]
}

// Authenticate verifies that a valid Registration header was provided
func (controller *BasicAuthenticatedController[T]) Authenticate(context *http.Context[T]) bool {
	context.Response.Header().Set("WWW-Authenticate", `Basic realm="Basic Auth Required"`)
	username, password, ok := context.Request.BasicAuth()
	if !ok {
		return false
	}
	return controller.Validator.Validate(context, username, password)
}

// HeaderAuthValidator handles verification by checking header values
type HeaderAuthValidator[T http.ContextBase] interface {
	Validate(context *http.Context[T]) bool
}

// HeaderAuthenticatedController handles operations on the Registration Collection
type HeaderAuthenticatedController[T http.ContextBase] struct {
	MethodNotAllowedController[T]
	validator HeaderAuthValidator[T]
}

// Authenticate verifies that a valid Registration header was provided
func (controller *HeaderAuthenticatedController[T]) Authenticate(context *http.Context[T]) bool {
	return controller.validator.Validate(context)
}
