package Database

import (
	"backend/Class/Logger"
	"backend/Class/Utils"
	. "backend/Entity"
)

// InsertChart Insert a new chart on the table 'charts'
func InsertChart(entity ChartDTO) {
	var sql = `
		INSERT INTO charts (
			name, description, version, created, digest, path, home, sources      
		)
		VALUES (
		    $1, $2, $3, $4, $5, $6, $7, $8
		);
	`

	_, err := DB.Exec(sql, entity.Name, Utils.NullToString(entity.Description), entity.Version, entity.Created, entity.Digest, Utils.NullToString(entity.Path), Utils.NullToString(entity.Home), Utils.NullToString(entity.Sources))
	if err != nil {
		Logger.Error("Unable to insert in table 'charts'")
		Logger.Raise(err.Error())
	}
}
