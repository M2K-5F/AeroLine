package rest_auth

import (
	"aeroline/src/domain/user_domain"
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var (
	sessionPrefix     = func(sessionID SessionID) string { return "session:" + sessionID.String() }
	userSessionPrefix = func(userID user_domain.UserID) string { return "user_sessions:" + userID.String() }
)

type AuthStorage struct {
	client *redis.Client
	config Config
}

func (ths AuthStorage) SaveSession(ctx context.Context, session Session) error {
	pipe := ths.client.Pipeline()

	pipe.Set(ctx, sessionPrefix(session.SessionID), session.toDTO(), ths.config.SessionTTL)

	pipe.SAdd(ctx, userSessionPrefix(session.UserID), session.SessionID.String())
	pipe.Expire(ctx, userSessionPrefix(session.UserID), ths.config.SessionTTL)

	_, err := pipe.Exec(ctx)
	return err
}

func (ths AuthStorage) GetSessionByID(ctx context.Context, sessionID SessionID) (*Session, error) {
	var dto sessionDTO

	err := ths.client.Get(ctx, sessionPrefix(sessionID)).Scan(&dto)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrSessionNotFound
		}
		return nil, err
	}

	session := dto.toDomain()
	return &session, nil
}

func (ths AuthStorage) GetUserSessions(ctx context.Context, userID user_domain.UserID) ([]Session, error) {
	sessionIDs, err := ths.client.SMembers(ctx, userSessionPrefix(userID)).Result()
	if err != nil {
		return nil, err
	}

	sessions := make([]Session, 0, len(sessionIDs))

	for _, id := range sessionIDs {
		sess, err := ths.GetSessionByID(ctx, SessionID(uuid.MustParse(id)))
		if err != nil {
			if errors.Is(err, ErrSessionNotFound) {
				ths.client.SRem(ctx, userSessionPrefix(userID), id)
				continue
			}
			return nil, err
		}
		sessions = append(sessions, *sess)
	}

	return sessions, nil
}

func (ths AuthStorage) DeleteSession(ctx context.Context, userID user_domain.UserID, sessionID SessionID) error {
	pipe := ths.client.Pipeline()

	pipe.Del(ctx, sessionPrefix(sessionID))
	pipe.SRem(ctx, userSessionPrefix(userID), sessionID.String())

	_, err := pipe.Exec(ctx)
	return err
}

func NewAuthStorage(config Config, redisClient *redis.Client) *AuthStorage {
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return &AuthStorage{
		client: redisClient,
		config: config,
	}
}
