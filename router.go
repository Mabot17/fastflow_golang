package main

import (
	"github.com/gin-gonic/gin"
	"inventaris-app/router"
	_ "inventaris-app/docs"


	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")

	// register modules
	router.RegisterProductRoutes(api)
	router.RegisterStockInRoutes(api)
	router.RegisterStockOutRoutes(api)

	return r
}