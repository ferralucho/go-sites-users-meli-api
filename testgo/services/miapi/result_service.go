package miapi

import (
	"../../domains/miapi"
	"../../utils/apierrors"
	"sync"
)

func GetResultFromApi(userID int64) (*miapi.Result, *apierrors.ApiError) {
	var wg sync.WaitGroup

	wg.Add(1)
	user, err := GetUserAsyncFromApi(userID, &wg)

	if err != nil {
		return nil,err
	}

	wg.Wait()

	wg.Add(1)
	country, err := GetCountryAsyncFromApi(user.CountryID, &wg)

	if err != nil {
		return nil,err
	}

	wg.Wait()

	wg.Add(1)
	site, err := GetSiteAsyncFromApi(user.SiteID, &wg)

	if err != nil {
		return nil,err
	}

	wg.Wait()

	result := &miapi.Result{
		User: *user,
		Country: *country,
		Site: *site,
	}

	return result, nil
}