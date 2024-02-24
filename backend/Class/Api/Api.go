package Api

import (
	"backend/Class/Database"
	"backend/Class/Directory"
	"backend/Class/Logger"
	"backend/Entity"
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"net/http"
	"os"
	"strings"
)

func StartServer() {
	port := 8080

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		Logger.Error("Fail to launch HTTP Server")
		Logger.Raise(err.Error())
	} else {
		Logger.Success("HTTP Server is on listening")
	}

}

// <editor-fold desc="Endpoints"> Endpoints

func EndpointRoot() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		traceRequest(req)
		var infoDTO Entity.DTORegistry

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

func EndpointTest() {
	http.HandleFunc("/test", func(w http.ResponseWriter, req *http.Request) {
		traceRequest(req)

		io.WriteString(w, "Hello, Test !\n")
	})
}

func EndpointIndexYAML() {
	indexFilePath := os.Getenv("INDEX_FILE_PATH")

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
	chartDir := os.Getenv("REPOSITORY_DIR")
	chartHandler := http.FileServer(http.Dir(chartDir))
	http.Handle("/charts/", http.StripPrefix("/charts/", chartHandler))
}

// </editor-fold>

func traceRequest(req *http.Request) {
	Logger.Info("Request to '" + req.URL.Path + "'")
}
