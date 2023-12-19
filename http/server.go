package http

import (
	"encoding/json"
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
type RESTServer struct {
	Address string
	Port    int
	Router  Router
}

// CreateRESTServer initializes a RESTServer struct and returns it.
func CreateRESTServer(address string, rootController Controller) RESTServer {
	server := RESTServer{Address: address}
	server.Router = CreateRouter(rootController)
	return server
}

// ServeHTTP processes the HTTP Request
func (server *RESTServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	context := CreateContext(w, r, server.Router)
	if !context.HasError() {
		switch context.Request.Method {
		case http.MethodGet:
			context.Route.Controller.Get(context)
		case http.MethodPost:
			context.Route.Controller.Post(context)
		case http.MethodPut:
			context.Route.Controller.Put(context)
		case http.MethodPatch:
			context.Route.Controller.Patch(context)
		case http.MethodDelete:
			context.Route.Controller.Delete(context)
		case http.MethodOptions:
			context.Route.Controller.Options(context)
		default:
			context.SetError(qerror.NewRestError(qerror.MethodNotAllowed, "", nil), http.StatusMethodNotAllowed)
		}
	}
	server.CompleteRequest(start, context)
}

const (
	contentTypeHeader             = "Content-Type"
	contentTypeJSON               = "application/json"
	contentTypeForm               = "application/x-www-form-urlencoded"
	contentTypeMultiPartFormData  = "multipart/form-data"
	contentTypeMultiPartFormData1 = "multipart/mixed"
)

// CompleteRequest generates output and completes the Request
func (server *RESTServer) CompleteRequest(start time.Time, context *Context) {
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
	if stringutil.IsWhiteSpace(context.Response.Header().Get(contentTypeHeader)) {
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
		b = []byte(context.Model.(string))
	}

	context.Response.WriteHeader(context.responseStatus)
	_ = log.Warn(context.Write(b))
}

func (server *RESTServer) completeRequestJSON(context *Context) {
	context.Response.Header().Add(contentTypeHeader, contentTypeJSON)
	var b []byte
	var e error
	if context.HasError() {
		b, e = json.Marshal(context.Error)
		if e != nil {
			context.SetError(qerror.NewRestError("system-error", "", nil), http.StatusInternalServerError)
			b, _ = json.Marshal(context.Error)
		}
	} else {
		b, e = json.Marshal(context.Model)
		if e != nil {
			context.SetError(qerror.NewRestError("system-error", "", nil), http.StatusInternalServerError)
			b, _ = json.Marshal(context.Error)
		}
	}
	context.Response.WriteHeader(context.responseStatus)
	err := context.Write(b)
	_ = log.Warn(err)
}

// ListenAndServe starts a http server listening on the address specified
// on the RESTServer instance.
func (server *RESTServer) ListenAndServe() error {
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
func (server *RESTServer) ListenAndServeTLS(address string, port int) error {
	// qserver := &QuimbyServer{Address: addr}
	// err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	// log.Fatal(err)
	return nil
}
