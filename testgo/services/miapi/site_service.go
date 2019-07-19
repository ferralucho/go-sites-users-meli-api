package miapi

import (
	"../../domains/miapi"
	"../../utils/apierrors"
)

func GetSiteFromApi(siteID string) (*miapi.Site, *apierrors.ApiError) {
	site := &miapi.Site{
		ID: siteID,
	}
	if err := site.Get(); err != nil {
		return nil, err
	}
	return site, nil
}

func GetSiteAsyncFromApi(siteID string, values chan *miapi.Result, errors chan *apierrors.ApiError) {
	site := &miapi.Site{
		ID: siteID,
	}
	if err := site.Get(); err != nil {
		errors <- err
	} else{
		errors <- nil
	}
	result := &miapi.Result{
		Site: site,
	}

	values <- result
}

func GetAllSitesFromApi() (*miapi.Sites, *apierrors.ApiError) {
	sites := &miapi.Sites{}
	if err := miapi.GetAllSites(sites); err != nil {
		return nil, err
	}
	return sites, nil

}


