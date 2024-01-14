package Database

import (
	"backend/Class/Logger"
	"database/sql"
)

var db *sql.DB

func OpenConnection(driver string, dataSource string) {
	conn, err := sql.Open(driver, dataSource)
	if err != nil {
		Logger.Error("Fail to create/connect to the Database")
	} else {
		Logger.Success("Database connected")
	}

	db = conn
}

func CreateTableRegistry() {
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

func Fixtures() {
	Logger.Info("Insert fixtures data")

	insertFixturesSQL := `
		INSERT INTO registry (
			id, name, description, version, created, digest, home, sources, urls                 
		) 
		VALUES (
			1, 'test', 'Deploy a basic test pod' , '0.1.0', 
			'2016-10-06T16:23:20.499543808-06:00', '515c58e5f79d8b2913a10cb400ebb6fa9c77fe713287afbacf1a0b897cd78727',
		    'https://helm.sh/helm', 'https://github.com/helm/helm', 'charts/test/test-0.1.0.tgz' 
		);
		INSERT INTO registry (
			id, name, description, version, created, digest, home, sources, urls                 
		) 
		VALUES (
			2, 'test', 'Deploy a basic test pod' , '0.2.0', 
			'2016-10-06T16:23:20.499814565-06:00', '99c76e403d752c84ead610644c4b1c2f2b453a74b921f422b9dcb8a7c8b559cd',
		    'https://helm.sh/helm', 'https://github.com/helm/helm', 'charts/test/test-0.2.0.tgz' 
		);
		INSERT INTO registry (
			id, name, description, version, created, digest, home, sources, urls                 
		) 
		VALUES (
			3, 'toto', 'Deploy a basic toto pod' , '1.1.0', 
			'2016-10-06T16:23:20.499543808-06:00', '515c58e5f79d8b2913a10cb400ebb6fa9c77fe813289afbacf1a0b897cd78727',
		    'https://helm.sh/helm', 'https://github.com/helm/helm', 'charts/toto/toto-1.1.0.tgz' 
		);
	`

	_, err := db.Exec(insertFixturesSQL)
	if err != nil {
		Logger.Warning("Fail to insert fixtures")
		Logger.Raise(err.Error())
	}
}
