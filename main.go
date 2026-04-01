package main

import "inventaris-app/config"

func main() {

	config.LoadEnv()
	config.ConnectDB()

	r := SetupRouter()
	r.Run(":8002")
}