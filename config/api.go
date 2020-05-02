package config

import (
	. "golang-gateway-microservice-template/utils"
	"os"
)

const (
	envAPIVersion = "API_VERSION"

	defaultAPIVersion = "v1"
)

var (
	// APIVersion is the API version to be used in calls to other microservices.
	APIVersion = defaultAPIVersion
)

// loadAPISettings gets API settings using environment variables, or if not set uses a default value.
func loadAPISettings() {
	if value, exists := os.LookupEnv(envAPIVersion); exists {
		APIVersion = value
	}

	Log.Infof("using API version %s", APIVersion)
}

func init() {
	loadAPISettings()
}
