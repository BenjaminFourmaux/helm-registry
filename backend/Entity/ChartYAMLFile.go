package Entity

/*
ChartFile schema : https://helm.sh/docs/topics/charts/
*/
type ChartFile struct {
	APIVersion   string            `yaml:"apiVersion"`             // The chart API version (required)
	Name         string            `yaml:"name"`                   // The name of the chart (required)
	Version      string            `yaml:"version"`                // A SemVer 2 version (required)
	KubeVersion  string            `yaml:"kubeVersion,omitempty"`  // A SemVer range of compatible Kubernetes versions (optional)
	Description  string            `yaml:"description,omitempty"`  // A single-sentence description of this project (optional)
	Type         string            `yaml:"type,omitempty"`         // The type of the chart (optional)
	Keywords     []string          `yaml:"keywords,omitempty"`     // A list of keywords about this project (optional)
	Home         string            `yaml:"home,omitempty"`         // The URL of this project home page (optional)
	Sources      []string          `yaml:"sources,omitempty"`      // A list of URLs to source code for this project (optional)
	Dependencies []ChartDependency `yaml:"dependencies,omitempty"` // A list of the chart requirements (optional)
	Maintainers  []ChartMaintainer `yaml:"maintainers,omitempty"`  // A list of maintainers (optional)
	Icon         string            `yaml:"icon,omitempty"`         // A URL to an SVG or PNG image to be used as an icon (optional).
	AppVersion   string            `yaml:"appVersion,omitempty"`   // The version of the app that this contains (optional). Needn't be SemVer. Quotes recommended.
	Deprecated   bool              `yaml:"deprecated,omitempty"`   // Whether this chart is deprecated (optional, boolean)
	Annotations  map[string]string `yaml:"annotations,omitempty"`  // A list of annotations keyed by name (optional)
}

type ChartDependency struct {
	Name         string   `yaml:"name"`
	Version      string   `yaml:"version"`
	Repository   string   `yaml:"repository,omitempty"`
	Condition    string   `yaml:"condition,omitempty"`
	Tags         []string `yaml:"tags,omitempty"`
	ImportValues []string `yaml:"import-values,omitempty"`
	Alias        string   `yaml:"alias,omitempty"`
}

type ChartMaintainer struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email,omitempty"`
	URL   string `yaml:"url,omitempty"`
}
