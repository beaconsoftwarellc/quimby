package http

import (
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/beaconsoftwarellc/gadget/v2/errors"
	"github.com/beaconsoftwarellc/gadget/v2/generator"
	"github.com/beaconsoftwarellc/gadget/v2/log"
	"github.com/beaconsoftwarellc/gadget/v2/stringutil"
	qerror "github.com/beaconsoftwarellc/quimby/v2/errors"
	"github.com/beaconsoftwarellc/quimby/v2/http/multipartform"
	"github.com/beaconsoftwarellc/quimby/v2/http/urlencodedform"
)

type Context[T ContextBase] struct {
	base          T
	URIParameters map[string]string
	URLParameters url.Values
	URI           string
	URL           *url.URL
	Method        string

	Request  *http.Request
	Response http.ResponseWriter
	Route    *RouteNode[T]

	responseStatus int
	Authentication Authentication
	Model          interface{}
	Error          *qerror.RestError

	Body     string
	bodyRead bool

	Extended map[string]interface{}
}

func (ctx *Context[T]) GetBase() T {
	return ctx.base
}

// GetMethod from the request on this context
func (ctx *Context[T]) GetMethod() string {
	return ctx.Method
}

func (ctx *Context[T]) GetRoute() *RouteNode[T] {
	return ctx.Route
}

// GetUrl from the request on this context
func (ctx *Context[T]) GetUrl() *url.URL {
	return ctx.URL
}

func (ctx *Context[T]) GetRequest() *http.Request {
	return ctx.Request
}

func (ctx *Context[T]) GetResponse() http.ResponseWriter {
	return ctx.Response
}

// Status returns the HTTP status of the response
func (ctx *Context[T]) Status() int {
	return ctx.responseStatus
}

// SetError sets the HTTP status of the response and the error to be returned
func (ctx *Context[T]) SetError(err *qerror.RestError, status int) {
	ctx.responseStatus = status
	ctx.Error = err
	if http.StatusInternalServerError == status && err.Code != qerror.Canceled {
		ctx.GetBase().GetLogger().Errorf("[HTTP.CTX.100] InternalServerError encountered "+
			"servicing request %s %s: %s", ctx.Method, ctx.URL, err)
	}
}

// HasError checks if there is an Error set on the Context
func (ctx *Context[T]) HasError() bool {
	return nil != ctx.Error
}

// SetResponse and the HTTP status and result model, returns a boolean indicating
// whether this [Context] is in an error state.
func (ctx *Context[T]) SetResponse(model interface{}, status int) bool {
	if ctx.HasError() {
		return false
	}
	if ctx.responseStatus != 0 {
		return false
	}
	ctx.responseStatus = status
	ctx.Model = model
	return true
}

func (ctx *Context[T]) SetRedirect(model interface{}, location string) bool {
	if ctx.HasError() {
		return false
	}
	if ctx.responseStatus != 0 {
		return false
	}
	ctx.responseStatus = http.StatusFound
	ctx.Response.Header().Set(locationHeader, location)
	ctx.Model = model
	return true
}

// CreateContext initializes a Context from the passed Response and Request
// pair, and router. The router is used for de-templating and populating the
// URIParameters
func CreateContext[T ContextBase](base T, writer http.ResponseWriter, request *http.Request,
	router Router[T], logger log.Logger) *Context[T] {
	var err error
	qctx := &Context[T]{
		base:     base,
		Request:  request,
		Extended: make(map[string]interface{}),
	}
	qctx.Response = writer
	qctx.URL = request.URL
	qctx.URI = request.RequestURI
	qctx.Method = request.Method
	qctx.URLParameters, err = url.ParseQuery(request.URL.RawQuery)

	// set the session on the base in case it has not been set yet.
	qctx.base.SetSession("http")
	// then set the logger
	logger = qctx.base.SetLogger(logger)

	if err != nil {
		// take a hard stance on malformed URL's
		qctx.SetError(qerror.NewRestError(qerror.MalformedURL,
			fmt.Sprintf("Malformed URL Parameters '%s'.", request.URL), nil),
			http.StatusBadRequest)
		return qctx
	}

	cleanPath := strings.Split(qctx.URI, "?")[0]
	cleanPath = strings.Trim(cleanPath, " /")
	qctx.Route, err = router.FindRouteForPath(cleanPath)
	if err != nil || qctx.Route == nil {
		qctx.SetError(qerror.NewRestError(qerror.InvalidRoute, "", nil), http.StatusBadRequest)
		return qctx
	}

	qctx.URIParameters = make(map[string]string)
	if !stringutil.IsWhiteSpace(cleanPath) {
		qctx.URIParameters, err = stringutil.Detemplate(qctx.Route.TemplateRoute, cleanPath)
		if err != nil {
			qctx.SetError(qerror.NewRestError(qerror.InvalidRoute, "", nil), http.StatusInternalServerError)
			return qctx
		}
	}
	var ok bool
	if http.MethodOptions != qctx.Request.Method {
		if qctx.Authentication, ok = qctx.Route.Controller.Authenticate(qctx); !ok {
			qctx.SetError(
				qerror.NewRestError(qerror.AuthenticationFailed,
					InvalidCredentialsErrorMessage, nil), http.StatusUnauthorized)
		}
	}

	return qctx
}

// Read reads the entire body of the request and returns it as a slice of
// bytes
func (ctx *Context[T]) Read() ([]byte, error) {
	if ctx.bodyRead {
		return []byte(ctx.Body), nil
	}
	body := make([]byte, ctx.Request.ContentLength)
	n, err := io.ReadFull(ctx.Request.Body, body)

	if errors.Is(err, io.ErrUnexpectedEOF) && ctx.base.Err() != nil {
		ctx.base.GetLogger().Warnf("%s:%s: Request.ContentLength (%d) "+
			"mismatch with actual body length (%d)", ctx.URI,
			ctx.Request.RemoteAddr, n, ctx.Request.ContentLength)
		err = ctx.base.Err()
	}
	if err == io.EOF {
		err = nil
	}

	ctx.Body = string(body)
	ctx.bodyRead = true
	return body, err
}

// readJSON takes the JSON content type body of the Request and unmarshals a JSON object the
// same type as the passed implementation of interface{}
func (ctx *Context[T]) readJSON(body []byte, target interface{}) error {
	return json.Unmarshal(body, target)
}

// ReadObject reads the body of the Request and unmarshals an object the
// same type as the passed implementation of interface{}
func (ctx *Context[T]) ReadObject(target interface{}) error {
	var (
		body        []byte
		err         error
		contentType string
	)
	contentType, _, err = mime.
		ParseMediaType(ctx.Request.Header.Get(contentTypeHeader))
	if nil != err {
		ctx.SetError(qerror.NewRestError(qerror.ValidationError,
			err.Error(), nil), http.StatusNotAcceptable)
		return err
	}
	switch contentType {
	case contentTypeForm:
		body, err = ctx.Read()
		if nil == err {
			err = urlencodedform.Unmarshal(body, target)
		}
	case contentTypeMultiPartFormData:
		fallthrough
	case contentTypeMultiPartFormData1:
		err = ctx.Request.
			ParseMultipartForm(multipartform.MaxMemoryBytes)
		if nil == err {
			err = multipartform.Unmarshal(ctx.Request.MultipartForm, target)
		}
	case contentTypeJSON:
		body, err = ctx.Read()
		if nil == err {
			err = ctx.readJSON(body, target)
		}
	default:
		err = errors.Newf("Unsupported contentType (%s) provided", contentType)
	}
	if nil != err {
		ctx.SetError(qerror.NewRestError(qerror.ValidationError,
			err.Error(), nil), http.StatusNotAcceptable)
	}
	return err
}

// Write writes a string to the response body.
func (ctx *Context[T]) Write(bytes []byte) error {
	var (
		lastWrite int
		err       error
	)
	for written := 0; written < len(bytes); written += lastWrite {
		toWrite := bytes[written:]
		lastWrite, err = ctx.Response.Write(toWrite)
		if nil != err {
			return err
		}
	}
	return nil
}

// ReadQueryParams converts URL Parameters into an Object
func (ctx *Context[T]) ReadQueryParams(target interface{}) error {
	return urlencodedform.URLValuesToObject(ctx.URLParameters, target)
}

// Context Implementation

func (ctx *Context[T]) Deadline() (deadline time.Time, ok bool) {
	return ctx.base.Deadline()
}
func (ctx *Context[T]) Done() <-chan struct{} {
	return ctx.base.Done()
}
func (ctx *Context[T]) Err() error {
	return ctx.base.Err()
}
func (ctx *Context[T]) Value(key any) any {
	return ctx.base.Value(key)
}

// ContextBase Implementation

func (ctx *Context[T]) GetLogger() log.Logger {
	return ctx.base.GetLogger()
}

func (ctx *Context[T]) SetLogger(logger log.Logger) log.Logger {
	return ctx.base.SetLogger(logger)
}

func (ctx *Context[T]) GetSession() string {
	return ctx.base.GetSession()
}

func (ctx *Context[T]) SetSession(prefix generator.IDPrefix) {
	ctx.base.SetSession(prefix)
}

func (ctx *Context[T]) GetAuthentication() Authentication {
	return ctx.GetAuthentication()
}

func (ctx *Context[T]) SetAuthentication(auth Authentication) {
	ctx.SetAuthentication(auth)
}
