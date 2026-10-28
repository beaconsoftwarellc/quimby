package http

import (
	"context"

	"github.com/beaconsoftwarellc/gadget/v2/generator"
	"github.com/beaconsoftwarellc/gadget/v2/log"
	"github.com/beaconsoftwarellc/gadget/v2/stringutil"
)

//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock$GOFILE

// ContextBase is an interface for base functionality that is common to all contexts
type ContextBase interface {
	context.Context
	// GetLogger for this context.
	GetLogger() log.Logger
	// SetLogger for this context as a sub-logger of the passed logger with the
	// session off this context.
	SetLogger(logger log.Logger) log.Logger
	// GetSession for this context identified in log messages
	GetSession() string
	// SetSession for this context's log messages
	SetSession(prefix generator.IDPrefix)
	// GetAuthentication for this context
	GetAuthentication() Authentication
	// SetAuthentication for this context
	SetAuthentication(auth Authentication)
}

// NewDefaultBase returns a default implementation of ContextBase
func NewDefaultBase(ctx context.Context) ContextBase {
	return &DefaultBase{Context: ctx}
}

type DefaultBase struct {
	context.Context
	Session        string
	logger         log.Logger
	authentication Authentication
}

func (c *DefaultBase) GetLogger() log.Logger {
	if c.logger == nil {
		c.logger = log.Global().New(c.Session)
	}
	return c.logger
}

func (c *DefaultBase) SetLogger(logger log.Logger) log.Logger {
	if logger == nil {
		logger = log.Global()
	}
	c.logger = logger.New(c.Session)
	return c.logger
}

func (c *DefaultBase) GetSession() string {
	return c.Session
}

func (c *DefaultBase) SetSession(prefix generator.IDPrefix) {
	if stringutil.IsWhiteSpace(c.GetSession()) {
		c.Session = generator.ID(prefix)
	}
}

func (c *DefaultBase) GetAuthentication() Authentication {
	return c.authentication
}

func (c *DefaultBase) SetAuthentication(auth Authentication) {
	c.authentication = auth
}
