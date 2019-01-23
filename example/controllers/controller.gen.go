package controllers

// THIS IS A GENERATED FILE. DO NOT MODIFY
// controller.tmpl

import (
	"fmt"
	"net/http"

	"gitlab.com/beacon-software/gadget/errors"
	qcontrollers "gitlab.com/beacon-software/quimby/controllers"
	qerror "gitlab.com/beacon-software/quimby/error"
	qhttp "gitlab.com/beacon-software/quimby/http"

	"gitlab.com/beacon-software/quimby/example/config"
	"gitlab.com/beacon-software/quimby/example/models"
	"gitlab.com/beacon-software/quimby/example/security"
)

// EchoController is a debugging tool for echo'ing back the request sent in as the body of the response
type EchoController interface {
	qhttp.Controller
}

type echoController struct {
	qcontrollers.MethodNotAllowedController
	qcontrollers.NoAuthenticationController
	Specification *config.Specification
}

// NewEchoController returns an initialized EchoController
func NewEchoController(spec *config.Specification) EchoController {
	controller := &echoController{}
	controller.Specification = spec

	return controller
}

// GetRoutes establishes routes for the EchoController
func (controller *echoController) GetRoutes() []string {
	return []string{
		"echo",
		"echo/{{toEcho}}",
	}
}

// ResourceController is a is a sample controller implementation
type ResourceController interface {
	qhttp.Controller
}

type resourceController struct {
	qcontrollers.MethodNotAllowedController
	qcontrollers.NoAuthenticationController
	Specification *config.Specification
}

// NewResourceController returns an initialized ResourceController
func NewResourceController(spec *config.Specification) ResourceController {
	controller := &resourceController{}
	controller.Specification = spec

	return controller
}

// GetRoutes establishes routes for the ResourceController
func (controller *resourceController) GetRoutes() []string {
	return []string{
		"resource/{{id}}",
		"resource/{{id}}/{{subresource}}",
	}
}

// WidgetsController handles List and Create functions for the Widget Collection
type WidgetsController interface {
	qhttp.Controller
	doGet(context *qhttp.Context, username string, password string) (*models.WidgetCollection, errors.TracerError)
	doPost(context *qhttp.Context, username string, password string, request *models.WidgetRequest) (*models.Widget, errors.TracerError)
}

type widgetsController struct {
	qcontrollers.BasicAuthenticatedController
	Specification *config.Specification
}

// NewWidgetsController returns an initialized WidgetsController
func NewWidgetsController(spec *config.Specification) WidgetsController {
	controller := &widgetsController{}
	controller.Specification = spec
	controller.Validator = security.NewBasicValidator()
	return controller
}

// GetRoutes establishes routes for the WidgetsController
func (controller *widgetsController) GetRoutes() []string {
	return []string{
		"api/widgets",
	}
}

// Get returns a collection Widgets
func (controller *widgetsController) Get(context *qhttp.Context) {
	username, password, ok := context.Request.BasicAuth()
	if !ok {
		context.SetError(qerror.NewRestError(qerror.AuthenticationFailed, "", nil), http.StatusUnauthorized)
	}
	resp, err := controller.doGet(context, username, password)
	if nil != err {
		qerror.TranslateError(context, err)
		controller.Specification.Log.Infof("%#v", map[string]string{"context": fmt.Sprintf("%#v", context), "error": fmt.Sprintf("%#v", err)})
		return
	}
	context.SetResponse(resp, http.StatusOK)
}

// Post creates a new Widget in the collection
func (controller *widgetsController) Post(context *qhttp.Context) {
	username, password, ok := context.Request.BasicAuth()
	if !ok {
		context.SetError(qerror.NewRestError(qerror.AuthenticationFailed, "", nil), http.StatusUnauthorized)
	}
	request := &models.WidgetRequest{}
	if err := context.ReadObject(request); nil != err {
		context.SetError(&qerror.RestError{Code: qerror.ValidationError, Message: err.Error()}, http.StatusNotAcceptable)
		return
	}
	resp, err := controller.doPost(context, username, password, request)
	if nil != err {
		qerror.TranslateError(context, err)
		controller.Specification.Log.Infof("%#v", map[string]string{"context": fmt.Sprintf("%#v", context), "payload": fmt.Sprintf("%#v", request), "error": fmt.Sprintf("%#v", err)})
		return
	}
	context.SetResponse(resp, http.StatusCreated)
}

// WidgetController handles Read, Update, Replace and Delete for a single Widget
type WidgetController interface {
	qhttp.Controller
	doGet(context *qhttp.Context, username string, password string, widgetID string) (*models.Widget, errors.TracerError)
	doPut(context *qhttp.Context, username string, password string, widgetID string, request *models.WidgetRequest) (*models.Widget, errors.TracerError)
	doPatch(context *qhttp.Context, username string, password string, widgetID string, request *models.WidgetPatch) (*models.Widget, errors.TracerError)
	doDelete(context *qhttp.Context, username string, password string, widgetID string) (*models.Widget, errors.TracerError)
}

type widgetController struct {
	qcontrollers.BasicAuthenticatedController
	Specification *config.Specification
}

// NewWidgetController returns an initialized WidgetController
func NewWidgetController(spec *config.Specification) WidgetController {
	controller := &widgetController{}
	controller.Specification = spec
	controller.Validator = security.NewBasicValidator()
	return controller
}

// GetRoutes establishes routes for the WidgetController
func (controller *widgetController) GetRoutes() []string {
	return []string{
		"api/widgets/{{id}}",
	}
}

// Get returns a the requested Widget
func (controller *widgetController) Get(context *qhttp.Context) {
	username, password, ok := context.Request.BasicAuth()
	if !ok {
		context.SetError(qerror.NewRestError(qerror.AuthenticationFailed, "", nil), http.StatusUnauthorized)
	}
	resp, err := controller.doGet(context, username, password, context.URIParameters["id"])
	if nil != err {
		qerror.TranslateError(context, err)
		controller.Specification.Log.Infof("%#v", map[string]string{"context": fmt.Sprintf("%#v", context), "error": fmt.Sprintf("%#v", err)})
		return
	}
	context.SetResponse(resp, http.StatusOK)
}

// Put replaces the identified Widget with the new Widget
func (controller *widgetController) Put(context *qhttp.Context) {
	username, password, ok := context.Request.BasicAuth()
	if !ok {
		context.SetError(qerror.NewRestError(qerror.AuthenticationFailed, "", nil), http.StatusUnauthorized)
	}
	request := &models.WidgetRequest{}
	if err := context.ReadObject(request); nil != err {
		context.SetError(&qerror.RestError{Code: qerror.ValidationError, Message: err.Error()}, http.StatusNotAcceptable)
		return
	}
	resp, err := controller.doPut(context, username, password, context.URIParameters["id"], request)
	if nil != err {
		qerror.TranslateError(context, err)
		controller.Specification.Log.Infof("%#v", map[string]string{"context": fmt.Sprintf("%#v", context), "payload": fmt.Sprintf("%#v", request), "error": fmt.Sprintf("%#v", err)})
		return
	}
	context.SetResponse(resp, http.StatusOK)
}

// Patch updates fields on the identified Widget
func (controller *widgetController) Patch(context *qhttp.Context) {
	username, password, ok := context.Request.BasicAuth()
	if !ok {
		context.SetError(qerror.NewRestError(qerror.AuthenticationFailed, "", nil), http.StatusUnauthorized)
	}
	request := &models.WidgetPatch{}
	if err := context.ReadObject(request); nil != err {
		context.SetError(&qerror.RestError{Code: qerror.ValidationError, Message: err.Error()}, http.StatusNotAcceptable)
		return
	}
	resp, err := controller.doPatch(context, username, password, context.URIParameters["id"], request)
	if nil != err {
		qerror.TranslateError(context, err)
		controller.Specification.Log.Infof("%#v", map[string]string{"context": fmt.Sprintf("%#v", context), "payload": fmt.Sprintf("%#v", request), "error": fmt.Sprintf("%#v", err)})
		return
	}
	context.SetResponse(resp, http.StatusOK)
}

// Delete deletes the identified Widget from the collection
func (controller *widgetController) Delete(context *qhttp.Context) {
	username, password, ok := context.Request.BasicAuth()
	if !ok {
		context.SetError(qerror.NewRestError(qerror.AuthenticationFailed, "", nil), http.StatusUnauthorized)
	}
	resp, err := controller.doDelete(context, username, password, context.URIParameters["id"])
	if nil != err {
		qerror.TranslateError(context, err)
		controller.Specification.Log.Infof("%#v", map[string]string{"context": fmt.Sprintf("%#v", context), "error": fmt.Sprintf("%#v", err)})
		return
	}
	context.SetResponse(resp, http.StatusNoContent)
}
