package authentication

import (
	"time"

	"github.com/beaconsoftwarellc/quimby/v2/http"
)

const Reject = "reject"

func NewRejectAll() http.Authenticator {
	return &rejectAllAuthenticator{}
}

type rejectAllAuthenticator struct{}

func (r *rejectAllAuthenticator) SetUserAuthentication(context *http.Context, userID string) (http.Authentication, bool) {
	context.Authentication = &rejectionAuthentication{}
	return context.Authentication, context.Authentication.GetValidity()
}

func (r *rejectAllAuthenticator) Authenticate(*http.Context) (http.Authentication, bool) {
	return &rejectionAuthentication{}, false
}

type rejectionAuthentication struct {
	ctime time.Time
}

func (r *rejectionAuthentication) GetKind() string {
	return Reject
}

func (r *rejectionAuthentication) GetUserID() string {
	return ""
}

func (r *rejectionAuthentication) GetCreated() time.Time {
	return r.ctime
}

func (r *rejectionAuthentication) GetExpiry() time.Time {
	return r.ctime
}

func (r *rejectionAuthentication) GetValidity() bool {
	return false
}
