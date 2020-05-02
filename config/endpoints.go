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

func getURL(defaultURL string, service string) string {
	envName := fmt.Sprintf("%s_SERVICE_URL", strings.ToUpper(service))
	if value, exists := os.LookupEnv(envName); exists {
		return value
	}
	return defaultURL
}

func loadServiceEndpoints() {
	PizzaServiceURL = getURL(localPizzaServiceURL, "data")

	utils.Log.Infof(`Using the following endpoints: pizza-service: %s`, PizzaServiceURL)
}

// init determines service endpoint URLs by environment variable. If not set, uses default values.
func init() {
	loadServiceEndpoints()
}
