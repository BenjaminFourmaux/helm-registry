package main

import (
	"backend/Class/Api"
	"backend/Class/Database"
	"backend/Class/Directory"
	"backend/Class/Logger"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func main() {
	Logger.Info("Helm Registry - Started")

	// Get env var
	os.Setenv("INDEX_FILE_PATH", "index.yaml")

	// Database
	Database.OpenConnection("sqlite3", "./charts_info.db")
	Database.CreateTableInfo()
	Database.CreateTableRegistry()
	Database.InitInfo(
		os.Getenv("REGISTRY_NAME"),
		os.Getenv("REGISTRY_DESCRIPTION"),
		os.Getenv("REGISTRY_VERSION"),
		os.Getenv("REGISTRY_MAINTAINER"),
		os.Getenv("REGISTRY_MAINTAINER_URL"),
		os.Getenv("REGISTRY_LABELS"),
	)
	//Database.Fixtures() // Insert test fixtures

	// Update file
	Directory.UpdateIndex()

	// Endpoints registration
	Logger.Info("Registering HTTP Endpoints")
	Api.EndpointTest()
	Api.EndpointIndexYAML()

	// Start HTTP Server
	Logger.Info("Start HTTP Server")
	Api.StartServer()
}
