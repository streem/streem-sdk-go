package tokenbuilder

import (
	"crypto"
	"encoding/base64"
	"fmt"
	"github.com/streem/streem-sdk-go/config"

	"gopkg.in/square/go-jose.v2"
)

// getSigner creates new signer with secret as Jwk.
func getSigner() (jose.Signer, error) {

	jwk, err := getJsonWebKey()
	if err != nil {
		fmt.Println("Unable to unmarshal key from secret: ", err)
		return nil, fmt.Errorf("unable to unmarshal key from secret: %w", err)
	}

	signingKey := jose.SigningKey{
		Key:       jwk,
		Algorithm: jose.ES256,
	}
	signerOptions := &jose.SignerOptions{
		EmbedJWK: false,
	}
	signer, err := jose.NewSigner(signingKey, signerOptions)

	if err != nil {
		return nil, fmt.Errorf("unable to create signer: %w", err)
	}

	return signer, nil
}

// getJsonWebKey converts provided secret key into jwk.
// It mimics Streem SDK with thumbprint as kid.
func getJsonWebKey() (jwk jose.JSONWebKey, err error) {
	secretBytes, err := base64.RawStdEncoding.DecodeString(config.Get().ApiKeySecret)
	if err != nil {
		return
	}

	err = jwk.UnmarshalJSON(secretBytes)
	if err != nil {
		return
	}

	thumb, err := jwk.Thumbprint(crypto.SHA256)
	if err != nil {
		return
	}

	jwk.KeyID = base64.RawURLEncoding.EncodeToString(thumb)
	return
}
