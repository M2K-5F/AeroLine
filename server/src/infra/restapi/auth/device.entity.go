package rest_auth

import (
	"encoding/json"

	"github.com/google/uuid"
)

type DeviceID uuid.UUID

func (ths DeviceID) String() string {
	return uuid.UUID(ths).String()
}

func (ths DeviceID) MarshalJSON() ([]byte, error) {
	return json.Marshal(uuid.UUID(ths).String())
}

func (ths *DeviceID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	id, err := uuid.Parse(s)
	if err != nil {
		return err
	}
	*ths = DeviceID(id)
	return nil
}

type Device struct {
	DeviceID  DeviceID `json:"id"`
	IP        string   `json:"ip"`
	OS        string   `json:"os"`
	Browser   string   `json:"browser"`
	UserAgent string   `json:"user_agent"`
}
