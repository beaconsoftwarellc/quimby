package http

// Controller is the main interface for the request handlers in the Router
type Controller interface {
	// GetRoutes that this controller serves
	GetRoutes() []string
	Get(context *Context)
	Post(context *Context)
	Put(context *Context)
	Patch(context *Context)
	Delete(context *Context)
	Options(context *Context)
	Authenticate(context *Context) (Authentication, bool)
}
