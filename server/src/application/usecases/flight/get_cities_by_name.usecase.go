package flight_usecases

import (
	"aeroline/src/domain/shared"
	"context"
	s "strings"
)

type FindCitiesByNameCMD struct {
	SearchQuery string
}

func (ths UseCase) FindCitiesByName(ctx context.Context, cmd FindCitiesByNameCMD) ([]shared.City, error) {
	cities := shared.Filter(shared.CityList, func(el shared.City) bool {
		return s.Contains(s.ToLower(el.Name), s.ToLower(cmd.SearchQuery))
	})

	return cities, nil
}
