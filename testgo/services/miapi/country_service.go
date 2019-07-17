package miapi

import (
	"../../domains/miapi"
	"../../utils/apierrors"
	"sync"
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

func GetCountryAsyncFromApi(countryID string, wg *sync.WaitGroup) (*miapi.Country, *apierrors.ApiError) {
	country := &miapi.Country{
		Id: countryID,
	}
	if err := country.Get(); err != nil {
		return nil, err
	}
	wg.Done()
	return country, nil
}