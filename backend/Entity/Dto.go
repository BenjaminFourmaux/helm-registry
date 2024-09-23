package Entity

import (
	"database/sql"
	"time"
)

type ChartDTO struct {
	Id          int            `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"` // nullable string field
	Version     string         `json:"version"`
	Created     time.Time      `json:"created"`
	Digest      string         `json:"digest"`
	Home        sql.NullString `json:"home"`    // nullable string field
	Sources     sql.NullString `json:"sources"` // nullable string field
	Path        sql.NullString `json:"path"`    // nullable string field
}

type RegistryDTO struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Version       int    `json:"version"`
	Maintainer    string `json:"maintainer"`
	MaintainerUrl string `json:"maintainer_url"`
	Labels        string `json:"labels"`
}
