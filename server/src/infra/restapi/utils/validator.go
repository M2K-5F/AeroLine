package rest_utils

import (
	"aeroline/src/domain/shared"
	"aeroline/src/domain/user_domain"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validatorSgt = validator.New()

func ParseBody[RequestType any](c fiber.Ctx) (*RequestType, error) {
	var request RequestType
	if err := c.Bind().Body(&request); err != nil {
		return nil, &shared.AppError{
			Type: shared.TypeValidation,
			Msg:  err.Error(),
		}
	}

	if err := validatorSgt.Struct(&request); err != nil {
		return nil, &shared.AppError{
			Type: shared.TypeValidation,
			Msg:  err.Error(),
		}
	}

	return &request, nil
}

type scannable interface {
	Scan(value any) error
}

func ParseIDFromQuery[T any, PT interface {
	*T
	scannable
}](c fiber.Ctx, key string) (PT, error) {
	value := c.Query(key)

	if value == "" {
		return nil, fiber.NewError(
			fiber.StatusUnprocessableEntity,
			"query parameter '"+key+"' is required",
		)
	}

	var id PT = new(T)

	if err := id.Scan(value); err != nil {
		return nil, fiber.NewError(
			fiber.StatusUnprocessableEntity,
			"invalid value for '"+key+"': ",
		)
	}

	return id, nil
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
