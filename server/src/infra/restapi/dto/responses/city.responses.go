package responses

import "aeroline/src/domain/shared"

type CityResponse struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func CityToResponse(city shared.City) CityResponse {
	return CityResponse{
		Code:    city.Code,
		Name:    city.Name,
		Country: string(city.Country),
	}
}
