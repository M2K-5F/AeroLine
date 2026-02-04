package rest_auth

import (
	"aeroline/src/domain/user_domain"
	"time"

	"github.com/google/uuid"
)

type SessionID uuid.UUID

func (ths SessionID) String() string {
	return uuid.UUID(ths).String()
}

type Session struct {
	SessionID           SessionID
	UserID              user_domain.UserID
	Device              Device
	LastActivity        time.Time
	IsBlocked           bool
	CurrentRefreshToken RefreshToken
}

func (ths *Session) UpdateToken(newToken *RefreshToken) {
	ths.CurrentRefreshToken = *newToken
}

func (ths *Session) UpdateActivity() {
	ths.LastActivity = time.Now()
}

func NewSession(userID user_domain.UserID, device *Device) *Session {
	return &Session{
		SessionID:    SessionID(uuid.New()),
		UserID:       userID,
		Device:       *device,
		LastActivity: time.Now(),
		IsBlocked:    false,
	}
}
