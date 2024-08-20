package Entity

type BffHomeResponse struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Maintainer    string   `json:"maintainer"`
	MaintainerUrl string   `json:"maintainer_url"`
	Labels        []string `json:"labels"`
	NumberOfRepos int      `json:"number_of_repos"`
}
