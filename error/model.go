package error

import (
	"fmt"
	"net/http"

	"github.com/beaconsoftwarellc/gadget/database"
	"github.com/beaconsoftwarellc/gadget/errors"
)

const (
	// CannotBeBlank indicates a field that was submitted blank, but is required
	CannotBeBlank = "cannot-be-blank"
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
	// NotAuthorized indicates that the currently authenticated user is not permitted to perform an action
	NotAuthorized = "not-authorized"
	// SystemError indicates that a systemic issue has occurred with the request
	SystemError = "system-error"
	// NotFound indicates that the requested resource was not found
	NotFound = "not-found"
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
	return fmt.Sprintf("%s (%s): %#v", restError.Message, restError.Code, restError.Details)
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
	var status int
	var restError *RestError

	switch err.(type) {
	case *database.NotFoundError:
		restError = &RestError{Code: NotFound, Message: err.Error()}
		status = http.StatusNotFound
	case *NotFoundError:
		restError = &RestError{Code: NotFound, Message: err.Error()}
		status = http.StatusNotFound
	case *NotAuthenticatedError:
		restError = &RestError{Code: NotFound, Message: err.Error()}
		status = http.StatusUnauthorized
	default:
		restError = &RestError{Code: ValidationError, Message: err.Error()}
		status = http.StatusBadRequest
	}
	container.SetError(restError, status)
}
