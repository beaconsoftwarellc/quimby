package authentication

import (
    "github.com/beaconsoftwarellc/quimby/http"
    "time"
)

const Reject = "reject"

func NewRejectAll() http.Authenticator {
    return &rejectAllAuthenticator{}
}

type rejectAllAuthenticator struct {}

func (r *rejectAllAuthenticator) SetUserAuthentication(context *http.Context, userID string) (http.Authentication, bool) {
    context.Authentication = &rejectionAuthentication{}
    return context.Authentication, context.Authentication.Valid()
}

func (r *rejectAllAuthenticator) Authenticate(*http.Context) (http.Authentication, bool) {
    return &rejectionAuthentication{}, false
}

type rejectionAuthentication struct {
    ctime time.Time
}

func (r *rejectionAuthentication) Type() string {
    return Reject
}

func (r *rejectionAuthentication) UserID() string {
    return ""
}

func (r *rejectionAuthentication) Created() time.Time {
    return r.ctime
}

func (r *rejectionAuthentication) Expiry() time.Time {
    return r.ctime
}

func (r *rejectionAuthentication) Valid() bool {
    return false
}
