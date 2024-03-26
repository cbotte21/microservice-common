package jwtParser

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt/v5"
)

type JwtParser struct {
	ClientId, ClientSecret, Realm string
	keycloak                      *gocloak.GoCloak
}

func NewJwtParser(clientId, clientSecret, realm, url string) *JwtParser {
	return &JwtParser{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Realm:        realm,
		keycloak:     gocloak.NewClient(url),
	}
}

// Redeem returns the contents of a jwt
func (secret JwtParser) Redeem(accessToken string) (*jwt.MapClaims, error) {
	token, claims, err := secret.keycloak.DecodeAccessToken(context.Background(), accessToken, secret.Realm)
	if err != nil {
		return nil, err
	}
	fmt.Println(token)
	fmt.Println(claims)
	return claims, nil
}

// ValidateJWT returns nil if the token is valid
func (secret JwtParser) ValidateJWT(accessToken string) error {
	_, _, err := secret.keycloak.DecodeAccessToken(context.Background(), accessToken, secret.Realm)
	return err
}
