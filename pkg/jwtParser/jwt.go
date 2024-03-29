package jwtParser

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt/v5"
)

type JwtParser struct {
	Realm    string
	keycloak *gocloak.GoCloak
}

func NewJwtParser(realm, url string) *JwtParser {
	return &JwtParser{
		Realm:    realm,
		keycloak: gocloak.NewClient(url),
	}
}

// Redeem returns the contents of a jwt
func (secret JwtParser) Redeem(accessToken string) (*jwt.MapClaims, error) {
	_, claims, err := secret.keycloak.DecodeAccessToken(context.Background(), accessToken, secret.Realm)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

// ValidateJWT returns nil if the token is valid
func (secret JwtParser) ValidateJWT(accessToken string) error {
	_, _, err := secret.keycloak.DecodeAccessToken(context.Background(), accessToken, secret.Realm)
	return err
}
