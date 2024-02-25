package Utils

import (
	"backend/Entity"
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
