package authentication

import (
	"time"

	"github.com/beaconsoftwarellc/gadget/crypto"
	"github.com/beaconsoftwarellc/gadget/stringutil"
	"github.com/beaconsoftwarellc/quimby/http"
)

const (
	Anonymous = "anonymous"
	Basic     = "Basic"
)

// GetHashAndSalt for a specified user identity string. Boolean indicates whether the retrieval was successful
type GetHashAndSalt func(user string) (hash string, salt string, ok bool)

// NewBasic authenticator for doing basic authentication against a http context. Expects the username and password
// to be Base64 encoded separated by a ':'. Usernames must not contain a colon however password's can.
// The information should be included in the 'Authorization' header in the following format:
// Authorization: Basic QWxhZGRpbjpPcGVuU2VzYW1l
// where 'QWxhZGRpbjpPcGVuU2VzYW1l' is the base64 encoded string 'Aladdin:OpenSesame'
func NewBasic(get GetHashAndSalt, authenticationTTL time.Duration) http.Authenticator {
	return &basicAuthenticator{
		get: get,
		ttl: authenticationTTL,
	}
}

// NewBasicHashAndSalt creates a hash and salt for the passed password that will authenticate using the BasicAuthenticator
// in this package.
func NewBasicHashAndSalt(password string) (hash, salt string) {
	return crypto.HashAndSalt(password)
}

type basicAuthenticator struct {
	get GetHashAndSalt
	ttl time.Duration
}

func (s *basicAuthenticator) Authenticate(context *http.Context) (http.Authentication, bool) {
	basicAuthentication := &basicAuthentication{
		created: time.Now(),
		expiry:  time.Now().Add(s.ttl),
		valid:   false,
		user:    Anonymous,
	}
	var password string
	var ok bool
	basicAuthentication.user, password, ok = context.Request.BasicAuth()
	if !ok {
		return basicAuthentication, false
	}
	hash, salt, ok := s.get(basicAuthentication.user)
	if ok && stringutil.ConstantTimeComparison(hash, crypto.Hash(password, salt)) {
		basicAuthentication.valid = true
	}
	context.Authentication = basicAuthentication
	return basicAuthentication, basicAuthentication.IsValid()
}

func (s *basicAuthenticator) SetUserAuthentication(context *http.Context, userID string) (http.Authentication, bool) {
	context.Authentication = &basicAuthentication{created: time.Now(), expiry: time.Now().Add(s.ttl), valid: true, user: userID}
	return context.Authentication, context.Authentication.IsValid()
}

type basicAuthentication struct {
	user    string
	created time.Time
	expiry  time.Time
	valid   bool
}

func (b basicAuthentication) GetKind() string {
	return Basic
}

func (b basicAuthentication) GetUserID() string {
	return b.user
}

func (b basicAuthentication) GetCreated() time.Time {
	return b.created
}

func (b basicAuthentication) GetExpiry() time.Time {
	return b.expiry
}

func (b basicAuthentication) IsValid() bool {
	return b.valid
}
