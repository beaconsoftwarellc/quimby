package error

import (
	"errors"
	"net/http"
	"testing"

	"github.com/beaconsoftwarellc/gadget/generator"
	"github.com/beaconsoftwarellc/gadget/stringutil"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAddDetails(t *testing.T) {
	assert := assert.New(t)
	err := NewRestError("", "", nil)
	err.AddDetail("foo")
	assert.Equal("foo", err.Details[0])
	ferr := FieldError{}
	err.AddDetail(ferr)
	assert.Equal(ferr, err.Details[1])
}

type MockRestErrorContainer struct {
	SetErrorRestError *RestError
	SetErrorStatus    int
}

func (mock *MockRestErrorContainer) SetError(e *RestError, s int) {
	mock.SetErrorRestError = e
	mock.SetErrorStatus = s
}

func TestTranslateError(t *testing.T) {
	var tests = []struct {
		name             string
		err              error
		expectedRestCode string
		expectedStatus   int
	}{
		{
			name:             "no error",
			err:              nil,
			expectedRestCode: "",
			expectedStatus:   0,
		},
		{
			name:             "codes NotFound",
			err:              status.Error(codes.NotFound, ""),
			expectedRestCode: "not-found",
			expectedStatus:   http.StatusNotFound,
		},
		{
			name:             "codes.Unauthenticated",
			err:              status.Error(codes.Unauthenticated, ""),
			expectedRestCode: AuthenticationFailed,
			expectedStatus:   http.StatusUnauthorized,
		},
		{
			name:             "codes.AlreadyExists",
			err:              status.Error(codes.AlreadyExists, ""),
			expectedRestCode: ValidationError,
			expectedStatus:   http.StatusBadRequest,
		},
		{
			name:             "codes.PermissionDenied",
			err:              status.Error(codes.PermissionDenied, ""),
			expectedRestCode: NotAuthorized,
			expectedStatus:   http.StatusForbidden,
		},
		{
			name:             "codes.FailedPrecondition",
			err:              status.Error(codes.FailedPrecondition, ""),
			expectedRestCode: ValidationError,
			expectedStatus:   http.StatusBadRequest,
		},
		{
			name:             "codes.InvalidArgument",
			err:              status.Error(codes.InvalidArgument, ""),
			expectedRestCode: ValidationError,
			expectedStatus:   http.StatusBadRequest,
		},
		{
			name:             "codes.Canceled",
			err:              status.Error(codes.Canceled, ""),
			expectedRestCode: SystemError,
			expectedStatus:   http.StatusInternalServerError,
		},
		{
			name:             "FieldError",
			err:              &FieldError{},
			expectedRestCode: ValidationError,
			expectedStatus:   http.StatusBadRequest,
		},
		{
			name:             "NotFoundError",
			err:              &NotFoundError{},
			expectedRestCode: NotFound,
			expectedStatus:   http.StatusNotFound,
		},
		{
			name:             "NotAuthenticatedError",
			err:              &NotAuthenticatedError{},
			expectedRestCode: AuthenticationFailed,
			expectedStatus:   http.StatusUnauthorized,
		},
		{
			name:             "any",
			err:              errors.New(generator.String(20)),
			expectedRestCode: SystemError,
			expectedStatus:   http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			container := &MockRestErrorContainer{}
			TranslateError(container, tt.err)
			assert.Equal(tt.expectedStatus, container.SetErrorStatus)
			if stringutil.IsWhiteSpace(tt.expectedRestCode) {
				assert.Nil(container.SetErrorRestError)
			} else {
				assert.Equal(tt.expectedRestCode, container.SetErrorRestError.Code)
			}
		})
	}
}
