package rest_auth

import (
	"aeroline/src/domain/shared"
	"aeroline/src/domain/user_domain"
	"context"
	"errors"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService struct {
	storage IAuthStorage
	userRdr UserRdr
	config  Config
}

func (ths AuthService) signAccessToken(claims AccessTokenClaims) (*AccessToken, error) {
	plain, err := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		claims,
	).SignedString(ths.config.privateKey)
	if err != nil {
		return nil, err
	}

	token := AccessToken(plain)
	return &token, nil
}

func (ths AuthService) signRefreshToken(claims RefreshTokenClaims) (*RefreshToken, error) {
	plain, err := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		claims,
	).SignedString(ths.config.privateKey)
	if err != nil {
		return nil, err
	}

	token := RefreshToken(plain)
	return &token, nil
}

func (ths AuthService) verifyAccessToken(token AccessToken) (*AccessTokenClaims, error) {
	var claims *AccessTokenClaims

	if token, err := jwt.ParseWithClaims(
		string(token),
		&AccessTokenClaims{},
		func(t *jwt.Token) (any, error) { return ths.config.publicKey, nil },
	); err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, err
	} else {
		claims = token.Claims.(*AccessTokenClaims)
	}

	return claims, nil
}

func (ths AuthService) verifyRefreshToken(token RefreshToken) (*RefreshTokenClaims, error) {
	var claims *RefreshTokenClaims

	if token, err := jwt.ParseWithClaims(
		string(token),
		&RefreshTokenClaims{},
		func(t *jwt.Token) (any, error) { return ths.config.publicKey, nil },
	); err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, err
	} else {
		claims = token.Claims.(*RefreshTokenClaims)
	}

	return claims, nil
}

func (ths AuthService) Login(ctx context.Context, user *user_domain.User, device *Device) (*AccessToken, *RefreshToken, error) {
	session := NewSession(user.ID(), device)

	now := time.Now()

	accessToken, err := ths.signAccessToken(
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
	)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := ths.signRefreshToken(
		RefreshTokenClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Subject:   user.ID().String(),
				IssuedAt:  jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(now.Add(ths.config.SessionTTL)),
			},
			SessionID: session.SessionID.String(),
		},
	)
	if err != nil {
		return nil, nil, err
	}

	session.UpdateToken(refreshToken)

	if err := ths.storage.SaveSession(ctx, *session); err != nil {
		return nil, nil, err
	}

	return accessToken, refreshToken, nil
}

func (ths AuthService) RefreshToken(ctx context.Context, token RefreshToken, device *Device) (*AccessToken, *RefreshToken, error) {
	refreshClaims, err := ths.verifyRefreshToken(token)
	if err != nil {
		return nil, nil, err
	}

	session, err := ths.storage.GetSessionByID(ctx, SessionID(uuid.MustParse(refreshClaims.SessionID)))
	if err != nil {
		return nil, nil, err
	}

	cachedToken := session.CurrentRefreshToken

	if cachedToken != token {
		return nil, nil, ErrRefreshTokenInvalid
	}

	if device.DeviceID != session.Device.DeviceID {
		return nil, nil, ErrUndefinedDevice
	}

	if session.IsBlocked {
		return nil, nil, ErrSessionBlocked
	}

	now := time.Now()

	refreshToken, err := ths.signRefreshToken(
		RefreshTokenClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Subject:   session.UserID.String(),
				IssuedAt:  jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(now.Add(ths.config.SessionTTL)),
			},
			SessionID: session.SessionID.String(),
		},
	)
	if err != nil {
		return nil, nil, err
	}

	session.UpdateActivity()
	session.UpdateToken(refreshToken)

	user, err := ths.userRdr.GetUserByID(ctx, session.UserID)
	if err != nil {
		return nil, nil, err
	}

	err = ths.storage.SaveSession(ctx, *session)
	if err != nil {
		return nil, nil, err
	}

	accessToken, err := ths.signAccessToken(
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
	)
	if err != nil {
		return nil, nil, err
	}

	return accessToken, refreshToken, nil
}

func (ths AuthService) VerifyAccessToken(ctx context.Context, token AccessToken) (*user_domain.UserID, *[]user_domain.Permission, error) {
	claims, err := ths.verifyAccessToken(token)
	if err != nil {
		return nil, nil, err
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

func NewAuthService(redisClient *redis.Client, config Config, userRdr UserRdr) *AuthService {
	return &AuthService{
		config:  config,
		userRdr: userRdr,
		storage: NewAuthStorage(config, redisClient),
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
