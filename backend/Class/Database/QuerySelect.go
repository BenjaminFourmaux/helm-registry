package Database

import (
	"database/sql"
	"fmt"
)

func GetALlChartsOrderedByName() (*sql.Rows, error) {
	return DB.Query(`SELECT * FROM charts GROUP BY name;`)
}

func GetInfo() *sql.Row {
	return DB.QueryRow(`SELECT * FROM registry;`)
}

// GetChartByFilename Retrieve a Chart by his filename in URLs
func GetChartByFilename(filename string) *sql.Row {
	return DB.QueryRow(fmt.Sprintf("SELECT * FROM charts WHERE urls LIKE '%%%s%%'", filename))
}
