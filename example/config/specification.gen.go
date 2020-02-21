package config

// THIS IS A GENERATED FILE. DO NOT MODIFY
// config.tmpl

import (
	"github.com/beaconsoftwarellc/gadget/environment"
	"github.com/beaconsoftwarellc/gadget/log"
)

// Specification details the expected values for the config
type Specification struct {
  Log log.Logger
  
  Port int `env:"PORT,optional"`
  
  Storage storage.WidgetStorage
}

// New returns a Specification based on the environment
func New() *Specification {
	return NewValues(environment.GetEnvMap())
}

// NewValues returns a Specification based on the env var map passed in
func NewValues(envVars map[string]string) *Specification {
	s := &Specification{ 
		Port: 8080,
	}
	err := environment.ProcessMap(s, envVars)
	if nil != err {
		panic(log.Error(err))
	}

	s.Log = log.New("ExampleGateway", log.FunctionFromEnv())

	
  	s.Storage = storage.NewWidgetStorage()

	return s
}
