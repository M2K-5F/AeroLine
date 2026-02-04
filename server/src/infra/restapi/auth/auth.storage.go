package rest_auth

import (
	"aeroline/src/domain/user_domain"
	"context"
	"errors"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

const (
	refreshPrefix     = "refresh:"
	sessionPrefix     = "session:"
	userSessionPrefix = "user_sessions:"
)

type AuthStorage struct {
	client *redis.Client
	config Config
}

func (ths AuthStorage) SaveToken(ctx context.Context, token RefreshToken, sessionID SessionID) error {
	return ths.client.Set(ctx, refreshPrefix+sessionID.String(), string(token), ths.config.SessionTTL).Err()
}

func (ths AuthStorage) SaveSession(ctx context.Context, session Session) error {
	pipe := ths.client.Pipeline()

	pipe.Set(ctx, sessionPrefix+session.SessionID.String(), session, ths.config.SessionTTL)

	pipe.SAdd(ctx, userSessionPrefix+session.UserID.String(), session.SessionID.String())
	pipe.Expire(ctx, userSessionPrefix+session.UserID.String(), ths.config.SessionTTL)

	_, err := pipe.Exec(ctx)
	return err
}

func (ths AuthStorage) GetSessionByID(ctx context.Context, sessionID SessionID) (*Session, error) {
	var session Session

	err := ths.client.Get(ctx, sessionPrefix+sessionID.String()).Scan(&session)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrSessionNotFound
		}
		return nil, err
	}

	return &session, nil
}

func (ths AuthStorage) GetTokenBySessionID(ctx context.Context, sessionID SessionID) (RefreshToken, error) {
	val, err := ths.client.Get(ctx, refreshPrefix+sessionID.String()).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrRefreshTokenInvalid
		}
		return "", err
	}

	return RefreshToken(val), nil
}

func (ths AuthStorage) GetUserSessions(ctx context.Context, userID user_domain.UserID) ([]Session, error) {

	sessionIDs, err := ths.client.SMembers(ctx, userSessionPrefix+userID.String()).Result()
	if err != nil {
		return nil, err
	}

	if len(sessionIDs) == 0 {
		return nil, nil
	}

	sessions := make([]Session, 0, len(sessionIDs))

	for _, id := range sessionIDs {
		sess, err := ths.GetSessionByID(ctx, SessionID(uuid.MustParse(id)))
		if err != nil {
			if err == redis.Nil || err == ErrSessionNotFound {
				ths.client.SRem(ctx, userSessionPrefix+userID.String(), id)
				continue
			}
			return nil, err
		}
		sessions = append(sessions, *sess)
	}

	return sessions, nil
}

func NewAuthStorage(config Config) *AuthStorage {
	dsn := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	client := redis.NewClient(&redis.Options{
		Addr: dsn,
		DB:   0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return &AuthStorage{
		client: client,
		config: config,
	}
}
