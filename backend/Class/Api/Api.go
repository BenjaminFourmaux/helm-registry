package Api

import (
	"backend/Class/Database"
	"backend/Class/Directory"
	"backend/Class/Logger"
	"backend/Class/Utils/env"
	"backend/Entity"
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"net/http"
	"strings"
)

func StartServer() {
	Logger.Success(fmt.Sprintf("HTTP Server is on listening on port: %d", env.Port))

	err := http.ListenAndServe(fmt.Sprintf(":%d", env.Port), nil)
	if err != nil {
		Logger.Error("Fail to launch HTTP Server")
		Logger.Raise(err)
	} else {
		Logger.Success("HTTP Server is on listening")
	}

}

// <editor-fold desc="Endpoints"> Endpoints

func EndpointTest() {
	http.HandleFunc("/test", func(w http.ResponseWriter, req *http.Request) {
		traceRequest(req)

		if req.URL.Path != "/test" {
			Logger.Warning("404 not found")
			http.NotFound(w, req)
			return
		}

		io.WriteString(w, "Hello, Test !\n")
	})
}

func EndpointHelpRedirect() {
	http.HandleFunc("/help", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/help" {
			Logger.Warning("404 not found")
			http.NotFound(w, req)
			return
		}

		http.Redirect(w, req, "https://helm.sh/docs/helm/helm_repo/", http.StatusSeeOther)
	})
}

func EndpointRoot() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		traceRequest(req)

		if req.URL.Path != "/" && req.URL.Path != "/favicon.ico" {
			Logger.Warning("404 not found")
			http.NotFound(w, req)
			return
		}

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
				Name:          infoDTO.Name.String,
				Description:   infoDTO.Description.String,
				Version:       infoDTO.Version.String,
				Maintainer:    infoDTO.Maintainer.String,
				MaintainerUrl: infoDTO.MaintainerUrl.String,
				Labels:        strings.Split(infoDTO.Labels.String, ";"),
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

		if req.URL.Path != "/index.yaml" {
			Logger.Warning("404 not found")
			http.NotFound(w, req)
			return
		}

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
	chartDir := env.REPOSITORY_DIR
	chartHandler := http.FileServer(http.Dir(chartDir))

	http.HandleFunc("/charts/", func(w http.ResponseWriter, req *http.Request) {
		traceRequest(req)

		if !strings.Contains(req.URL.Path, "/charts/") {
			Logger.Warning("404 not found")
			http.NotFound(w, req)
			return
		}

		http.StripPrefix("/charts/", chartHandler)
	})
}

// </editor-fold>

func traceRequest(req *http.Request) {
	Logger.Info("Request to '" + req.URL.Path + "'")
}
