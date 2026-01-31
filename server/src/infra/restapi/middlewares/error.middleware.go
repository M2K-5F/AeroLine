package middlewares

import (
	"aeroline/src/domain/shared"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
)

const (
	colorReset    = "\033[0m"
	colorCritical = "\033[1;41;37m"
	colorRed      = "\033[31m"
	colorYellow   = "\033[33m"
	colorCyan     = "\033[36m"
	colorBold     = "\033[1m"
	colorGray     = "\033[90m"
)

func Error() fiber.Handler {
	return func(c fiber.Ctx) error {
		err := c.Next()

		if err == nil {
			return nil
		}
		var appError *shared.AppError
		if errors.As(err, &appError) {
			status := getStatusCode(appError.Type)
			return c.Status(status).JSON(appError)
		}

		var fiberError *fiber.Error
		if errors.As(err, &fiberError) {
			return c.Status(fiberError.Code).JSON(fiber.Map{
				"message": fiberError.Message,
			})
		}

		logInternalError(c, err)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
}

func getStatusCode(errType shared.ErrorType) int {
	switch errType {
	case shared.TypeValidation:
		return fiber.StatusUnprocessableEntity

	case
		shared.TypeAlreadyExists,
		shared.TypeBusinessLogic,
		shared.TypeMissingData,
		shared.TypeIntegrity:
		return fiber.StatusBadRequest

	case shared.TypeNotFound:
		return fiber.StatusNotFound

	case shared.TypeForbidden:
		return fiber.StatusForbidden

	case shared.TypeUnauthorized:
		return fiber.StatusUnauthorized
	}

	return fiber.StatusUnprocessableEntity
}

func logInternalError(c fiber.Ctx, err error) {
	timestamp := time.Now().Format("15:04:05")
	method := c.Method()
	path := c.Path()

	fmt.Printf("%s%s %s%s%s %s%s  %s%-15s  %s500%s  \n\t %s[CRITICAL ERROR]%s %v%s\n",
		colorGray, timestamp,
		colorCritical, "[WARNING]", colorReset,
		getMethodColorSimple(method), method,
		colorCyan, path,
		colorCritical, colorReset,
		colorRed,
		colorYellow, err, colorReset,
	)
}
