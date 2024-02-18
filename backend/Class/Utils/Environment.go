package Utils

import (
	"backend/Class/Logger"
	"os"
)

func SetupEnv() {
	// Get the running env (Docker or not)
	isDocker := IsDocker()

	// Get Env vars and if not declared, init it with default value

	if os.Getenv("INDEX_FILE_PATH") == "" {
		_ = os.Setenv("INDEX_FILE_PATH", "index.yaml")
	}

	if os.Getenv("REPOSITORY_DIR") == "" {
		if isDocker {
			_ = os.Setenv("REPOSITORY_DIR", "/usr/helm-registry/charts")
		} else {
			_ = os.Setenv("REPOSITORY_DIR", "../test/chart") // TODO: change that for c:\user\user\document ..
		}
	}

	// Create directories

	// if REPOSITORY_DIR not exist, create it
	if _, err := os.Stat(os.Getenv("REPOSITORY_DIR")); os.IsNotExist(err) {
		err := os.MkdirAll(os.Getenv("REPOSITORY_DIR"), 0755)
		if err != nil {
			Logger.Error("Error creating directory : " + os.Getenv("REPOSITORY_DIR"))
		}
	}
}

func IsDocker() bool {
	_, err := os.Stat("/.dockerenv")
	if err == nil {
		Logger.Info("App running on Docker")
		return true
	} else if os.IsNotExist(err) {
		Logger.Info("App running not on Docker")
		return false
	} else {
		return false
	}
}
