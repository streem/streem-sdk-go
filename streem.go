// Package streem provides sdk for initilializing streem sso token builder and generating
// tokens with necessary claims.
package streem

import (
	"errors"
	"github.com/streem/streem-sdk-go/config"
	"github.com/streem/streem-sdk-go/tokenbuilder"
)

// Init initializes streem sdk with api key id, secret and environment.
func Init(apiKeyId, apiKeySecret, apiEnvironment string) error {
	if apiKeyId == "" {
		return errors.New("Cannot initialize Streem without an API Key ID")
	}

	config.NewConfig(apiKeyId, apiKeySecret, apiEnvironment)

	return nil
}

// NewTokenBuilder creates a token builder
func NewTokenBuilder() tokenbuilder.ITokenBuilder {
	return tokenbuilder.NewTokenBuilder()
}
