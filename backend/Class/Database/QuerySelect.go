package Database

import "database/sql"

func GetALlChartsOrderedByName() (*sql.Rows, error) {
	return DB.Query(`SELECT * FROM charts GROUP BY name;`)
}

func GetInfo() *sql.Row {
	return DB.QueryRow(`SELECT * FROM info;`)
}
