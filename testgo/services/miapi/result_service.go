package miapi

import (
	"../../domains/miapi"
	"../../utils/apierrors"
	"sync"
)

func GetResultWaitFromApi( userID int64) (*miapi.Result, *apierrors.ApiError) {
	var wg sync.WaitGroup
	var wgErrors sync.WaitGroup
	result := &miapi.Result{}
	retError := &apierrors.ApiError{}

	resultChannel := make(chan *miapi.Result)
	defer close(resultChannel)

	errorsChannel := make(chan *apierrors.ApiError)
	defer close(errorsChannel)
	wgErrors.Add(2)

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

	go func() {
		for error := range errorsChannel {
			if error != nil {
				retError = error
			}
			wgErrors.Done()
		}
	}()

	wg.Wait()

	return result, retError
}