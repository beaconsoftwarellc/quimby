package security

import (
	qcontrollers "gitlab.com/beacon-software/quimby/controllers"
	qhttp "gitlab.com/beacon-software/quimby/http"
)

type basicValidator struct{}

func (validator *basicValidator) Validate(context *qhttp.Context, username, password string) bool {
	return username == password
}

// NewBasicValidator returns an initialized BasicAuthValidator
func NewBasicValidator() qcontrollers.BasicAuthValidator {
	return &basicValidator{}
}

type tokenValidator struct{}

func (validator *tokenValidator) Validate(context *qhttp.Context) bool {
	return context.Request.Header.Get("Authorization") == "valid"
}

// NewTokenValidator returns an initialized HeaderAuthValidator
func NewTokenValidator() qcontrollers.HeaderAuthValidator {
	return &tokenValidator{}
}
