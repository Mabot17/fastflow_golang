// @title Inventaris API
// @version 1.0
// @description API untuk sistem inventory
// @host localhost:8002
// @BasePath /api
package main

import "inventaris-app/config"

func main() {

	config.LoadEnv()
	config.ConnectDB()

	r := SetupRouter()
	r.Run(":8002")
}