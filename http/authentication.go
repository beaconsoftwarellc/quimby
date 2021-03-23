package http

// Authentication provides a base for authentication implementations that are used as part of the http
// context chain.
type Authentication interface {
	// Type of authentication this instance provides
	GetType() string
	// UserID that is requesting authentication
	GetUserID() string
	// Valid indicates whether the user id on this instance is authentic
	IsValid() bool
}

// Authenticator is responsible for the initialization of a Authentication from a http context.
type Authenticator interface {
	// Authenticate the passed context and provide the authentication and bool indicating whether the
	// context is authenticated
	Authenticate(context *Context) (Authentication, bool)
	// SetUserAuthentication on the passed context for the passed user ID.
	SetUserAuthentication(context *Context, userID string) (Authentication, bool)
}
