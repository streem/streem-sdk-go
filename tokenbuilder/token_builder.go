package tokenbuilder

import (
	"errors"
	"time"
)

type ITokenBuilder interface {
	SetUserId(string)
	GetUserId() string

	SetName(string)
	GetName() string

	SetEmail(string)
	GetEmail() string

	SetAvatarUrl(string)
	GetAvatarUrl() string

	SetTokenExpirationMs(int64)
	GetTokenExpirationMs() int64

	SetSessionExpirationMs(int64)
	GetSessionExpirationMs() int64

	SetReservationSid(string)
	GetReservationSid() string

	Build() (string, error)
	AddTime() func(dur time.Duration) int64
}

// tokenBuilder struct for handling token generation.
type tokenBuilder struct {
	userId              string
	name                string
	email               string
	avatarUrl           string
	tokenExpirationMs   int64
	sessionExpirationMs int64
	reservationSid      string
}

func NewTokenBuilder() *tokenBuilder {
	return &tokenBuilder{}
}

// SetUserId sets the user id onto the token builder. This is required for building a token.
func (t *tokenBuilder) SetUserId(uid string) {
	t.userId = uid
}

// GetUserId returns the user id.
func (t *tokenBuilder) GetUserId() string {
	return t.userId
}

// SetName sets the name onto the token builder. This is recommended for building a token.
func (t *tokenBuilder) SetName(name string) {
	t.name = name
}

// GetName returns the name.
func (t *tokenBuilder) GetName() string {
	return t.name
}

// SetEmail sets the email onto the token builder. This is recommended for building a token.
func (t *tokenBuilder) SetEmail(email string) {
	t.email = email
}

// GetEmail returns the email.
func (t *tokenBuilder) GetEmail() string {
	return t.email
}

// SetAvatarUrl sets the avatar url onto the token builder. This is recommended for building a token.
func (t *tokenBuilder) SetAvatarUrl(avatarUrl string) {
	t.avatarUrl = avatarUrl
}

// GetAvatarUrl returns the avatar url.
func (t *tokenBuilder) GetAvatarUrl() string {
	return t.avatarUrl
}

// SetTokenExpirationMs sets the token expiration in MS onto the token builder.
// This is optional for building a token. Tokens by default expire in 5 minutes.
func (t *tokenBuilder) SetTokenExpirationMs(tokenExpirationMs int64) {
	t.tokenExpirationMs = tokenExpirationMs
}

// GetTokenExpirationMs returns the token expiration in ms.
func (t *tokenBuilder) GetTokenExpirationMs() int64 {
	return t.tokenExpirationMs
}

// SetSessionExpirationMs sets the session expiration onto the token builder.
// This is optional for building a token. Session by default expire in 4 hours.
func (t *tokenBuilder) SetSessionExpirationMs(sessionExpirationMs int64) {
	t.sessionExpirationMs = sessionExpirationMs
}

// GetSessionExpirationMs returns the session expiration in ms.
func (t *tokenBuilder) GetSessionExpirationMs() int64 {
	return t.sessionExpirationMs
}

// SetReservationSid sets the reservation sid onto the token builder. This is optional for building a token.
func (t *tokenBuilder) SetReservationSid(reservationSid string) {
	t.reservationSid = reservationSid
}

// GetReservationSid returns the reservation sid.
func (t *tokenBuilder) GetReservationSid() string {
	return t.reservationSid
}

// Build generates JWS token for the provided user info param.
// It returns the JWS signed compact serialized token or any error encountered.
func (t *tokenBuilder) Build() (string, error) {

	if t.userId == "" {
		return "", errors.New("cannot build token with an empty userid")
	}

	addDurationToNow := t.AddTime()
	if t.tokenExpirationMs == 0 {
		t.SetTokenExpirationMs(addDurationToNow(time.Minute * 5))
	}

	if t.sessionExpirationMs == 0 {
		t.SetSessionExpirationMs(addDurationToNow(time.Hour * 4))
	}

	signer, err := getSigner()

	if err != nil {
		return "", err
	}

	claim, err := newClaim(t)

	if err != nil {
		return "", err
	}

	jws, err := signer.Sign(claim)

	if err != nil {
		return "", err
	}

	token, err := jws.CompactSerialize()

	if err != nil {
		return "", err
	}

	return token, nil
}

// AddTime adds a specific duration of time to time.Now and returns it in ms.
func (t *tokenBuilder) AddTime() func(dur time.Duration) int64 {
	now := time.Now()
	return func(dur time.Duration) int64 {
		return now.Add(dur).UnixNano() / 1e6
	}
}
