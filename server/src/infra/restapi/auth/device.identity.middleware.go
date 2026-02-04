package rest_auth

import (
	"strings"

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
			SameSite: "Lax",
		})
	}

	ua := c.Get("User-Agent")
	meta := Device{
		DeviceID:  deviceID,
		IP:        c.IP(),
		UserAgent: ua,
		OS:        inferOS(ua),
		Browser:   inferBrowser(ua),
	}

	c.Locals("deviceID", &deviceID)
	c.Locals("deviceMeta", &meta)

	return c.Next()
}

func inferOS(ua string) string {
	switch {
	case strings.Contains(ua, "Android"):
		return "Android"
	case strings.Contains(ua, "iPhone"), strings.Contains(ua, "iPad"):
		return "iOS"
	case strings.Contains(ua, "Windows"):
		return "Windows"
	case strings.Contains(ua, "Macintosh"):
		return "MacOS"
	case strings.Contains(ua, "Linux"):
		return "Linux"
	default:
		return "Unknown"
	}
}

func inferBrowser(ua string) string {
	switch {
	case strings.Contains(ua, "Firefox"):
		return "Firefox"
	case strings.Contains(ua, "Chrome"):
		return "Chrome"
	case strings.Contains(ua, "Safari") && !strings.Contains(ua, "Chrome"):
		return "Safari"
	case strings.Contains(ua, "Edge"):
		return "Edge"
	default:
		return "Unknown"
	}
}
