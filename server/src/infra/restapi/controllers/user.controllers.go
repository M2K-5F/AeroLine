package controllers

import (
	uc "aeroline/src/application/usecases/user"
	reqs "aeroline/src/infra/restapi/dto/requests"
	"aeroline/src/infra/restapi/dto/responses"
	rest "aeroline/src/infra/restapi/utils"

	"github.com/gofiber/fiber/v3"
)

type UserController struct {
	uc *uc.UseCase
}

// @router /users/{id} [get]
// @param id path string true "id"
// @tags users
// @success 200 {object} responses.UserResponse
// @security Bearer
func (ths UserController) GetByID(c fiber.Ctx) error {
	ctx, cancel := rest.BaseRequestContext(c)
	defer cancel()

	cmd, err := rest.ParseCommand[reqs.GetUserByIDRequest](c)
	if err != nil {
		return err
	}

	user, err := ths.uc.GetByID(ctx, *cmd)
	if err != nil {
		return err
	}

	return c.JSON(responses.UserToResponse(*user))
}

func NewUserControllers(useCase *uc.UseCase) *UserController {
	return &UserController{
		uc: useCase,
	}
}
