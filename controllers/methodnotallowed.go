package controllers

import (
	"net/http"

	http2 "github.com/beaconsoftwarellc/quimby/v2/http"

	qerror "github.com/beaconsoftwarellc/quimby/v2/errors"
)

// MethodNotAllowedController serves as a base for controllers that do not
// implement all the complete controller interface.
type MethodNotAllowedController[T http2.ContextBase] struct{}

// GetRoutes returns an emtpy string array.
func (controller MethodNotAllowedController[T]) GetRoutes() []string {
	return []string{}
}

// Get returns a method not allowed status
func (controller MethodNotAllowedController[T]) Get(context *http2.Context[T]) {
	context.SetError(qerror.NewRestError(qerror.MethodNotAllowed, "", nil), http.StatusMethodNotAllowed)
}

// Post returns a method not allowed status
func (controller MethodNotAllowedController[T]) Post(context *http2.Context[T]) {
	context.SetError(qerror.NewRestError(qerror.MethodNotAllowed, "", nil), http.StatusMethodNotAllowed)
}

// Put returns a method not allowed status
func (controller MethodNotAllowedController[T]) Put(context *http2.Context[T]) {
	context.SetError(qerror.NewRestError(qerror.MethodNotAllowed, "", nil), http.StatusMethodNotAllowed)
}

// Patch returns a method not allowed status
func (controller MethodNotAllowedController[T]) Patch(context *http2.Context[T]) {
	context.SetError(qerror.NewRestError(qerror.MethodNotAllowed, "", nil), http.StatusMethodNotAllowed)
}

// Delete returns a method not allowed status
func (controller MethodNotAllowedController[T]) Delete(context *http2.Context[T]) {
	context.SetError(qerror.NewRestError(qerror.MethodNotAllowed, "", nil), http.StatusMethodNotAllowed)
}

// Options returns a method not allowed status
func (controller MethodNotAllowedController[T]) Options(context *http2.Context[T]) {
	context.SetError(qerror.NewRestError(qerror.MethodNotAllowed, "", nil), http.StatusMethodNotAllowed)
}
