package Api

import (
	"backend/Class/Logger"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func StartServer() {
	port := 8080

	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
	if err != nil {
		Logger.Error("Fail to launch HTTP Server")
		Logger.Raise(err.Error())
	} else {
		Logger.Success("HTTP Server is on listening")
	}

}

func EndpointTest() {
	http.HandleFunc("/test", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, Test !\n")
	})
}

func EndpointIndexYAML() {
	indexFilePath := "index.yaml"

	http.HandleFunc("/index.yaml", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/yaml")

		// Open index.yaml file
		file, err := os.ReadFile(indexFilePath)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error %s", err), http.StatusInternalServerError)
			Logger.Error("Enable to open file")
		}

		// Paste file in the HTTP response
		_, err = io.Copy(w, bytes.NewReader(file))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error %s", err), http.StatusInternalServerError)
		}

		traceRequest(req)
	})
}

func traceRequest(req *http.Request) {
	Logger.Info("Request to '" + req.URL.Path + "' - 200 OK")
}
