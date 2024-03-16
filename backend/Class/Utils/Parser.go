package Utils

import (
	"backend/Class/Logger"
	"backend/Entity"
	"database/sql"
	"strings"
	"time"
)

// ParserChartToDTO Parser ChartFile entity to Chart DTO
func ParserChartToDTO(entity Entity.ChartFile, urls []string) Entity.ChartDTO {
	var dto = Entity.ChartDTO{
		Name:        entity.Name,
		Description: entity.Description,
		Version:     entity.Version,
		Created:     time.Now(),
		Digest:      "", // TODO : Compute manually the hash via sha-256 algorithm
		Home:        entity.Home,
		Sources:     strings.Join(entity.Sources, ";"),
		Urls:        strings.Join(urls, ";"),
	}
	return dto
}

// ParserRowToChartDTO Parse the result of an DB row into a ChartDTO
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
		&dto.Urls,
	)
	if err != nil {
		Logger.Error("To parse SQL row in a DTO object")
		Logger.Raise(err.Error())
	}

	return dto
}
