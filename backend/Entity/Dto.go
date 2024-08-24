package Entity

import (
	"database/sql"
	"time"
)

// Represent database tables structure

type ChartDTO struct {
	Id          int            `json:"id"` // Auto generate
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"` // nullable string field
	Version     string         `json:"version"`
	Created     time.Time      `json:"created"`
	Digest      string         `json:"digest"`
	Path        sql.NullString `json:"path"`    // Location in chart directory
	Home        sql.NullString `json:"home"`    // nullable string field
	Sources     sql.NullString `json:"sources"` // nullable string field
}

type RegistryDTO struct {
	Name          sql.NullString `json:"name"`
	Description   sql.NullString `json:"description"`
	Version       sql.NullString `json:"version"`
	Maintainer    sql.NullString `json:"maintainer"`
	MaintainerUrl sql.NullString `json:"maintainer_url"`
	Labels        sql.NullString `json:"labels"`
}
