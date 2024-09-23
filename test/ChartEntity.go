package test

import "time"

type ChartEntry struct {
	APIVersion  string    `yaml:"apiVersion"`
	AppVersion  string    `yaml:"appVersion"`
	Created     time.Time `yaml:"created"`
	Description string    `yaml:"description"`
	Digest      string    `yaml:"digest"`
	Name        string    `yaml:"name"`
	Type        string    `yaml:"type"`
	URLs        []string  `yaml:"urls"`
	Version     string    `yaml:"version"`
}

type Index struct {
	APIVersion string                  `yaml:"apiVersion"`
	Entries    map[string][]ChartEntry `yaml:"entries"`
	Generated  time.Time               `yaml:"generated"`
}
