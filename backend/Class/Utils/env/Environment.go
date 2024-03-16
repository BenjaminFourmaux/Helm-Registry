package env

import (
	"backend/Class/Logger"
	"os"
)

// Using var instead of os.getEnv() for improve security against EnvVarInjection
// That store en var in the program memory instead of each time calling os.getEnv and get changed value
var (
	IS_DOCKER               bool   // IS_DOCKER Env var - True if on Docker env, False if is on Windows or Linux
	INDEX_FILE_PATH         string // INDEX_FILE_PATH Env var - Path to the index.yaml file
	REPOSITORY_DIR          string // REPOSITORY_DIR Env var - Path to the repository folder
	REGISTRY_NAME           string // REGISTRY_NAME Env var - Name of this registry
	REGISTRY_DESCRIPTION    string // REGISTRY_DESCRIPTION Env var - Description of this registry
	REGISTRY_VERSION        string // REGISTRY_VERSION Env var - Version of this registry
	REGISTRY_MAINTAINER     string // REGISTRY_MAINTAINER Env var - Name of the maintainer of this registry
	REGISTRY_MAINTAINER_URL string // REGISTRY_MAINTAINER_URL Env var - URL of the maintainer of this registry
	REGISTRY_LABELS         string // REGISTRY_LABELS Env var - List of labels for this registry
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
			//userDocs := os.Getenv("USERPROFILE") + "\\Documents\\helm-registry\\charts"
			_ = os.Setenv("REPOSITORY_DIR", "../test/chart") // TODO : remplace by userDocs
		}
	}

	// Save env var change after permutation in class properties
	IS_DOCKER = isDocker
	INDEX_FILE_PATH = os.Getenv("INDEX_FILE_PATH")
	REPOSITORY_DIR = os.Getenv("REPOSITORY_DIR")
	REGISTRY_NAME = os.Getenv("REGISTRY_NAME")
	REGISTRY_DESCRIPTION = os.Getenv("REGISTRY_DESCRIPTION")
	REGISTRY_VERSION = os.Getenv("REGISTRY_VERSION")
	REGISTRY_MAINTAINER = os.Getenv("REGISTRY_MAINTAINER")
	REGISTRY_MAINTAINER_URL = os.Getenv("REGISTRY_MAINTAINER_URL")
	REGISTRY_LABELS = os.Getenv("REGISTRY_LABELS")

	// Create directories

	// if REPOSITORY_DIR not exist, create it
	if _, err := os.Stat(REPOSITORY_DIR); os.IsNotExist(err) {
		err := os.MkdirAll(REPOSITORY_DIR, 0755)
		if err != nil {
			Logger.Error("Error creating directory : " + REPOSITORY_DIR)
		} else {
			Logger.Success("Creating REPOSITORY_DIR on : " + REPOSITORY_DIR)
		}
	}
}

func IsDocker() bool {
	_, err := os.Stat("/.dockerenv")
	if err == nil {
		Logger.Info("App running on Docker")
		return true
	} else if os.IsNotExist(err) {
		Logger.Info("App not running on Docker")
		return false
	} else {
		return false
	}
}
