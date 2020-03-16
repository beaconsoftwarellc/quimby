package test

// THIS IS A GENERATED FILE. DO NOT MODIFY
// test_config.tmpl

import (
	"github.com/beaconsoftwarellc/gadget/log"
	"github.com/beaconsoftwarellc/quimby/example/config"
)

// NewMockSpec for use in unit tests.
func NewMockSpec() (*config.Specification) {
	spec := &config.Specification{
		Log: log.NewStackLogger(),
	}
	return spec
}
