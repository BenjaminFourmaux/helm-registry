package Database

import (
	"backend/Entity"
	"database/sql"
	"errors"
	"fmt"
)

// <editor-fold desc="For Table: charts">

func GetAllCharts() (*sql.Rows, error) {
	return DB.Query(`SELECT * FROM charts`)
}

func GetALlChartsOrderedByName() (*sql.Rows, error) {
	return DB.Query(`SELECT * FROM charts GROUP BY name;`)
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
	var id int
	err := result.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return false
	}
	return true
}

// </editor-fold>

// <editor-fold desc="For Table: registry">

func GetInfo() *sql.Row {
	return DB.QueryRow(`SELECT * FROM registry;`)
}

// </editor-fold>
