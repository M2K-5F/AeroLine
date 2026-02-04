package rest_auth

import (
	"aeroline/src/domain/user_domain"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type SessionResponse struct {
	SessionID    string    `json:"session_id"`
	UserID       string    `json:"user_id"`
	DeviceID     string    `json:"device_id"`
	IP           string    `json:"ip"`
	OS           string    `json:"os"`
	Browser      string    `json:"browser"`
	UserAgent    string    `json:"user_agent"`
	LastActivity time.Time `json:"last_activity"`
	IsBlocked    bool      `json:"is_blocked"`
}

type sessionDTO struct {
	SessionResponse

	CurrentRefreshToken string `json:"current_refresh_token"`
}

func SessionToResponse(session Session) SessionResponse {
	return SessionResponse{
		SessionID:    session.SessionID.String(),
		UserID:       session.UserID.String(),
		DeviceID:     session.Device.DeviceID.String(),
		IP:           session.Device.IP,
		OS:           session.Device.OS,
		Browser:      session.Device.Browser,
		UserAgent:    session.Device.UserAgent,
		LastActivity: session.LastActivity,
		IsBlocked:    session.IsBlocked,
	}
}

func (d sessionDTO) MarshalBinary() ([]byte, error)     { return json.Marshal(d) }
func (d *sessionDTO) UnmarshalBinary(data []byte) error { return json.Unmarshal(data, d) }

func (ths Session) toDTO() sessionDTO {
	return sessionDTO{
		SessionResponse: SessionResponse{
			SessionID:    ths.SessionID.String(),
			UserID:       ths.UserID.String(),
			DeviceID:     ths.Device.DeviceID.String(),
			IP:           ths.Device.IP,
			OS:           ths.Device.OS,
			Browser:      ths.Device.Browser,
			UserAgent:    ths.Device.UserAgent,
			LastActivity: ths.LastActivity,
			IsBlocked:    ths.IsBlocked,
		},
		CurrentRefreshToken: string(ths.CurrentRefreshToken),
	}
}

func (dto sessionDTO) toDomain() Session {
	var userID user_domain.UserID
	userID.Scan(dto.UserID)

	return Session{
		SessionID: SessionID(uuid.MustParse(dto.SessionID)),
		UserID:    userID,
		Device: Device{
			DeviceID:  DeviceID(uuid.MustParse(dto.DeviceID)),
			IP:        dto.IP,
			OS:        dto.OS,
			Browser:   dto.Browser,
			UserAgent: dto.UserAgent,
		},
		LastActivity:        dto.LastActivity,
		IsBlocked:           dto.IsBlocked,
		CurrentRefreshToken: RefreshToken(dto.CurrentRefreshToken),
	}
}
