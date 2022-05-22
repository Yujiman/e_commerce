package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/dgrijalva/jwt-go"
)

type AccessTokenClaims struct {
	DomainID string `json:"domain_id,omitempty"`
	Scopes   string `json:"scopes,omitempty"`
	jwt.StandardClaims
}

var onceInitAccessExpiredMinutes sync.Once
var accessExpiredMinutes *int

func GetAccessTokenLifeTimeMinutes() int {
	onceInitAccessExpiredMinutes.Do(func() {
		expiresAt, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRES_AT"))
		if err != nil {
			log.Panicln("ACCESS_TOKEN_EXPIRES_AT environment not valid")
		}
		accessExpiredMinutes = &expiresAt
	})
	return *accessExpiredMinutes
}

var onceInitAccessClaims sync.Once
var accessClaims *AccessTokenClaims

func GetAccessTokenClaims() *AccessTokenClaims {
	onceInitAccessClaims.Do(func() {
		issuer := os.Getenv("TOKEN_ISSUER")
		if issuer == "" {
			log.Panicln("TOKEN_ISSUER environment not valid")
		}

		audience := os.Getenv("TOKEN_AUDIENCE")
		if audience == "" {
			log.Panicln("TOKEN_AUDIENCE environment not valid")
		}

		accessClaims = &AccessTokenClaims{
			Scopes: "[]",
			StandardClaims: jwt.StandardClaims{
				Issuer:   issuer,
				Audience: audience,
			},
		}
	})

	return accessClaims
}

type RefreshTokenClaims struct {
	AccessTokenClaims AccessTokenClaims
	jwt.StandardClaims
}

func GetRefreshTokenClaims() *RefreshTokenClaims {
	return &RefreshTokenClaims{}
}

var onceInitRefreshExpiredMinutes sync.Once
var refreshExpiredMinutes *int

func GetRefreshTokenLifeTimeMinutes() int {
	onceInitRefreshExpiredMinutes.Do(func() {
		expiresAt, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRES_AT"))
		if err != nil {
			log.Panicln("REFRESH_TOKEN_EXPIRES_AT environment not valid")
		}
		refreshExpiredMinutes = &expiresAt
	})

	return *refreshExpiredMinutes
}
