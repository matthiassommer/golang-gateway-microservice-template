package main

import (
	"golang-gateway-microservice-template/router"
	. "golang-gateway-microservice-template/utils"
	"strconv"
)

const (
	defaultPort = 8080
	localPort   = 8080
)

func main() {
	Log.Info("[Gateway] Start")

	router := router.NewRouter()

	port := ":" + strconv.Itoa(defaultPort)
	if Environment() == ENV_LOCAL {
		port = ":" + strconv.Itoa(localPort)
	}

	Log.Fatal(router.Start(port))

	Log.Info("[Gateway] Started")
}
