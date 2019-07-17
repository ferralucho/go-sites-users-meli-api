package miapi

import (
	"../../domains/miapi"
	"../../utils/apierrors"
	"sync"
)

func GetSiteFromApi(siteID string) (*miapi.Site, *apierrors.ApiError) {
	site := &miapi.Site{
		Id: siteID,
	}
	if err := site.Get(); err != nil {
		return nil, err
	}
	return site, nil
}

func GetSiteAsyncFromApi(siteID string, wg *sync.WaitGroup) (*miapi.Site, *apierrors.ApiError) {
	site := &miapi.Site{
		Id: siteID,
	}
	if err := site.Get(); err != nil {
		return nil, err
	}
	wg.Done()

	return site, nil
}


/*
func GetAllSitesFromApi() (*miapi.Sites, *apierrors.ApiError) {
	sites := &miapi.Sites{}
	if err := miapi.GetAllSites(&sites); err != nil {
		return nil, err
	}
	return sites, nil

}
*/

