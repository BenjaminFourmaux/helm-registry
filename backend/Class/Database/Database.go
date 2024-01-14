package Database

import (
	"backend/Class/Logger"
	"database/sql"
)

func OpenConnection(driver string, dataSource string) *sql.DB {
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		Logger.Error("Fail to create/connect to the Database")
	} else {
		Logger.Success("Database connected")
	}

	return db
}

func CreateTableRegistry(db *sql.DB) {
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS registry (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NULL,
			version TEXT NOT NULL,
			created DATETIME NOT NULL,
			digest TEXT NOT NULL,
			home TEXT NULL,
			sources TEXT NULL,
			urls TEXT NULL
		);
	`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		Logger.Error("Fail to create table 'registry'")
		Logger.Raise(err.Error())
	}
}
