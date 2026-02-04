package rest_auth

import (
	"aeroline/src/domain/user_domain"
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type SessionID uuid.UUID

func (ths SessionID) String() string {
	return uuid.UUID(ths).String()
}

type DeviceID uuid.UUID

func (ths DeviceID) String() string {
	return uuid.UUID(ths).String()
}

type Session struct {
	SessionID    SessionID
	UserID       user_domain.UserID
	DeviceID     DeviceID
	LastActivity time.Time
	IsBlocked    bool
}

func (ths Session) MarshalBinary() ([]byte, error) {
	return json.Marshal(ths)
}

func (ths *Session) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, ths)
}

func NewSession(userID user_domain.UserID, deviceID DeviceID) *Session {
	return &Session{
		SessionID:    SessionID(uuid.New()),
		UserID:       userID,
		DeviceID:     deviceID,
		LastActivity: time.Now(),
		IsBlocked:    false,
	}
}

type AccessToken string

type RefreshToken string

type AccessTokenClaims struct {
	jwt.RegisteredClaims
	Permissions []string `json:"permissions"`
}

type RefreshTokenClaims struct {
	jwt.RegisteredClaims
	SessionID SessionID `json:"session_id"`
}
