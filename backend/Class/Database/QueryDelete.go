package Database

import "database/sql"

func DeleteChart(id int) (sql.Result, error) {
	return DB.Exec(`DELETE FROM charts WHERE id = @id`, map[string]interface{}{"id": id})
}
