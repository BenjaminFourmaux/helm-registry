package Database

import (
	"backend/Class/Logger"
	. "backend/Entity"
)

// InsertChart Insert a new chart on the table 'charts'
func InsertChart(entity ChartDTO) {
	var sql = `
		INSERT INTO charts (
			name, description, version, created, digest, home, sources, urls      
		)
		VALUES (
		    $1, $2, $3, $4, $5, $6, $7, $8
		);
	`

	_, err := DB.Exec(sql, entity.Name, entity.Description, entity.Version, entity.Created, entity.Digest, entity.Home, entity.Sources, entity.Urls)
	if err != nil {
		Logger.Error("Unable to insert in table 'charts'")
		Logger.Raise(err.Error())
	}
}
