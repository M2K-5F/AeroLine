package rest_auth

import (
	"aeroline/src/domain/user_domain"

	"github.com/gofiber/fiber/v3"
)

func GetUserID(c fiber.Ctx) *user_domain.UserID {
	userID, ok := c.Locals("user-id").(user_domain.UserID)
	if !ok {
		panic("You need add PermissionFilter to route to use user id in controllers")
	}

	return &userID
}

func GetDeviceID(c fiber.Ctx) *DeviceID {
	o, ok := c.Locals("deviceID").(*DeviceID)
	if !ok {
		panic("You need add device id middleware to use device id in controllers")
	}
	return o
}

func GetDevice(c fiber.Ctx) *Device {
	device, ok := c.Locals("deviceMeta").(*Device)
	if !ok {
		panic("You need add device id middleware to use device in controllers")
	}

	return device
}

func GetRefreshToken(c fiber.Ctx) (*RefreshToken, error) {
	plain := c.Cookies("refr")
	if plain == "" {
		return nil, ErrTokenNotFound
	}
	token := RefreshToken(plain)

	return &token, nil
}
