package miapi

import (
	"../../domains/miapi"
	"../../utils/apierrors"
)

func GetCountryFromApi(countryID string) (*miapi.Country, *apierrors.ApiError) {
	country := &miapi.Country{
		Id: countryID,
	}
	if err := country.Get(); err != nil {
		return nil, err
	}
	return country, nil
}

func GetCountryAsyncFromApi(countryID string, values chan *miapi.Result, errors chan *apierrors.ApiError){
	country := &miapi.Country{
		Id: countryID,
	}
	if err := country.Get(); err != nil {
		errors <- err
	} else{
		errors <- nil
	}
	result := &miapi.Result{
		Country: country,
	}

	values <- result
}