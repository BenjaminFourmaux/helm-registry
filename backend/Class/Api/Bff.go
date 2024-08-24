package Api

import (
	"backend/Class/Database"
	"backend/Class/Logger"
	"backend/Entity"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func EndpointBFFHome() {
	http.HandleFunc("/bff/home", func(w http.ResponseWriter, req *http.Request) {
		traceRequest(req)

		if req.URL.Path != "/bff/home" {
			Logger.Warning("404 not found")
			http.NotFound(w, req)
			return
		}

		var infoDTO Entity.RegistryDTO

		chartInfo := Database.GetInfo()
		err := chartInfo.Scan(&infoDTO.Name, &infoDTO.Description, &infoDTO.Version, &infoDTO.Maintainer, &infoDTO.MaintainerUrl, &infoDTO.Labels)
		if err != nil {
			Logger.Error("error")
			Logger.Raise(err.Error())
		}

		var query = Database.GetRepositoriesCharts()
		if query.Err() != nil {
			fmt.Println(query.Err())
		}
		count := 0
		for query.Next() {
			count++
		}

		response := Entity.BffHomeResponse{
			Name:          infoDTO.Name.String,
			Description:   infoDTO.Description.String,
			Maintainer:    infoDTO.Maintainer.String,
			MaintainerUrl: infoDTO.MaintainerUrl.String,
			NumberOfRepos: count,
		}

		if infoDTO.Labels.Valid {
			response.Labels = strings.Split(infoDTO.Labels.String, ";")
		} else {
			response.Labels = []string{}
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(response)
	})
}
