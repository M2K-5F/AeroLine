package rest_utils

import (
	"aeroline/src/domain/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validatorSgt = validator.New()

func ValidateAndGetBody[RequestType any](c fiber.Ctx) (*RequestType, error) {
	var request RequestType
	if err := c.Bind().Body(&request); err != nil {
		return nil, err
	}

	if err := validatorSgt.Struct(&request); err != nil {
		return nil, err
	}

	return &request, nil
}

func GetUserId(c fiber.Ctx) (user.UserID, error) {
	plainID := c.Locals("X-userID").(string)
	var userID user.UserID
	err := userID.Scan(plainID)
	if err != nil {
		return user.UserID{}, err
	}

	return userID, nil
}
