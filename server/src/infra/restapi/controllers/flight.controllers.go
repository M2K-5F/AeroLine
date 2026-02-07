package controllers

import (
	uc "aeroline/src/application/usecases/flight"
	"aeroline/src/domain/shared"
	"aeroline/src/infra/restapi/dto/requests"
	"aeroline/src/infra/restapi/dto/responses"
	utils "aeroline/src/infra/restapi/utils"

	"github.com/gofiber/fiber/v3"
)

type FlightController struct {
	uc *uc.UseCase
}

// @router /flight/cities [get]
// @param q query string true "name of city"
// @success 200 {array} responses.CityResponse
func (ths FlightController) FindCitiesByName(c fiber.Ctx) error {
	ctx, cancel := utils.BaseRequestContext(c)
	defer cancel()

	cmd, err := utils.ParseCommand[requests.FindCitiesByNameRequest](c)
	if err != nil {
		return err
	}

	cities, err := ths.uc.FindCitiesByName(ctx, *cmd)
	if err != nil {
		return err
	}

	return c.JSON(shared.Map(cities, func(c shared.City) responses.CityResponse {
		return responses.CityToResponse(c)
	}))
}

func NewFlightController(uc uc.UseCase) *FlightController {
	return &FlightController{
		uc: &uc,
	}
}
