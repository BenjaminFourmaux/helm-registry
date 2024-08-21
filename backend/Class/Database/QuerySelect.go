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

// GetRepositoriesCharts Get distinct charts (by name) with last version
func GetRepositoriesCharts() *sql.Rows {
	var queryResult, _ = DB.Query(`
		SELECT *
		FROM charts c1
		WHERE c1.version = (
			SELECT MAX(c2.version)
			FROM charts c2
			WHERE c2.name = c1.name
		);
	`)

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
	return DB.QueryRow(`SELECT *  FROM registry LIMIT 1;`)
}

// </editor-fold>
