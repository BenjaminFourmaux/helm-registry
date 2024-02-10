package Database

import "database/sql"

func GetALlChartsOrderedByName() (*sql.Rows, error) {
	return DB.Query(`SELECT * FROM registry GROUP BY name;`)
}

func GetInfo() *sql.Row {
	return DB.QueryRow(`SELECT * FROM info;`)
}
