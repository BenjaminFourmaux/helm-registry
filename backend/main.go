package main

import (
	"backend/Class/Api"
	"backend/Class/Database"
	"backend/Class/Logger"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	Logger.Info("Helm Registry - Started")

	// Database
	db := Database.OpenConnection("sqlite3", "./charts_info.db")
	Database.CreateTableRegistry(db)

	// Endpoints registration
	Logger.Info("Registering HTTP Endpoints")
	Api.EndpointTest()
	Api.EndpointIndexYAML()

	// Start HTTP Server
	Logger.Info("Start HTTP Server")
	Api.StartServer()
}
