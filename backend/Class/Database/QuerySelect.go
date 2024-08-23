package Database

import (
	"backend/Class/Logger"
	"backend/Entity"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
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
	// Maybe calculate the digest by sha256 the Chart.yaml of a chart and compare with chart's digest in db
	var queryResult = DB.QueryRow(`
		SELECT id FROM charts
		WHERE name = ?
		AND version = ?
		AND path = ?
	`, chart.Name, chart.Version, chart.Path)

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
	Logger.Error(err.Error())
	Logger.Debug(strconv.Itoa(id))
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
