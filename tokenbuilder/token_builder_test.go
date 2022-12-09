package tokenbuilder

import (
	"fmt"
	"github.com/streem/streem-sdk-go/config"
	"testing"
)

func TestUserId(t *testing.T) {
	var tests = []struct {
		expectedUserId string
	}{
		{expectedUserId: "123abc"},
		{expectedUserId: "!*@&$!()++"},
		{expectedUserId: ""},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("(%d): setting userId %s. Expecting userId %s", i, tt.expectedUserId, tt.expectedUserId)
		t.Run(testName, func(t *testing.T) {
			instance := NewTokenBuilder()

			instance.SetUserId(tt.expectedUserId)
			userId := instance.GetUserId()

			if userId != tt.expectedUserId {
				t.Errorf("got %s, want %s", userId, tt.expectedUserId)
			}
		})
	}
}

func TestName(t *testing.T) {
	var tests = []struct {
		expectedName string
	}{
		{expectedName: "Alex"},
		{expectedName: "A Name With A Lot of Spaces And $p3C14L Characters"},
		{expectedName: ""},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("(%d): setting name %s. Expecting name %s", i, tt.expectedName, tt.expectedName)
		t.Run(testName, func(t *testing.T) {
			instance := NewTokenBuilder()

			instance.SetName(tt.expectedName)
			name := instance.GetName()

			if name != tt.expectedName {
				t.Errorf("got %s, want %s", name, tt.expectedName)
			}
		})
	}
}

func TestEmail(t *testing.T) {
	var tests = []struct {
		expectedEmail string
	}{
		{expectedEmail: "some@email.com"},
		{expectedEmail: "some@email"},
		{expectedEmail: ""},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("(%d): setting email %s. Expecting email %s", i, tt.expectedEmail, tt.expectedEmail)
		t.Run(testName, func(t *testing.T) {
			instance := NewTokenBuilder()

			instance.SetEmail(tt.expectedEmail)
			email := instance.GetEmail()

			if email != tt.expectedEmail {
				t.Errorf("got %s, want %s", email, tt.expectedEmail)
			}
		})
	}
}

func TestAvatarUrl(t *testing.T) {
	var tests = []struct {
		expectedAvatarUrl string
	}{
		{expectedAvatarUrl: "some/avatar/url"},
		{expectedAvatarUrl: "https://some.avatar.url"},
		{expectedAvatarUrl: ""},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("(%d): setting avatar url %s. Expecting avatar url %s", i, tt.expectedAvatarUrl, tt.expectedAvatarUrl)
		t.Run(testName, func(t *testing.T) {
			instance := NewTokenBuilder()

			instance.SetAvatarUrl(tt.expectedAvatarUrl)
			avatarUrl := instance.GetAvatarUrl()

			if avatarUrl != tt.expectedAvatarUrl {
				t.Errorf("got %s, want %s", avatarUrl, tt.expectedAvatarUrl)
			}
		})
	}
}

func TestTokenExpirationMs(t *testing.T) {
	var tests = []struct {
		tokenExpiration int64
	}{
		{tokenExpiration: 1},
		{tokenExpiration: 1e6},
		{tokenExpiration: -10},
		{tokenExpiration: 0},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("(%d): setting token expiration %d. Expecting token expiration %d", i, tt.tokenExpiration, tt.tokenExpiration)
		t.Run(testName, func(t *testing.T) {
			instance := NewTokenBuilder()

			instance.SetTokenExpirationMs(tt.tokenExpiration)
			tokenExpiration := instance.GetTokenExpirationMs()

			if tokenExpiration != tt.tokenExpiration {
				t.Errorf("got %d, want %d", tokenExpiration, tt.tokenExpiration)
			}
		})
	}
}

func TestSessionExpirationMs(t *testing.T) {
	var tests = []struct {
		sessionExpiration int64
	}{
		{sessionExpiration: 1},
		{sessionExpiration: 1e6},
		{sessionExpiration: -1010},
		{sessionExpiration: 0},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("(%d): setting session expiration %d. Expecting session expiration %d", i, tt.sessionExpiration, tt.sessionExpiration)
		t.Run(testName, func(t *testing.T) {
			instance := NewTokenBuilder()

			instance.SetSessionExpirationMs(tt.sessionExpiration)
			sessionExpiration := instance.GetSessionExpirationMs()

			if sessionExpiration != tt.sessionExpiration {
				t.Errorf("got %d, want %d", sessionExpiration, tt.sessionExpiration)
			}
		})
	}
}

func TestReservationSid(t *testing.T) {
	var tests = []struct {
		expectedReservationSid string
	}{
		{expectedReservationSid: "rsv_abc123"},
		{expectedReservationSid: ""},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("(%d): setting reservation sid %s. Expecting reservation sid %s", i, tt.expectedReservationSid, tt.expectedReservationSid)
		t.Run(testName, func(t *testing.T) {
			instance := NewTokenBuilder()

			instance.SetReservationSid(tt.expectedReservationSid)
			reservationSid := instance.GetReservationSid()

			if reservationSid != tt.expectedReservationSid {
				t.Errorf("got %s, want %s", reservationSid, tt.expectedReservationSid)
			}
		})
	}
}

func TestBuild(t *testing.T) {
	var tests = []struct {
		givenInstance func() ITokenBuilder
	}{
		{
			givenInstance: func() ITokenBuilder {

			    // These values are from our Embedded SSO notion document
			    // https://www.notion.so/streemers/Embedded-SSO-d12c4dd8864b4304ad2d023446d59cf0#daced2833ad34cfabf395786fbd3d4d3
				config.NewConfig("api_1mY3yMnSp4DUa97vnSgrOW",
				    "eyJrdHkiOiJFQyIsImQiOiI1cTJCTE5CTG8wR2tCdjJhbHZjaU9VQjh2M0tCQWZUYU02VVd4TDllN3lBIiwidXNlIjoic2lnIiwiY3J2IjoiUC0yNTYiLCJ4IjoiblpEVHNZVlpKVGtsTnBpU19SaVYxdkVGaEZLVVgtVHpNWUxfVjhuSXpSdyIsInkiOiJFeXRLSW1KeEtkQlctVHpWZjBhNmhxTEpJV0R6cHdfTnBkTHZhd1VNOVBRIiwiYWxnIjoiRVMyNTYifQ",
				    "someEnvironment",
				)
				i := NewTokenBuilder()
				i.SetUserId("123ID")

				return i
			},
		},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("(%d): Testing TokenBuilder#Build.", i)
		t.Run(testName, func(t *testing.T) {

			instance := tt.givenInstance()
			token, err := instance.Build()

			if err != nil {
				t.Errorf("got error %v", err)
			}

			if token == "" {
                t.Errorf("got empty token %s", token)
            }
		})
	}
}
