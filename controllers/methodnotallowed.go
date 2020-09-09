package controllers

import (
	http2 "github.com/beaconsoftwarellc/quimby/http"
	"net/http"

	qerror "github.com/beaconsoftwarellc/quimby/error"
)

// MethodNotAllowedController serves as a base for controllers that do not
// implement all the complete controller interface.
type MethodNotAllowedController struct{}

// GetRoutes returns an emtpy string array.
func (controller MethodNotAllowedController) GetRoutes() []string {
	return []string{}
}

// Get returns a method not allowed status
func (controller MethodNotAllowedController) Get(context *http2.Context) {
	context.SetError(qerror.NewRestError(qerror.MethodNotAllowed, "", nil), http.StatusMethodNotAllowed)
}

// Post returns a method not allowed status
func (controller MethodNotAllowedController) Post(context *http2.Context) {
	context.SetError(qerror.NewRestError(qerror.MethodNotAllowed, "", nil), http.StatusMethodNotAllowed)
}

// Put returns a method not allowed status
func (controller MethodNotAllowedController) Put(context *http2.Context) {
	context.SetError(qerror.NewRestError(qerror.MethodNotAllowed, "", nil), http.StatusMethodNotAllowed)
}

// Patch returns a method not allowed status
func (controller MethodNotAllowedController) Patch(context *http2.Context) {
	context.SetError(qerror.NewRestError(qerror.MethodNotAllowed, "", nil), http.StatusMethodNotAllowed)
}

// Delete returns a method not allowed status
func (controller MethodNotAllowedController) Delete(context *http2.Context) {
	context.SetError(qerror.NewRestError(qerror.MethodNotAllowed, "", nil), http.StatusMethodNotAllowed)
}

// Options returns a method not allowed status
func (controller MethodNotAllowedController) Options(context *http2.Context) {
	context.SetError(qerror.NewRestError(qerror.MethodNotAllowed, "", nil), http.StatusMethodNotAllowed)
}
