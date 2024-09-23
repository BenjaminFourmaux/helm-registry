package Database

import (
	"database/sql"
)

func DeleteChart(id int) (sql.Result, error) {
	return DB.Exec(`DELETE FROM charts WHERE id = @id`, id)
}

/*
DeleteCharts Delete all charts via id passed in list parameter
*/
func DeleteCharts(ids []int) error {
	var err error

	// This method doesn't work
	/*var strList []string
	for _, v := range ids {
		strList = append(strList, strconv.Itoa(v))
	}
	param := strings.Join(strList, ", ")

	return DB.Exec(`DELETE FROM charts WHERE id IN ($1)`, param)*/

	for _, id := range ids {
		_, errTmp := DB.Exec(`DELETE FROM charts WHERE id IN ($1)`, id)

		if errTmp != nil {
			err = errTmp
		}
	}

	return err
}
