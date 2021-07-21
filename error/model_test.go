package error

import (
	"net/http"
	"testing"

	"github.com/beaconsoftwarellc/gadget/database"
	"github.com/stretchr/testify/assert"
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

type TestRestErrorContainer struct {
	Error  *RestError
	Status int
}

func (context *TestRestErrorContainer) SetError(err *RestError, status int) {
	context.Error = err
	context.Status = status
}

func TestTranslateError(t *testing.T) {
	assert := assert.New(t)
	expected := NewRestError(ValidationError, "", nil)
	rec := &TestRestErrorContainer{}
	TranslateError(rec, expected)
	assert.Equal(expected.Code, rec.Error.Code)
	assert.Equal(expected.Error(), rec.Error.Message)
	assert.Equal(http.StatusBadRequest, rec.Status)
}

func TestTranslateError_NotAuthenticatedError(t *testing.T) {
	assert := assert.New(t)
	expected := NewNotAuthenticatedError()
	rec := &TestRestErrorContainer{}
	TranslateError(rec, expected)
	assert.Equal(NotFound, rec.Error.Code)
	assert.Equal(expected.Error(), rec.Error.Message)
	assert.Equal(http.StatusUnauthorized, rec.Status)
}

func TestTranslateError_NotFoundError(t *testing.T) {
	assert := assert.New(t)
	expected := NewNotFoundError()
	rec := &TestRestErrorContainer{}
	TranslateError(rec, expected)
	assert.Equal(NotFound, rec.Error.Code)
	assert.Equal(expected.Error(), rec.Error.Message)
	assert.Equal(http.StatusNotFound, rec.Status)
}

func TestTranslateError_DatabaseNotFoundError(t *testing.T) {
	assert := assert.New(t)
	expected := database.NewNotFoundError()
	rec := &TestRestErrorContainer{}
	TranslateError(rec, expected)
	assert.Equal(NotFound, rec.Error.Code)
	assert.Equal(expected.Error(), rec.Error.Message)
	assert.Equal(http.StatusNotFound, rec.Status)
}
