package requests

import uc "aeroline/src/application/usecases/flight"

type FindCitiesByNameRequest struct {
	SearchQuery string `query:"q"`
}

func (ths FindCitiesByNameRequest) ToCMD() (*uc.FindCitiesByNameCMD, error) {
	return &uc.FindCitiesByNameCMD{
		SearchQuery: ths.SearchQuery,
	}, nil
}
