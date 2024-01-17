package Entity

import "time"

type ChartEntry struct {
	Version     string    `yaml:"version"`
	Created     time.Time `yaml:"created"`
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	Digest      string    `yaml:"digest"`
	Home        string    `yaml:"home"`
	Sources     []string  `yaml:"sources"`
	Urls        []string  `yaml:"urls"`
}

type Index struct {
	APIVersion string                  `yaml:"apiVersion"`
	Entries    map[string][]ChartEntry `yaml:"entries"`
	Generated  time.Time               `yaml:"generated"`
}
