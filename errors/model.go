package errors

import (
	"fmt"
	"net/http"

	"github.com/beaconsoftwarellc/gadget/v2/errors"
	"github.com/beaconsoftwarellc/gadget/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// CannotBeBlank indicates a field that was submitted blank, but is required
	CannotBeBlank = "cannot-be-blank"
	// Canceled indicates the operation was canceled (typically by the caller).
	Canceled = "canceled"
	// ValidationError indicates that a validation rule such as min / max value was violated
	ValidationError = "validation-error"
	// MethodNotAllowed indicates that the attempted VERB is not implemented for that endpoint
	MethodNotAllowed = "method-not-allowed"
	// MalformedURL indicates that the URL was not parsable as input
	MalformedURL = "malformed-url"
	// InvalidRoute indicates that the route does not exist
	InvalidRoute = "invalid-route"
	// AuthenticationFailed indicates that authentication did not complete successfully
	AuthenticationFailed = "authentication-failed"
	// NotAuthorized indicates that the currently authenticated user is not
	// permitted to perform an action
	NotAuthorized = "not-authorized"
	// SystemError indicates that a systemic issue has occurred with the request
	SystemError = "system-error"
	// NotFound indicates that the requested resource was not found
	NotFound = "not-found"
	// AlreadyExists indicates that the requested resource already exist
	AlreadyExists = "already-exists"
)

// RestError represents the standard error returned by the API Gateway
type RestError struct {
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Details []interface{} `json:"details"`
}

// NewRestError instantiates a RestError
func NewRestError(code string, message string, details []interface{}) *RestError {
	return &RestError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

func (restError *RestError) Error() string {
	return fmt.Sprintf("%s (%s): %#v", restError.Message, restError.Code,
		restError.Details)
}

// AddDetail adds a detail such as a FieldError to an Error response
func (restError *RestError) AddDetail(errorDetail interface{}) {
	restError.Details = append(restError.Details, errorDetail)
}

// FieldError represents a validation error related to a specific input field
type FieldError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field"`
}

// NewFieldError instantiates a FieldError
func NewFieldError(code, message, field string) *FieldError {
	return &FieldError{
		Code:    code,
		Message: message,
		Field:   field,
	}
}

func (err *FieldError) Error() string {
	return fmt.Sprintf("validation failed for field %s: %s", err.Field,
		err.Message)
}

// RestErrorContainer holds a RestError
type RestErrorContainer interface {
	SetError(*RestError, int)
}

// NotFoundError is returned when the requested resource isn't found
type NotFoundError struct {
	trace []string
}

// NewNotFoundError instantiates a NotFoundError with a stack trace
func NewNotFoundError() errors.TracerError {
	return &NotFoundError{
		trace: errors.GetStackTrace(),
	}
}

func (err *NotFoundError) Error() string {
	return "not-found"
}

// Trace returns the stack trace for the error
func (err *NotFoundError) Trace() []string {
	return err.trace
}

// NotAuthenticatedError is returned for a failed login attempt
type NotAuthenticatedError struct {
	trace []string
}

// NewNotAuthenticatedError instantiates a NotAuthenticatedError with a stack trace
func NewNotAuthenticatedError() errors.TracerError {
	return &NotAuthenticatedError{
		trace: errors.GetStackTrace(),
	}
}

func (err *NotAuthenticatedError) Error() string {
	return "login-failed"
}

// Trace returns the stack trace for the error
func (err *NotAuthenticatedError) Trace() []string {
	return err.trace
}

// TranslateError from an ErrorMessage to a RestError and set it on a ErrorContainer
func TranslateError(container RestErrorContainer, err error) {
	if nil == err {
		return
	}
	var httpStatus int
	var restError *RestError
	// handle GRPC first
	statusError, ok := status.FromError(err)
	if ok {
		switch statusError.Code() {
		case codes.NotFound:
			restError = &RestError{Code: NotFound, Message: statusError.Message()}
			httpStatus = http.StatusNotFound
		case codes.Unauthenticated:
			restError = &RestError{Code: AuthenticationFailed, Message: statusError.Message()}
			httpStatus = http.StatusUnauthorized
		case codes.PermissionDenied:
			restError = &RestError{Code: NotAuthorized, Message: statusError.Message()}
			httpStatus = http.StatusForbidden
		case codes.AlreadyExists:
			restError = &RestError{Code: AlreadyExists, Message: statusError.Message()}
			httpStatus = http.StatusConflict
		case codes.OutOfRange:
			fallthrough
		case codes.InvalidArgument:
			fallthrough
		case codes.FailedPrecondition:
			restError = &RestError{Code: ValidationError, Message: statusError.Message()}
			httpStatus = http.StatusBadRequest
		case codes.Canceled:
			restError = &RestError{Code: Canceled, Message: statusError.Message()}
			// it may have been the caller or someone internally that cancelled the
			// request. So either way we are covered with StatusInternalServerError.
			httpStatus = http.StatusInternalServerError
		case codes.DeadlineExceeded:
			restError = &RestError{Code: SystemError, Message: statusError.Message()}
			httpStatus = http.StatusInternalServerError
		default:
			log.Errorf("[QMY.ERR.162] unhandled system error: %s", err)
			restError = &RestError{Code: SystemError, Message: statusError.Message()}
			httpStatus = http.StatusInternalServerError
		}
	} else {
		switch err.(type) {
		case *FieldError:
			restError = &RestError{Code: ValidationError, Message: err.Error()}
			httpStatus = http.StatusBadRequest
		case *NotFoundError:
			restError = &RestError{Code: NotFound, Message: err.Error()}
			httpStatus = http.StatusNotFound
		case *NotAuthenticatedError:
			restError = &RestError{Code: AuthenticationFailed, Message: err.Error()}
			httpStatus = http.StatusUnauthorized
		default:
			log.Errorf("[QMY.ERR.178] unhandled system error: %s", err)
			restError = &RestError{Code: SystemError, Message: err.Error()}
			httpStatus = http.StatusInternalServerError
		}
	}
	container.SetError(restError, httpStatus)
}
