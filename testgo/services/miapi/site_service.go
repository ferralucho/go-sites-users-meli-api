package miapi

import (
	"../../domains/miapi"
	"../../utils/apierrors"
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
