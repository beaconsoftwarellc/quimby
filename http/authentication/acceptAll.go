package authentication

import (
    "github.com/beaconsoftwarellc/quimby/http"
    "time"
)

const (
	Accept = "accept"
)

// NewAcceptAll authenticator that will treat all incoming requests as authenticated
func NewAcceptAll() http.Authenticator {
    return &acceptAllAuthenticator{}
}

type acceptAllAuthenticator struct {}

func (r acceptAllAuthenticator) Authenticate(*http.Context) (http.Authentication, bool) {
    return &acceptAuthentication{}, false
}

type acceptAuthentication struct {
    ctime time.Time
}

func (r *acceptAuthentication) Type() string {
    return Accept
}

func (r *acceptAuthentication) UserID() string {
    return Anonymous
}

func (r *acceptAuthentication) Created() time.Time {
    return r.ctime
}

func (r *acceptAuthentication) Expiry() time.Time {
    return time.Now().Add(time.Minute)
}

func (r *acceptAuthentication) Valid() bool {
    return true
}
