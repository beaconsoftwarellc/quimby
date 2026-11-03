package authentication

import (
	"time"

	"github.com/beaconsoftwarellc/quimby/v2/http"
)

const Reject = "reject"

func NewRejectAll[T http.ContextBase]() http.Authenticator[T] {
	return &rejectAllAuthenticator[T]{}
}

type rejectAllAuthenticator[T http.ContextBase] struct{}

func (r *rejectAllAuthenticator[T]) SetUserAuthentication(context *http.Context[T],
	_ string) (http.Authentication, bool) {
	context.Authentication = &rejectionAuthentication{}
	return context.Authentication, context.Authentication.GetValidity()
}

func (r *rejectAllAuthenticator[T]) Authenticate(*http.Context[T]) (http.Authentication, bool) {
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
