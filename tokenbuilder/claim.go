package tokenbuilder

import (
	"encoding/json"
	"fmt"
	"github.com/streem/streem-sdk-go/config"
	"time"
)

// claim struct holds properties required for token payload
type claim struct {
	Audience       string `json:"aud"`
	Email          string `json:"email"`
	ExpirationTime int64  `json:"exp"`
	SessionExp     int64  `json:"session_exp"`
	IssuedAt       int64  `json:"iat"`
	Issuer         string `json:"iss"`
	Name           string `json:"name"`
	Picture        string `json:"picture"`
	Subject        string `json:"sub"`
	ReservationSid string `json:"streem:reservation_sid"`
}

// newClaim generates a byte array for a claim
func newClaim(info *tokenBuilder) ([]byte, error) {
	claim := claim{
		Audience:       fmt.Sprintf("https://api.%s.streem.cloud/", config.Get().ApiEnvironment),
		Issuer:         fmt.Sprintf("streem:api:%s", config.Get().ApiKeyId),
		IssuedAt:       time.Now().Unix(),
		ExpirationTime: info.GetTokenExpirationMs() / 1000,
		SessionExp:     info.GetSessionExpirationMs() / 1000,
		Subject:        info.GetUserId(),
		Email:          info.GetEmail(),
		Name:           info.GetName(),
		Picture:        info.GetAvatarUrl(),
		ReservationSid: info.GetReservationSid(),
	}

	return claim.toBytes()
}

// toBytes converts claim struct to byte array using json encoded.
func (claim *claim) toBytes() ([]byte, error) {
	bytes, err := json.Marshal(claim)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
