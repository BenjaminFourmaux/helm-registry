package Database

import (
	"database/sql"
	"strconv"
	"strings"
)

func DeleteChart(id int) (sql.Result, error) {
	return DB.Exec(`DELETE FROM charts WHERE id = @id`, id)
}

// DeleteCharts Delete all charts via id passed in list parameter
func DeleteCharts(ids []int) (sql.Result, error) {
	var strList []string
	for _, v := range ids {
		strList = append(strList, strconv.Itoa(v))
	}
	param := strings.Join(strList, ",")

	return DB.Exec(`DELETE FROM charts WHERE id IN (?)`, param)
}
