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
