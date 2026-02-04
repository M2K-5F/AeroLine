package rest_auth

import (
	"aeroline/src/domain/shared"
	"aeroline/src/domain/user_domain"
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	storage IAuthStorage
	userRdr UserRdr
	config  Config
}

func (ths AuthService) Login(ctx context.Context, user *user_domain.User, deviceID DeviceID) (AccessToken, RefreshToken, error) {
	session := NewSession(user.ID(), deviceID)

	now := time.Now()

	accessPlain, err := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		AccessTokenClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Subject:   user.ID().String(),
				IssuedAt:  jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(now.Add(ths.config.AccessTokenTTl)),
			},
			Permissions: shared.Map(user.Permissions(),
				func(p user_domain.Permission) string {
					return p.String()
				},
			),
		},
	).SignedString(ths.config.privateKey)
	if err != nil {
		return "", "", err
	}

	refreshPlain, err := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		RefreshTokenClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Subject:   user.ID().String(),
				IssuedAt:  jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(now.Add(ths.config.SessionTTL)),
			},
			SessionID: session.SessionID,
		},
	).SignedString(ths.config.privateKey)
	if err != nil {
		return "", "", err
	}

	accessToken := AccessToken(accessPlain)
	refreshToken := RefreshToken(refreshPlain)

	err = ths.storage.SaveToken(ctx, refreshToken, session.SessionID)
	if err != nil {
		return "", "", err
	}

	err = ths.storage.SaveSession(ctx, *session)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (ths AuthService) RefreshToken(ctx context.Context, token RefreshToken, deviceID DeviceID) (AccessToken, RefreshToken, error) {
	var claims *RefreshTokenClaims

	if token, err := jwt.ParseWithClaims(
		string(token),
		&RefreshTokenClaims{},
		func(t *jwt.Token) (any, error) { return ths.config.publicKey, nil },
	); err != nil {
		return "", "", err
	} else {
		claims = token.Claims.(*RefreshTokenClaims)
	}

	var userID user_domain.UserID
	userID.Scan(claims.Subject)

	cachedToken, err := ths.storage.GetTokenBySessionID(ctx, claims.SessionID)
	if err != nil {
		return "", "", err
	}

	if cachedToken != token {
		return "", "", ErrRefreshTokenInvalid
	}

	session, err := ths.storage.GetSessionByID(ctx, claims.SessionID)
	if err != nil {
		return "", "", err
	}

	if session.DeviceID != deviceID {
		return "", "", ErrUndefinedDevice
	}

	if session.IsBlocked {
		return "", "", ErrSessionBlocked
	}

	now := time.Now()

	refreshPlain, err := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		RefreshTokenClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Subject:   userID.String(),
				IssuedAt:  jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(now.Add(ths.config.SessionTTL)),
			},
			SessionID: session.SessionID,
		},
	).SignedString(ths.config.privateKey)
	if err != nil {
		return "", "", err
	}

	refreshToken := RefreshToken(refreshPlain)

	err = ths.storage.SaveToken(ctx, refreshToken, session.SessionID)
	if err != nil {
		return "", "", err
	}

	session.LastActivity = time.Now()

	err = ths.storage.SaveSession(ctx, *session)
	if err != nil {
		return "", "", err
	}

	user, err := ths.userRdr.GetUserByID(ctx, userID)
	if err != nil {
		return "", "", err
	}

	accessPlain, err := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		AccessTokenClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Subject:   user.ID().String(),
				IssuedAt:  jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(now.Add(ths.config.AccessTokenTTl)),
			},
			Permissions: shared.Map(user.Permissions(),
				func(p user_domain.Permission) string {
					return p.String()
				},
			),
		},
	).SignedString(ths.config.privateKey)
	if err != nil {
		return "", "", err
	}

	accessToken := AccessToken(accessPlain)

	return accessToken, refreshToken, nil
}

func (ths AuthService) VerifyAccessToken(ctx context.Context, token AccessToken) (*user_domain.UserID, *[]user_domain.Permission, error) {
	var claims *AccessTokenClaims

	if token, err := jwt.ParseWithClaims(
		string(token),
		&AccessTokenClaims{},
		func(t *jwt.Token) (any, error) { return ths.config.publicKey, nil },
	); err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, nil, ErrTokenExpired
		}
		return nil, nil, err
	} else {
		claims = token.Claims.(*AccessTokenClaims)
	}

	var userID user_domain.UserID
	userID.Scan(claims.RegisteredClaims.Subject)

	permissions := shared.Map(claims.Permissions, func(p string) user_domain.Permission {
		return user_domain.Permission(p)
	})

	return &userID, &permissions, nil
}

func (ths AuthService) GetUserSessions(ctx context.Context, userID user_domain.UserID) ([]Session, error) {
	return ths.storage.GetUserSessions(ctx, userID)
}

func NewAuthService(config Config, userRdr UserRdr) *AuthService {
	return &AuthService{
		config:  config,
		userRdr: userRdr,
		storage: NewAuthStorage(config),
	}
}

func NewConfigFromEnv() Config {
	priv, _ := os.ReadFile(os.Getenv("PRIVATE_KEY_PATH"))
	pub, _ := os.ReadFile(os.Getenv("PUBLIC_KEY_PATH"))

	privateKey, err := jwt.ParseECPrivateKeyFromPEM(priv)
	if err != nil {
		panic(err)
	}

	publicKey, err := jwt.ParseECPublicKeyFromPEM(pub)
	if err != nil {
		panic(err)
	}

	return Config{
		publicKey:      publicKey,
		privateKey:     privateKey,
		AccessTokenTTl: 10 * time.Minute,
		SessionTTL:     7 * 24 * time.Hour,
	}
}
