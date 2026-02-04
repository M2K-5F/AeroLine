package rest_auth

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func DeviceIDMiddleware(c fiber.Ctx) error {
	var deviceID DeviceID

	plain := c.Cookies("devid")

	if id, err := uuid.Parse(plain); err == nil {
		deviceID = DeviceID(id)
	} else {
		deviceID = DeviceID(uuid.New())
		c.Cookie(&fiber.Cookie{
			Name:     "devid",
			Value:    uuid.UUID(deviceID).String(),
			HTTPOnly: true,
			Secure:   true,
			MaxAge:   365 * 24 * 3600,
			Path:     "/",
		})
	}

	c.Locals("deviceID", &deviceID)

	return c.Next()
}
