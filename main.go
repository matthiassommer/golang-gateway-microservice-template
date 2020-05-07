package main

import (
	"golang-gateway-microservice-template/router"
	. "golang-gateway-microservice-template/utils"
	"os"
	"os/signal"
	"strconv"
	"syscall"
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

	go func() {
		Log.Fatal(router.Start(port))
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)
	<-done

	router.Shutdown()
}
