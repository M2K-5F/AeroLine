package controllers

import (
	uc "aeroline/src/application/usecases/user"
	"aeroline/src/domain/shared"
	auth "aeroline/src/infra/restapi/auth"
	reqs "aeroline/src/infra/restapi/dto/requests"
	resps "aeroline/src/infra/restapi/dto/responses"
	rest "aeroline/src/infra/restapi/utils"
	"time"

	"github.com/gofiber/fiber/v3"
)

type AuthController struct {
	useCases    *uc.UseCase
	authService *auth.AuthService
}

// @router /auth/register [post]
// @tags auth
// @param body body reqs.RegisterUserRequest true "body"
// @success 200 {object} resps.UserResponse
func (ths AuthController) Register(c fiber.Ctx) error {
	ctx, cancel := rest.BaseRequestContext(c)
	defer cancel()

	cmd, err := rest.ParseCommand[reqs.RegisterUserRequest](c)
	if err != nil {
		return err
	}

	user, err := ths.useCases.Register(ctx, *cmd)
	if err != nil {
		return err
	}

	return c.JSON(resps.UserToResponse(*user))
}

// @router /auth/login [post]
// @tags auth
// @success 200 {object} responses.LoginUserResponse
// @param body body uc.LoginCMD true "body"
func (ths AuthController) Login(c fiber.Ctx) error {
	cmd, err := rest.ParseCommand[reqs.LoginUserRequest](c)
	if err != nil {
		return err
	}

	ctx, cancel := rest.BaseRequestContext(c)
	defer cancel()

	user, err := ths.useCases.Login(ctx, *cmd)
	if err != nil {
		return err
	}

	device := auth.GetDevice(c)

	access, refresh, err := ths.authService.Login(ctx, user, device)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Expires:  time.Now().Add(24 * 7 * time.Hour),
		Value:    string(*refresh),
		Name:     "refr",
		HTTPOnly: true,
		SameSite: "Lax",
		Path:     "/",
	})

	return c.JSON(resps.LoginUserResponse{
		UserResponse: resps.UserToResponse(*user),
		Access:       string(*access),
	})
}

// @router /auth/me [get]
// @tags auth
// @success 200 {object} resps.UserResponse
// @Security Bearer
func (ths AuthController) Me(c fiber.Ctx) error {
	ctx, cancel := rest.BaseRequestContext(c)
	defer cancel()

	cmd := new(uc.GetByIdCMD)

	cmd.UserID = *auth.GetUserID(c)

	user, err := ths.useCases.GetByID(ctx, *cmd)
	if err != nil {
		return err
	}

	return c.JSON(resps.UserToResponse(*user))
}

// @router /auth/refresh [patch]
// @tags auth
// @success 200 {object} map[string]string
// @security Bearer
func (ths AuthController) Refresh(c fiber.Ctx) error {
	ctx, cancel := rest.BaseRequestContext(c)
	defer cancel()

	device := auth.GetDevice(c)

	refresh, err := auth.GetRefreshToken(c)
	if err != nil {
		return err
	}

	access, refresh, err := ths.authService.RefreshToken(ctx, *refresh, device)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Expires:  time.Now().Add(24 * 7 * time.Hour),
		Value:    string(*refresh),
		Name:     "refr",
		HTTPOnly: true,
		SameSite: "Lax",
		Path:     "/",
	})

	return c.JSON(fiber.Map{"access": string(*access)})
}

// @router /auth/sessions [get]
// @tags auth
// @success 200 {array} rest_auth.SessionResponse
// @security Bearer
func (ths AuthController) GetSessions(c fiber.Ctx) error {
	userID := auth.GetUserID(c)

	ctx, cancel := rest.BaseRequestContext(c)
	defer cancel()

	sessions, err := ths.authService.GetUserSessions(ctx, *userID)
	if err != nil {
		return err
	}

	return c.JSON(shared.Map(sessions, auth.SessionToResponse))
}

type LogoutResponse struct {
	Message string `json:"message" example:"Successful logout"`
}

// @router /auth/logout [post]
// @tags auth
// @success 200 {object} LogoutResponse
// @security Bearer
func (ths AuthController) Logout(c fiber.Ctx) error {
	_ = auth.GetUserID(c)

	c.Cookie(&fiber.Cookie{
		Expires:  time.Now(),
		Value:    "",
		Name:     "refr",
		HTTPOnly: true,
		SameSite: "Lax",
		Path:     "/",
	})

	return c.JSON(LogoutResponse{Message: "Successful logout"})
}

func NewAuthController(useCases *uc.UseCase, authService *auth.AuthService) *AuthController {
	return &AuthController{
		useCases:    useCases,
		authService: authService,
	}
}
