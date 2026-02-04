package controllers

import (
	"aeroline/src/application/commands"
	user_usecase "aeroline/src/application/usecases/user"
	"aeroline/src/domain/shared"
	rest_auth "aeroline/src/infra/restapi/auth"
	"aeroline/src/infra/restapi/responses"
	rest_utils "aeroline/src/infra/restapi/utils"
	"time"

	"github.com/gofiber/fiber/v3"
)

type AuthController struct {
	useCases    *user_usecase.UseCase
	authService *rest_auth.AuthService
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
// @success 200 {object} responses.LoginUserResponse
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

	devid := rest_auth.GetDeviceID(c)

	access, refresh, err := ths.authService.Login(ctx, user, *devid)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Expires:  time.Now().Add(24 * 7 * time.Hour),
		Value:    string(refresh),
		Name:     "refr",
		HTTPOnly: true,
		SameSite: "Lax",
		Path:     "/",
	})

	return c.JSON(responses.LoginUserResponse{
		UserResponse: responses.UserToResponse(*user),
		Access:       string(access),
	})
}

// @router /auth/me [get]
// @success 200 {object} string
// @Security Bearer
func (ths AuthController) Me(c fiber.Ctx) error {
	userId := rest_auth.GetUserID(c)

	return c.SendString(userId.String())
}

// @router /auth/refresh [patch]
// @success 200 {object} map[string]string
// @security Bearer
func (ths AuthController) Refresh(c fiber.Ctx) error {
	ctx, cancel := rest_utils.BaseRequestContext(c)
	defer cancel()

	deviceID := rest_auth.GetDeviceID(c)

	token := c.Cookies("refr")
	if token == "" {
		return rest_auth.ErrTokenNotFound
	}

	access, refresh, err := ths.authService.RefreshToken(ctx, rest_auth.RefreshToken(token), *deviceID)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Expires:  time.Now().Add(24 * 7 * time.Hour),
		Value:    string(refresh),
		Name:     "refr",
		HTTPOnly: true,
		SameSite: "Lax",
		Path:     "/",
	})

	return c.JSON(fiber.Map{"access": string(access)})
}

// @router /auth/sessions [get]
// @success 200 {object} []SessionRequest
// @security Bearer
func (ths AuthController) GetSessions(c fiber.Ctx) error {
	userID := rest_auth.GetUserID(c)

	ctx, cancel := rest_utils.BaseRequestContext(c)
	defer cancel()

	sessions, err := ths.authService.GetUserSessions(ctx, *userID)
	if err != nil {
		return err
	}

	return c.JSON(shared.Map(sessions, func(s rest_auth.Session) SessionRequest {
		return SessionRequest{
			UserID:       s.UserID.String(),
			DeviceID:     s.DeviceID.String(),
			SessionID:    s.SessionID.String(),
			LastActivity: s.LastActivity,
		}
	}))
}

type SessionRequest struct {
	SessionID    string
	UserID       string
	LastActivity time.Time
	DeviceID     string
}

func NewAuthController(useCases *user_usecase.UseCase, authService *rest_auth.AuthService) *AuthController {
	return &AuthController{
		useCases:    useCases,
		authService: authService,
	}
}
