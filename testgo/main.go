package main

import (
	"./controllers/miapi"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

const (
	port = ":8080"
)

var (
	router = gin.Default()
	limiter = rate.NewLimiter(2, 5)
)

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main () {
	router.GET("/users/:userID", miapi.GetUser)
	router.GET("/sites/:siteID/categories/:categoryID", miapi.GetCategory)
	router.GET("/sites/:siteID", miapi.GetSite)
	router.GET("/countries/:countryID", miapi.GetCountry)
	router.GET("/result/:userID", miapi.GetResult)
	router.GET("/sites", miapi.GetAllSites)

	http.ListenAndServe(port, limit(router))
	//router.Run(port)
}