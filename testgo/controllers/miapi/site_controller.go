package miapi

import (
	"../../services/miapi"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	paramSiteID = "siteID"
)

func GetSite(c *gin.Context){
	siteID := c.Param(paramSiteID)

	site, apiError := miapi.GetSiteFromApi(siteID)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, site)
}

/*
func GetAllSites(c *gin.Context){

	sites, apiError := miapi.GetAllSitesFromApi()
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, sites)
}*/