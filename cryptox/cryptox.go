package cryptox

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

func GenerateJwkKeyPair() (*jwk.Key, *jwk.Key, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate private RSA k-ey: %v", err)
	}
	privateJwk, err := jwk.FromRaw(privateKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create private JWK: %v", err)
	}
	publicJwk, err := jwk.PublicKeyOf(privateJwk)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create public JWK: %v", err)
	}
	return &privateJwk, &publicJwk, nil
}

func GenerateJwkKeyPairFromPrivateKey(privateKey *rsa.PrivateKey) (*jwk.Key, *jwk.Key, error) {
	privateJwk, err := jwk.FromRaw(privateKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create private JWK: %v", err)
	}
	publicJwk, err := jwk.PublicKeyOf(privateJwk)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create public JWK: %v", err)
	}
	return &privateJwk, &publicJwk, nil
}