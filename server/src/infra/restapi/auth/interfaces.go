package rest_auth

import (
	"aeroline/src/domain/user_domain"
	"context"
	"crypto/ecdsa"
	"time"
)

type Config struct {
	AccessTokenTTl time.Duration
	SessionTTL     time.Duration

	publicKey  *ecdsa.PublicKey
	privateKey *ecdsa.PrivateKey
}

type UserRdr interface {
	GetUserByID(
		ctx context.Context,
		userID user_domain.UserID,
	) (*user_domain.User, error)
}

type IAuthStorage interface {
	GetTokenBySessionID(ctx context.Context, sessionID SessionID) (RefreshToken, error)

	SaveToken(ctx context.Context, token RefreshToken, sessionID SessionID) error

	SaveSession(ctx context.Context, session Session) error

	GetSessionByID(ctx context.Context, sessionID SessionID) (*Session, error)

	GetUserSessions(ctx context.Context, userID user_domain.UserID) ([]Session, error)
}
