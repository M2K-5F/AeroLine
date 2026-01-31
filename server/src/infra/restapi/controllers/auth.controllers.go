package controllers

import (
	"aeroline/src/application/commands"
	user_usecase "aeroline/src/application/usecases/user"
	"aeroline/src/infra/restapi/responses"
	rest_utils "aeroline/src/infra/restapi/utils"

	"github.com/gofiber/fiber/v3"
)

type AuthController struct {
	useCases *user_usecase.UseCase
}

// @router /auth/register [post]
// @param body body commands.RegisterUserCMD true "body"
// @success 200 {object} responses.UserResponse
func (ths AuthController) Register(c fiber.Ctx) error {
	cmd, err := rest_utils.ParseBody[commands.RegisterUserCMD](c)
	if err != nil {
		return err
	}

	ctx, cancel := rest_utils.BaseRequestContext(c)
	defer cancel()

	user, err := ths.useCases.Register(ctx, *cmd)
	if err != nil {
		return err
	}

	return c.JSON(responses.UserToResponse(*user))
}

// @router /auth/login [post]
// @success 200 {object} responses.UserResponse
// @param body body commands.LoginCMD true "body"
func (ths AuthController) Login(c fiber.Ctx) error {
	cmd, err := rest_utils.ParseBody[commands.LoginCMD](c)
	if err != nil {
		return err
	}

	ctx, cancel := rest_utils.BaseRequestContext(c)
	defer cancel()

	user, err := ths.useCases.Login(ctx, *cmd)
	if err != nil {
		return err
	}

	return c.JSON(responses.UserToResponse(*user))
}

func NewAuthController(useCases *user_usecase.UseCase) *AuthController {
	return &AuthController{
		useCases: useCases,
	}
}
