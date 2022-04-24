package service

import (
	"fmt"
	"time"

	"github.com/Yujiman/e_commerce/auth/jwt/internal/config"

	"github.com/dgrijalva/jwt-go"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateAccessTokenString(claims *config.AccessTokenClaims, privateKey []byte) (string, error) {
	rsaKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", status.Error(codes.Code(500), err.Error())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(rsaKey)
	if err != nil {
		return "", status.Error(codes.Code(500), err.Error())
	}

	return signedToken, nil
}

func CreateRefreshTokenString(claims *config.RefreshTokenClaims, privateKey []byte) (string, error) {
	rsaKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", status.Error(codes.Code(500), err.Error())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(rsaKey)
	if err != nil {
		return "", status.Error(codes.Code(500), err.Error())
	}

	return signedToken, nil
}

func VerifyRefreshTokenString(refreshToken string, keyPublic []byte) (*config.RefreshTokenClaims, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyPublic)
	if err != nil {
		return nil, status.Error(codes.Code(500), "VerifyRefreshTokenString parse public_key:"+err.Error())
	}
	token, err := jwt.ParseWithClaims(refreshToken, &config.RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, status.Error(
				codes.Code(401),
				fmt.Sprintf("Unexpected signing method: %v\n", token.Header["alg"]),
			)
		}
		return key, nil
	})
	if err != nil {
		return nil, status.Error(codes.Code(401), err.Error())
	}

	claims, ok := token.Claims.(*config.RefreshTokenClaims)

	if !ok {
		return nil, status.Error(codes.Code(401), "Couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, status.Error(codes.Code(401), "JWT is expired")
	}

	return claims, nil
}

func VerifyAccessTokenString(accessToken string, keyPublic []byte) (*config.AccessTokenClaims, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyPublic)
	if err != nil {
		return nil, status.Error(codes.Code(500), "VerifyAccessTokenString parse public_key:"+err.Error())
	}
	token, err := jwt.ParseWithClaims(accessToken, &config.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, status.Error(
				codes.Code(401),
				fmt.Sprintf("Unexpected signing method: %v\n", token.Header["alg"]),
			)
		}
		return key, nil
	})
	if err != nil {
		return nil, status.Error(codes.Code(401), err.Error())
	}

	claims, ok := token.Claims.(*config.AccessTokenClaims)

	if !ok {
		return nil, status.Error(codes.Code(401), "Couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, status.Error(codes.Code(401), "JWT is expired")
	}

	return claims, nil
}
