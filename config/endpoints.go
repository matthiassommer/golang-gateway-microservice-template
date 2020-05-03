package config

import (
	"fmt"
	"os"
	"strings"

	"golang-gateway-microservice-template/utils"
)

// Default service endpoints for local environment.
const (
	localPizzaServiceURL = "http://127.0.0.1:8081"
)

var (
	// PizzaServiceURL is the url to access the pizza microservice.
	PizzaServiceURL string
)

func init() {
	loadServiceEndpoints()
}

func loadServiceEndpoints() {
	PizzaServiceURL = getURL(localPizzaServiceURL, "pizza")

	utils.Log.Infof(`Using the following endpoints: pizza-service: %s`, PizzaServiceURL)
}

// getURL loads the microservice endpoint from a environment variable or returns the default for local testing.
func getURL(defaultURL string, service string) string {
	envName := fmt.Sprintf("%s_SERVICE_URL", strings.ToUpper(service))
	if value, exists := os.LookupEnv(envName); exists {
		return value
	}
	return defaultURL
}
