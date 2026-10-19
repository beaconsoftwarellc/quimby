package http

// Controller is the main interface for the request handlers in the Router
type Controller[T ContextBase] interface {
	// GetRoutes that this controller serves
	GetRoutes() []string
	Get(context *Context[T])
	Post(context *Context[T])
	Put(context *Context[T])
	Patch(context *Context[T])
	Delete(context *Context[T])
	Options(context *Context[T])
	Authenticate(context *Context[T]) (Authentication, bool)
}
