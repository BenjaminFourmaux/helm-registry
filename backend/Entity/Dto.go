package Entity

import (
	"time"
)

type ChartDTO struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Version     string    `json:"version"`
	Created     time.Time `json:"created"`
	Digest      string    `json:"digest"`
	Home        string    `json:"home"`
	Sources     string    `json:"sources"`
	Urls        string    `json:"urls"`
}

type RegistryDTO struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Version       int    `json:"version"`
	Maintainer    string `json:"maintainer"`
	MaintainerUrl string `json:"maintainer_url"`
	Labels        string `json:"labels"`
}
