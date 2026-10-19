package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/beaconsoftwarellc/gadget/v2/log"
	"github.com/beaconsoftwarellc/gadget/v2/stringutil"
	qerror "github.com/beaconsoftwarellc/quimby/v2/errors"
)

// HealthCheckRoute is the default URI for health checks
const HealthCheckRoute = "health"
const healthCheckURI = "/" + HealthCheckRoute

// RESTServer is a struct for managing the configuration and start up of a
// http/s server using the routing and controller logic in this package.
type RESTServer[T ContextBase] struct {
	Address    string
	Port       int
	Router     Router[T]
	newContext NewContextBase[T]
}

// NewContextBase for an incoming request
type NewContextBase[T ContextBase] func(context.Context) T

// CreateRESTServer initializes a RESTServer struct and returns it.
func CreateRESTServer[T ContextBase](address string, rootController Controller[T],
	newContext NewContextBase[T]) RESTServer[T] {
	server := RESTServer[T]{Address: address}
	server.Router = CreateRouter(rootController)
	server.newContext = newContext
	return server
}

// ServeHTTP processes the HTTP Request
func (server *RESTServer[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	ctx := CreateContext(server.newContext(r.Context()), w, r, server.Router, nil)
	if !ctx.HasError() {
		switch ctx.Request.Method {
		case http.MethodGet:
			ctx.Route.Controller.Get(ctx)
		case http.MethodPost:
			ctx.Route.Controller.Post(ctx)
		case http.MethodPut:
			ctx.Route.Controller.Put(ctx)
		case http.MethodPatch:
			ctx.Route.Controller.Patch(ctx)
		case http.MethodDelete:
			ctx.Route.Controller.Delete(ctx)
		case http.MethodOptions:
			ctx.Route.Controller.Options(ctx)
		default:
			ctx.SetError(qerror.NewRestError(qerror.MethodNotAllowed, "", nil),
				http.StatusMethodNotAllowed)
		}
	}
	server.CompleteRequest(start, ctx)
}

const (
	contentTypeHeader             = "Content-Type"
	contentTypeJSON               = "application/json"
	contentTypeForm               = "application/x-www-form-urlencoded"
	contentTypeMultiPartFormData  = "multipart/form-data"
	contentTypeMultiPartFormData1 = "multipart/mixed"
	locationHeader                = "Location"
)

// CompleteRequest generates output and completes the Request
func (server *RESTServer[T]) CompleteRequest(start time.Time, context *Context) {
	if healthCheckURI != context.URI {
		defer func(start time.Time) {
			log.Accessf("%s %s %s %s %d %dms %s %s %s",
				context.Request.RemoteAddr,
				context.Request.Method,
				context.Request.Proto,
				context.Request.URL.String(),
				context.Status(),
				time.Since(start).Milliseconds(),
				context.URLParameters.Encode(),
				context.Request.UserAgent(),
				context.Request.Referer(),
			)
		}(start)
	}

	contentType := context.Response.Header().Get(contentTypeHeader)
	if stringutil.IsWhiteSpace(contentType) {
		contentType = contentTypeJSON
		context.Response.Header().Set(contentTypeHeader, contentType)
	}

	if contentTypeJSON == contentType {
		server.completeRequestJSON(context)
		return
	}

	var b []byte
	if context.HasError() {
		b = []byte(context.Error.Message)
	} else if nil != context.Model {
		b = []byte(fmt.Sprintf("%v", context.Model))
	}

	context.Response.WriteHeader(context.responseStatus)
	_ = log.Warn(context.Write(b))
}

func (server *RESTServer[T]) completeRequestJSON(context *Context) {
	var (
		b []byte
		e error
	)
	if context.HasError() {
		b, e = json.Marshal(context.Error)
	} else if _, ok := context.Model.(string); ok {
		b = []byte(context.Model.(string))
	} else {
		b, e = json.Marshal(context.Model)
	}
	if e != nil {
		context.SetError(qerror.NewRestError("system-error", "", nil), http.StatusInternalServerError)
		b, _ = json.Marshal(context.Error)
	}
	context.Response.WriteHeader(context.responseStatus)
	_ = log.Warn(context.Write(b))
}

// ListenAndServe starts a http server listening on the address specified
// on the RESTServer instance.
func (server *RESTServer[T]) ListenAndServe() error {
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         server.Address,
		Handler:      server,
	}
	err := srv.ListenAndServe()
	return err
}

// ListenAndServeTLS starts a https server listening on the address specified
// on the RESTServer instance.
func (server *RESTServer[T]) ListenAndServeTLS(address string, port int) error {
	// qserver := &QuimbyServer{Address: addr}
	// err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	// log.Fatal(err)
	return nil
}
