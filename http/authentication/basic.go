package authentication

import (
    "encoding/base64"
    "github.com/beaconsoftwarellc/gadget/crypto"
    "github.com/beaconsoftwarellc/gadget/log"
    "github.com/beaconsoftwarellc/gadget/net"
    "github.com/beaconsoftwarellc/gadget/stringutil"
    "github.com/beaconsoftwarellc/quimby/http"
    "strings"
    "time"
)

const (
	Anonymous = "anonymous"
    Basic = "Basic"
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
    var header string
    basicAuthentication := &basicAuthentication{
        created: time.Now(),
        expiry: time.Now().Add(s.ttl),
        valid: false,
        user: Anonymous,
    }
    var success bool
    headerValues, ok := context.Request.Header[net.HeaderAuthorization]
    if len(headerValues) > 0 {
        header = headerValues[0]
    }
    // we expected 'Basic base64(USER:PASS)'
    if !ok || len(header) <= (len(Basic) + 1) {
        return basicAuthentication, success

    }
    encoded := strings.TrimSpace(header[len(Basic)+1:])
    decoded, err := base64.StdEncoding.DecodeString(encoded)
    if nil != err {
        log.Warnf("bad data in Authorization header, expected base64 encoded: %s", err)
        return basicAuthentication, success
    }
    colonSplit := strings.SplitN(string(decoded), ":", 1)
    if len(colonSplit) < 2 {
        log.Warnf("Authorization header value bad format, expected ':'")
        return basicAuthentication, success
    }
    basicAuthentication.user = colonSplit[0]
    salt, hash, ok := s.get(basicAuthentication.user)
    if ok && stringutil.ConstantTimeComparison(hash, crypto.Hash(colonSplit[1], salt)) {
        basicAuthentication.valid = true
    }
    return basicAuthentication, success
}

type basicAuthentication struct {
    user string
    created time.Time
    expiry time.Time
    valid bool
}

func (b basicAuthentication) Type() string {
    return Basic
}

func (b basicAuthentication) UserID() string {
    return b.user
}

func (b basicAuthentication) Created() time.Time {
    return b.created
}

func (b basicAuthentication) Expiry() time.Time {
    return b.expiry
}

func (b basicAuthentication) Valid() bool {
    return b.valid
}





