package main

import (
	"backend/Class/Api"
	"backend/Class/Database"
	"backend/Class/Directory"
	"backend/Class/Logger"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var db *sql.DB

func main() {
	Logger.Info("Helm Registry - Started")

	// Get env var
	os.Setenv("INDEX_FILE_PATH", "index.yaml")

	// Database
	Database.OpenConnection("sqlite3", "./charts_info.db")
	Database.CreateTableRegistry()
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
