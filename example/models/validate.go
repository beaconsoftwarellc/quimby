package models

import (
	"strings"

	"gitlab.com/beacon-software/gadget/errors"
	"gitlab.com/beacon-software/gadget/stringutil"
)

func issuesToError(issues []string) errors.TracerError {
	if 0 == len(issues) {
		return nil
	}
	return errors.New(strings.Join(issues, "\n"))
}

// Valid indicates if the request is complete
func (req *WidgetRequest) Valid() errors.TracerError {
	issues := []string{}
	if stringutil.IsWhiteSpace(req.SerialNumber) {
		issues = append(issues, "SerialNumber cannot be blank")
	}
	if stringutil.IsWhiteSpace(req.Description) {
		issues = append(issues, "Description cannot be blank")
	}
	return issuesToError(issues)
}

// Valid indicates if the request is complete
func (req *WidgetPatch) Valid() errors.TracerError {
	issues := []string{}
	if stringutil.IsWhiteSpace(req.Description) {
		issues = append(issues, "Description cannot be blank")
	}
	return issuesToError(issues)
}
