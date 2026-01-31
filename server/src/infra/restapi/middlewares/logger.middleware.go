package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
)

func Logger() fiber.Handler {
	return func(c fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		latency := time.Since(start)

		reset := "\033[0m"
		gray := "\033[90m"

		timestamp := time.Now().Format("15:04:05")
		timeStr := fmt.Sprintf("%7v", latency.Truncate(time.Microsecond))

		logLine := fmt.Sprintf(
			"%s%s%s  %s[INFO]%s   %s%s%s  %-15s  %s%3d%s  %s%s%s  %s%s%s",
			gray, timestamp, reset,
			"\033[1;32m", reset,
			getMethodColorSimple(c.Method()), c.Method(), reset,
			c.Path(),
			getStatusColorSimple(c.Response().StatusCode()),
			c.Response().StatusCode(), reset,
			gray, timeStr, reset,
			gray, c.IP(), reset,
		)

		fmt.Fprintln(os.Stdout, logLine)
		return err
	}
}

func getMethodColorSimple(method string) string {
	switch method {
	case "GET":
		return "\033[1;36m" // Bold Cyan
	case "POST":
		return "\033[1;32m" // Bold Green
	case "PUT":
		return "\033[1;33m" // Bold Yellow
	case "DELETE":
		return "\033[1;31m" // Bold Red
	default:
		return "\033[1;37m" // Bold White
	}
}

func getStatusColorSimple(status int) string {
	switch {
	case status >= 500:
		return "\033[1;41;37m" // Белый на Красном фоне (Критично)
	case status >= 400:
		return "\033[1;31m" // Красный
	case status >= 300:
		return "\033[1;33m" // Желтый
	case status >= 200:
		return "\033[1;32m" // Зеленый
	default:
		return "\033[1;36m" // Циан
	}
}
