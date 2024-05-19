package Api

import (
	"backend/Class/Database"
	"backend/Class/Directory"
	"backend/Class/Logger"
	"backend/Class/Utils"
	"backend/Class/Utils/env"
	"backend/Entity"
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func StartServer() {
	port := 8080

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		Logger.Error("Fail to launch HTTP Server")
		Logger.Raise(err.Error())
	} else {
		Logger.Success("HTTP Server listening on port " + strconv.Itoa(port))
	}

}

// <editor-fold desc="Endpoints"> Endpoints

func EndpointTest() {
	http.HandleFunc("/test", func(w http.ResponseWriter, req *http.Request) {
		traceRequest(req)

		var filename = Utils.GetFilenameFromPath("./charts/test-nginx-1.0.0.tgz")
		Logger.Debug(filename)

		chartId := Database.GetChartByFilename(filename)
		var chartToDelete = Utils.ParserRowToChartDTO(chartId)
		fmt.Println(chartToDelete.Id)

		_, err := Database.DeleteChart(chartToDelete.Id)
		if err != nil {
			Logger.Raise(err.Error())
		}

		io.WriteString(w, "Hello, Test !\n")
	})
}

func EndpointHelpRedirect() {
	http.HandleFunc("/help", func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "https://helm.sh/docs/helm/helm_repo/", http.StatusSeeOther)
	})
}

func EndpointRoot() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		traceRequest(req)
		var infoDTO Entity.RegistryDTO

		w.Header().Set("Content-Type", "text/yaml")

		// Get info from Database and convert it into DTO object
		data := Database.GetInfo()
		data.Scan(infoDTO.Name, infoDTO.Description, infoDTO.Version, infoDTO.Maintainer, infoDTO.MaintainerUrl, infoDTO.Labels)

		// DTO to YAML
		infoYAML := Entity.InfoEntity{
			ApiVersion: "v1",
			Kind:       "helm/registry",
			Registry: Entity.InfoRegistryEntity{
				Name:          infoDTO.Name,
				Description:   infoDTO.Description,
				Version:       infoDTO.Version,
				Maintainer:    infoDTO.Maintainer,
				MaintainerUrl: infoDTO.MaintainerUrl,
				Labels:        strings.Split(infoDTO.Labels, ";"),
			},
		}

		yamlData, _ := yaml.Marshal(&infoYAML)

		// Send response
		_, err := io.Copy(w, bytes.NewReader(yamlData))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error %s", err), http.StatusInternalServerError)
		}
	})
}

func EndpointIndexYAML() {
	indexFilePath := env.INDEX_FILE_PATH

	http.HandleFunc("/index.yaml", func(w http.ResponseWriter, req *http.Request) {
		traceRequest(req)

		w.Header().Set("Content-Type", "text/yaml")

		// Open index.yaml file
		file := Directory.ReadFile(indexFilePath)

		// Paste file in the HTTP response
		_, err := io.Copy(w, bytes.NewReader(file))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error %s", err), http.StatusInternalServerError)
		}
	})
}

func EndpointCharts() {
	chartHandler := http.FileServer(http.Dir(env.CHARTS_DIR))

	http.Handle("/charts/", http.StripPrefix("/charts/", chartHandler))
}

func EndpointIcons() {
	iconHandler := http.FileServer(http.Dir(env.ICONS_DIR))

	http.Handle("/icons/", http.StripPrefix("/icons/", iconHandler))
}

// </editor-fold>

func traceRequest(req *http.Request) {
	Logger.Info("HTTP - Request to '" + req.URL.Path + "'")
}
