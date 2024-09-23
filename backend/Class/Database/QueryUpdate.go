package Database

import "backend/Entity"

func UpdateChart(id int, chart Entity.ChartDTO) {
	DB.Exec(`
		UPDATE charts SET 
	    description = @description, 
	    digest = @digest,
	    created = @created,
	    home = @home,
	    sources = @sources,
	    path = @path
		WHERE id = @id
	`, chart.Description, chart.Digest, chart.Created, chart.Home, chart.Sources, chart.Path, id)
}
