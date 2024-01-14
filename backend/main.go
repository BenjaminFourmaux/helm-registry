package main

import (
	"backend/Class/Api"
	"backend/Class/Logger"
)

func main() {
	Logger.Info("Helm Registry - Started")

	// Endpoints registration
	Logger.Info("Registering HTTP Endpoints")
	Api.EndpointTest()
	Api.EndpointIndexYAML()

	// Start HTTP Server
	Logger.Info("Start HTTP Server")
	Api.StartServer()
}
