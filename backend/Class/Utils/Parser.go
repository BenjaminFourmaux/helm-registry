package Utils

import (
	"backend/Class/Logger"
	"backend/Entity"
	"database/sql"
	"strings"
	"time"
)

/*
ParserChartToDTO Parser ChartFile entity to Chart DTO
*/
func ParserChartToDTO(entity Entity.ChartFile, path string) Entity.ChartDTO {
	var dto = Entity.ChartDTO{
		Name:        entity.Name,
		Description: StringToNull(entity.Description),
		Version:     entity.Version,
		Created:     time.Now(),
		Digest:      "", // TODO : Compute manually the hash via sha-256 algorithm
		Home:        StringToNull(entity.Home),
		Sources:     StringToNull(strings.Join(entity.Sources, ";")),
		Urls:        StringToNull(strings.Join(entity.Urls, ";")),
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
			&dto.Path,
			&dto.Home,
			&dto.Sources,
			&dto.Urls,
		)

		if err != nil {
			Logger.Error("To parse SQL row in a DTO object")
			Logger.Raise(err.Error())
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
		&dto.Path,
		&dto.Home,
		&dto.Sources,
		&dto.Urls,
	)

	if err != nil {
		Logger.Error("To parse SQL row in a DTO object")
		Logger.Raise(err.Error())
	}

	return dto
}

/*
NullToString Convert a sql.NullString into a string (same empty)
*/
func NullToString(nullString sql.NullString) string {
	if nullString.Valid {
		return nullString.String
	} else {
		return ""
	}
}

/*
StringToNull Convert a string to a sql.NullString
*/
func StringToNull(str string) sql.NullString {
	if str == "" {
		return sql.NullString{}
	} else {
		return sql.NullString{String: str}
	}
}
