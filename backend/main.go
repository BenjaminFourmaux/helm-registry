package main

import (
	"backend/Class/Api"
	"backend/Class/Database"
	"backend/Class/Directory"
	"backend/Class/Logger"
	"backend/Class/Utils/env"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	Logger.Info("Helm Registry - Started")

	// Setup env
	env.SetupEnv()

	// Database
	Database.OpenConnection("sqlite3", "./charts_info.db")
	Database.CreateTableInfo()
	Database.CreateTableRegistry()
	Database.InitInfo(
		env.REGISTRY_NAME,
		env.REGISTRY_DESCRIPTION,
		env.REGISTRY_VERSION,
		env.REGISTRY_MAINTAINER,
		env.REGISTRY_MAINTAINER_URL,
		env.REGISTRY_LABELS,
	)
	//Database.Fixtures() // Insert test fixtures

	// Update file
	Directory.UpdateIndex()

	// Endpoints registration
	Logger.Info("Registering HTTP Endpoints")
	Api.EndpointRoot()
	Api.EndpointTest()
	Api.EndpointIndexYAML()
	Api.EndpointCharts()

	// Start HTTP Server
	Logger.Info("Start HTTP Server")
	Api.StartServer()
}
