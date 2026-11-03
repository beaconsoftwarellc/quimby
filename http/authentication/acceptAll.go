package authentication

import (
	"time"

	"github.com/beaconsoftwarellc/gadget/v2/stringutil"
	"github.com/beaconsoftwarellc/quimby/v2/http"
)

const (
	Accept = "accept"
)

// NewAcceptAll authenticator that will treat all incoming requests as authenticated
func NewAcceptAll[T http.ContextBase]() http.Authenticator[T] {
	return &acceptAllAuthenticator[T]{}
}

type acceptAllAuthenticator[T http.ContextBase] struct{}

func (r *acceptAllAuthenticator[T]) SetUserAuthentication(context *http.Context[T], userID string) (
	http.Authentication, bool) {
	context.Authentication = &acceptAuthentication{userID: userID}
	return context.Authentication, context.Authentication.GetValidity()
}

func (r *acceptAllAuthenticator[T]) Authenticate(context *http.Context[T]) (http.Authentication, bool) {
	context.Authentication = &acceptAuthentication{}
	return context.Authentication, context.Authentication.GetValidity()
}

type acceptAuthentication struct {
	ctime  time.Time
	userID string
}

func (r *acceptAuthentication) GetKind() string {
	return Accept
}

func (r *acceptAuthentication) GetUserID() string {
	if stringutil.IsEmpty(r.userID) {
		r.userID = Anonymous
	}
	return r.userID
}

func (r *acceptAuthentication) GetCreated() time.Time {
	return r.ctime
}

func (r *acceptAuthentication) GetExpiry() time.Time {
	return time.Now().Add(time.Minute)
}

func (r *acceptAuthentication) GetValidity() bool {
	return true
}
