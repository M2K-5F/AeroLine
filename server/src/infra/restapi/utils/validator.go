package rest_utils

import (
	"aeroline/src/domain/user_domain"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validatorSgt = validator.New()

func ValidateAndGetCommand[RequestType any](c fiber.Ctx) (*RequestType, error) {
	var request RequestType
	if err := c.Bind().Body(&request); err != nil {
		return nil, err
	}

	if err := validatorSgt.Struct(&request); err != nil {
		return nil, err
	}

	return &request, nil
}

func GetUserId(c fiber.Ctx) (user_domain.UserID, error) {
	plainID := c.Locals("X-userID").(string)
	var userID user_domain.UserID
	err := userID.Scan(plainID)
	if err != nil {
		return user_domain.UserID{}, err
	}

	return userID, nil
}
