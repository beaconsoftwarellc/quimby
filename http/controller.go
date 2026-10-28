package http

// Controller is the main interface for the request handlers in the Router
type Controller[T ContextBase] interface {
	// GetRoutes that this controller serves
	GetRoutes() []string
	// Get request handler
	Get(*Context[T])
	// Post request handler
	Post(*Context[T])
	// Put request handler
	Put(*Context[T])
	// Patch request handler
	Patch(*Context[T])
	// Delete request handler
	Delete(*Context[T])
	// Options request handler
	Options(*Context[T])
	// Authenticate the request
	Authenticate(*Context[T]) (Authentication, bool)
}
