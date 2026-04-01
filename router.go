package main

import (
	"github.com/gin-gonic/gin"
	"inventaris-app/router"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	// register modules
	router.RegisterProductRoutes(api)

	return r
}