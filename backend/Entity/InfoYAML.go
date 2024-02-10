package Entity

type InfoEntity struct {
	Kind       string             `yaml:"kind"`
	ApiVersion string             `yaml:"apiVersion"`
	Registry   InfoRegistryEntity `yaml:"registry"`
}

type InfoRegistryEntity struct {
	Name          string   `yaml:"name"`
	Description   string   `yaml:"description"`
	Version       int      `yaml:"version"`
	Maintainer    string   `yaml:"maintainer"`
	MaintainerUrl string   `yaml:"maintainer_url"`
	Labels        []string `yaml:"labels"`
}
