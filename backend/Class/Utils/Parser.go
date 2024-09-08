package Utils

import (
	"backend/Class/Logger"
	"backend/Entity"
	"database/sql"
	"helm.sh/helm/v3/pkg/chart"
	"strings"
	"time"
)

/*
ParserChartToDTO Parser ChartFile entity to Chart DTO
*/
func ParserChartToDTO(entity *chart.Chart, digest string, path string) Entity.ChartDTO {
	var dto = Entity.ChartDTO{
		Name:        entity.Name(),
		Description: StringToNull(entity.Metadata.Description),
		Version:     entity.Metadata.Version,
		Created:     time.Now(),
		Digest:      digest,
		Home:        StringToNull(entity.Metadata.Home),
		Sources:     StringToNull(strings.Join(entity.Metadata.Sources, ";")),
		Path:        StringToNull(path),
	}
	return dto
}

/*
ParserRowsToChartDTO Parse the result of a DB rows (multiple row result) in a list of ChartDTO
*/
func ParserRowsToChartDTO(rows *sql.Rows) []Entity.ChartDTO {
	var list []Entity.ChartDTO
	for rows.Next() {
		var dto Entity.ChartDTO
		err := rows.Scan(
			&dto.Id,
			&dto.Name,
			&dto.Description,
			&dto.Version,
			&dto.Created,
			&dto.Digest,
			&dto.Home,
			&dto.Sources,
			&dto.Path,
		)

		if err != nil {
			Logger.Error("To parse SQL row in a DTO object")
			Logger.Raise(err)
		}

		list = append(list, dto)
	}

	return list
}

/*
ParserRowToChartDTO Parse the result of a DB row into a ChartDTO
*/
func ParserRowToChartDTO(row *sql.Row) Entity.ChartDTO {
	var dto Entity.ChartDTO
	err := row.Scan(
		&dto.Id,
		&dto.Name,
		&dto.Description,
		&dto.Version,
		&dto.Created,
		&dto.Digest,
		&dto.Home,
		&dto.Sources,
		&dto.Path,
	)

	if err != nil {
		Logger.Error("To parse SQL row in a DTO object")
		Logger.Raise(err)
	}

	return dto
}

/*
NullToString Convert a sql.NullString into a string (same empty)
*/
func NullToString(nullString sql.NullString) string {
	return nullString.String
}

/*
StringToNull Convert a string to a sql.NullString
*/
func StringToNull(str string) sql.NullString {
	if str == "" {
		return sql.NullString{Valid: false}
	} else {
		return sql.NullString{String: str, Valid: true}
	}
}
