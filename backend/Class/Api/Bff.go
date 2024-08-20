package Api

import (
	"backend/Class/Database"
	"backend/Entity"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func EndpointBFFHome() {
	http.HandleFunc("/bff/home", func(w http.ResponseWriter, req *http.Request) {
		traceRequest(req)

		var infoDTO Entity.RegistryDTO

		chartInfo := Database.GetInfo()
		chartInfo.Scan(infoDTO.Name, infoDTO.Description, infoDTO.Version, infoDTO.Maintainer, infoDTO.MaintainerUrl, infoDTO.Labels)

		var query = Database.GetRepositoriesCharts()
		if query.Err() != nil {
			fmt.Println(query.Err())
		}
		count := 0
		for query.Next() {
			count++
		}

		response := Entity.BffHomeResponse{
			Name:          infoDTO.Name,
			Description:   infoDTO.Description,
			Maintainer:    infoDTO.Maintainer,
			MaintainerUrl: infoDTO.MaintainerUrl,
			Labels:        strings.Split(infoDTO.Labels, ";"),
			NumberOfRepos: count,
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(response)
	})
}
