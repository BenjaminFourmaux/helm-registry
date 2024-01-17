package Entity

import "time"

type DTORegistry struct {
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
