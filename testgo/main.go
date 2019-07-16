package main

//http://localhost:8080/miapp/123456

import (
	"./controllers/miapi"
	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

var (
	router = gin.Default()
)

func main () {
	router.GET("/users/:userID", miapi.GetUser)
	router.GET("/sites/:siteID/categories/:categoryID", miapi.GetCategory)
	router.GET("/sites/:siteID", miapi.GetSite)

	router.Run(port)
}