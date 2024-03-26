package jwtParser

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
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
func (secret JwtParser) Redeem(accessToken string) (*gocloak.IntroSpectTokenResult, error) {
	rptResult, err := secret.keycloak.RetrospectToken(context.Background(), accessToken, secret.ClientId, secret.ClientSecret, secret.Realm)
	if err != nil {
		panic("Inspection failed:" + err.Error())
	}

	if !*rptResult.Active {
		panic("Token is not active")
	}
	return rptResult, nil
}

// ValidateJWT returns nil if the token is valid
func (secret JwtParser) ValidateJWT(accessToken string) error {
	rptResult, err := secret.keycloak.RetrospectToken(context.Background(), accessToken, secret.ClientId, secret.ClientSecret, secret.Realm)
	if err != nil {
		panic("Inspection failed:" + err.Error())
	}

	if !*rptResult.Active {
		panic("Token is not active")
	}
	return nil
}
