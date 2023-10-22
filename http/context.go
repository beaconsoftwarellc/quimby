package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"strings"

	"github.com/beaconsoftwarellc/gadget/v2/errors"
	"github.com/beaconsoftwarellc/gadget/v2/log"
	"github.com/beaconsoftwarellc/gadget/v2/stringutil"
	qerror "github.com/beaconsoftwarellc/quimby/v2/error"
	"github.com/beaconsoftwarellc/quimby/v2/http/multipartform"
	"github.com/beaconsoftwarellc/quimby/v2/http/urlencodedform"
)

// NoContentError is returned when Read is called and the Request has a 0
// content length.
type NoContentError struct {
	RequestPath   string
	RequestMethod string
	trace         []string
}

// NewNoContentError instantiates a NoContentError with a stack trace
func NewNoContentError(path, method string) errors.TracerError {
	return &NoContentError{
		RequestPath:   path,
		RequestMethod: method,
		trace:         errors.GetStackTrace(),
	}
}

func (err *NoContentError) Error() string {
	return fmt.Sprintf("Request (%s %s) cannot be 'Read' as it has no content.",
		err.RequestMethod, err.RequestPath)
}

// Trace returns the stack trace for the error
func (err *NoContentError) Trace() []string {
	return err.trace
}

// Context serves as a structure that tracks the state of a given http Request
// Response chain.
type Context struct {
	context.Context
	URIParameters map[string]string
	URLParameters url.Values
	URI           string
	URL           *url.URL
	Method        string

	Request  *http.Request
	Response http.ResponseWriter
	Route    *RouteNode

	responseStatus int
	Authentication Authentication
	Model          interface{}
	Error          *qerror.RestError

	Extended map[string]interface{}

	Body     string
	bodyRead bool
}

// Status returns the HTTP status of the response
func (context *Context) Status() int {
	return context.responseStatus
}

// SetError sets the HTTP status of the response and the error to be returned
func (context *Context) SetError(err *qerror.RestError, status int) {
	context.responseStatus = status
	context.Error = err
}

// HasError checks if there is an Error set on the Context
func (context *Context) HasError() bool {
	return nil != context.Error
}

// AddError adds an error detail to the context
// Primary use case is to add a series of validation type errors
func (context *Context) AddError(err qerror.FieldError) {
	if !context.HasError() {
		context.SetError(qerror.NewRestError(qerror.ValidationError, "", nil), http.StatusBadRequest)
	}
	context.Error.AddDetail(err)
}

// SetResponse sets the HTTP status and model to be rendered in the response write
// Returns false if there is an Error on the context otherwise true
func (context *Context) SetResponse(model interface{}, status int) bool {
	if context.HasError() {
		return false
	}
	context.responseStatus = status
	context.Model = model
	return true
}

// CreateContext initializes a Context from the passed Response and Request
// pair, and router. The router is used for detemplating and populating the
// URIParameters
func CreateContext(writer http.ResponseWriter, request *http.Request,
	router Router) *Context {
	var err error
	qctx := &Context{Request: request, Extended: make(map[string]interface{})}
	qctx.Response = writer
	qctx.URL = request.URL
	qctx.URI = request.RequestURI
	qctx.Method = request.Method
	qctx.URLParameters, err = url.ParseQuery(request.URL.RawQuery)

	if ctx := request.Context(); ctx != nil {
		qctx.Context = ctx
	} else {
		qctx.Context = context.Background()
	}

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

// InvalidCredentialsErrorMessage is returned when Credentials are invalid
const InvalidCredentialsErrorMessage = "Invalid Credentials"

// Read reads the entire body of the request and returns it as a slice of
// bytes
func (context *Context) Read() ([]byte, error) {
	if context.bodyRead {
		return []byte(context.Body), nil
	}
	if context.Request.ContentLength <= 0 {
		return nil, NewNoContentError("", "")
	}
	body := make([]byte, context.Request.ContentLength)
	n, err := io.ReadFull(context.Request.Body, body)

	if err == io.ErrUnexpectedEOF {
		log.Errorf("warning:%s:%s: Request.ContentLength (%d) mismatch with actual body length (%d)", context.URI,
			context.Request.RemoteAddr, n, context.Request.ContentLength)
	}
	// Ignore EOF error
	if io.EOF == err {
		err = nil
	}
	context.Body = string(body)
	context.bodyRead = true
	return body, err
}

// readJSON takes the JSON content type body of the Request and unmarshals a JSON object the
// same type as the passed implementation of interface{}
func (context *Context) readJSON(body []byte, target interface{}) error {
	return json.Unmarshal(body, target)
}

// ReadObject reads the body of the Request and unmarshals an object the
// same type as the passed implementation of interface{}
func (context *Context) ReadObject(target interface{}) error {
	body, err := context.Read()

	if err != nil {
		return err
	}
	contentType, _, err := mime.
		ParseMediaType(context.Request.Header.Get(contentTypeHeader))
	if nil != err {
		context.SetError(qerror.NewRestError(qerror.ValidationError, err.Error(), nil), http.StatusNotAcceptable)
		return err
	}
	switch contentType {
	case contentTypeForm:
		err = urlencodedform.Unmarshal(body, target)
	case contentTypeMultiPartFormData:
		fallthrough
	case contentTypeMultiPartFormData1:
		err = context.Request.
			ParseMultipartForm(multipartform.MaxMemoryBytes)
		if nil == err {
			err = multipartform.Unmarshal(context.Request.MultipartForm, target)
		}
	case contentTypeJSON:
		err = context.readJSON(body, target)
	default:
		err = errors.New("Unsupported contentType (%s) provided", contentType)
	}
	if nil != err {
		context.SetError(qerror.NewRestError(qerror.ValidationError, err.Error(), nil), http.StatusNotAcceptable)
	}

	return err
}

// Write writes a string to the response body.
func (context *Context) Write(bytes []byte) error {
	var (
		lastWrite int
		err       error
	)
	for written := 0; written < len(bytes); written += lastWrite {
		toWrite := bytes[written:]
		lastWrite, err = context.Response.Write(toWrite)
		if nil != err {
			return err
		}
	}
	return nil
}

// ReadQueryParams converts URL Parameters into an Object
func (context *Context) ReadQueryParams(target interface{}) error {
	return urlencodedform.URLValuesToObject(context.URLParameters, target)
}
