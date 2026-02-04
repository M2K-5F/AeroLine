package rest_auth

import "github.com/golang-jwt/jwt/v5"

type AccessToken string

type RefreshToken string

type AccessTokenClaims struct {
	jwt.RegisteredClaims
	Permissions []string `json:"permissions"`
}

type RefreshTokenClaims struct {
	jwt.RegisteredClaims
	SessionID string `json:"session_id"`
}
