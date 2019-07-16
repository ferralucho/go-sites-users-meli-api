package main

//http://localhost:8080/miapp/123456

import (
	"./controllers/miapp"
	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

var (
	router = gin.Default()
)

func main () {
	router.GET("/miapp/:userID", miapp.GetUser)
	router.Run(port)
}