package miapi

import (
	"../../domains/miapi"
	"../../utils/apierrors"
	"sync"
)

func GetResultWaitFromApi( userID int64) (*miapi.Result, *apierrors.ApiError) {
	var wg sync.WaitGroup
	result := &miapi.Result{}

	resultChannel := make(chan *miapi.Result)
	defer close(resultChannel)

	errorsChannel := make(chan apierrors.ApiError)
	defer close(errorsChannel)

	user, err := GetUserFromApi(userID)
	result.User = user

	if err != nil {
		return nil,err
	}

	wg.Add(2)

	go GetCountryAsyncFromApi(user.CountryID, resultChannel, errorsChannel)
	go GetSiteAsyncFromApi(user.SiteID, resultChannel, errorsChannel)

	go func() {
		for res := range resultChannel {
			if res.Country != nil {
				result.Country = res.Country
			}
			if res.Site != nil {
				result.Site = res.Site
			}
			wg.Done()
		}
	}()

	wg.Wait()

	return result, nil
}
/*
func GetResultChannelFromApi(userID int64) (*miapi.Result, *apierrors.ApiError) {
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

 */