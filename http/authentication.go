package http

// Authentication provides a base for authentication implementations that are used as part of the http
// context chain.
type Authentication interface {
	// GetKind of authentication this instance provides
	GetKind() string
	// GetUserID that is requesting authentication
	GetUserID() string
	// GetValidity indicates whether the user id on this instance is authentic
	GetValidity() bool
}

// Authenticator is responsible for the initialization of an [Authentication] from an http context.
type Authenticator interface {
	// Authenticate the passed context and provide the authentication and bool indicating whether the
	// context is authenticated
	Authenticate(context *Context) (Authentication, bool)
	// SetUserAuthentication on the passed context for the passed user ID.
	SetUserAuthentication(context *Context, userID string) (Authentication, bool)
}
