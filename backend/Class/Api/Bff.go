package Api

import (
	"backend/Class/Database"
	"backend/Class/Directory"
	"backend/Class/Logger"
	"backend/Class/Utils/env"
	"backend/Entity"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
)

func EndpointBFFHome() {
	http.HandleFunc("/bff/home", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/bff/home" {
			Logger.Warning("404 not found")
			http.NotFound(w, req)
			return
		}

		var infoDTO Entity.RegistryDTO

		chartInfo := Database.GetInfo()

		chartInfo.Scan(&infoDTO.Name, &infoDTO.Description, &infoDTO.Version, &infoDTO.Maintainer, &infoDTO.MaintainerUrl, &infoDTO.Labels)

		var query, err = Database.GetAllCharts()
		if err != nil {
			fmt.Println(err)
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

func EndpointBFFIcons() {
	http.HandleFunc("/bff/icons", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/bff/icons" {
			Logger.Warning("404 not found")
			http.NotFound(w, req)
			return
		}

		icons, err := Directory.ListFiles(env.ICONS_DIR, []string{".png", ".jpg", ".svg"})
		if err != nil {
			Logger.Error("Failed to list icons")
			http.Error(w, "Failed to list icons", http.StatusInternalServerError)
			return
		}

		var iconList []Entity.IconResponse
		for _, icon := range icons {
			iconName := strings.TrimSuffix(icon, filepath.Ext(icon))
			iconList = append(iconList, Entity.IconResponse{
				Name: iconName,
				Uri:  fmt.Sprintf("%s://%s:%d/icons/%s", env.Scene, env.Hostname, env.Port, icon),
			})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(iconList)
	})
}
