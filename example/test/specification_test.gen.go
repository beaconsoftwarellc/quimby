package test

// THIS IS A GENERATED FILE. DO NOT MODIFY
// test_config.tmpl

import (
	"gitlab.com/beacon-software/gadget/log"
	"gitlab.com/beacon-software/quimby/example/config"
)

// NewMockSpec for use in unit tests.
func NewMockSpec() *config.Specification {
	spec := &config.Specification{
		Log: log.NewStackLogger(),
	}
	return spec
}
