package Database

import (
	"backend/Entity"
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

func GetChartByCriteria(chart Entity.ChartDTO) *sql.Row {
	var queryResult = DB.QueryRow(`
		SELECT * FROM charts
		WHERE name = @name
		AND version = @version
		AND urls LIKE '%%%@urls%%%'
	`, chart.Name, chart.Version, chart.Urls)

	return queryResult
}

func IfChartExist(chart Entity.ChartDTO) bool {
	var result = GetChartByCriteria(chart)
	return result != nil
}
