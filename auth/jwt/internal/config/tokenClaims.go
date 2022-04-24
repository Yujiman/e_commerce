package config

import (
	"github.com/dgrijalva/jwt-go"
)

const AccessExpiredMinutes = 20
const RefreshExpiredMinutes = 2000

type AccessTokenClaims struct {
	Scopes string `json:"scopes,omitempty"`
	jwt.StandardClaims
}

func GetAccessTokenClaims() *AccessTokenClaims {

	return &AccessTokenClaims{
		Scopes: "[]",
		StandardClaims: jwt.StandardClaims{
			Issuer: "eCommerce",
		},
	}
}

type RefreshTokenClaims struct {
	AccessTokenClaims AccessTokenClaims
	jwt.StandardClaims
}

func GetRefreshTokenClaims() *RefreshTokenClaims {
	return &RefreshTokenClaims{}
}
