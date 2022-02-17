package main

import (
	"flag"
	"fmt"
	"os"
	"url-shortener/internal/app/server"

	"url-shortener/go-helpers/logger"

	"url-shortener/go-helpers/config"
	"url-shortener/go-helpers/constants"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	service := "url-shortener"
	environment := os.Getenv("BOOT_CUR_ENV")
	if environment == "" {
		environment = constants.DevEnvironment
	}

	exitChannel := make(chan struct{})

	flag.Usage = func() {
		fmt.Println("Usage: server -s {service_name} -e {environment}")
		os.Exit(1)
	}
	flag.Parse()

	config.Init(service, environment, constants.ConfigFilePath)
	// logger.InitLogger()

	tracer.Start(tracer.WithEnv(environment),
		tracer.WithService(service),
		tracer.WithDebugMode(true))
	go server.Init()
	defer tracer.Stop()

	// Init Shutdown Signals & Actions
	// gracefulshutdown.Shutdown()

	// Blocking until the shutdown to complete then inform the main goroutine.
	<-exitChannel
	logger.SugarLogger.Info("main goroutine shutdown completed gracefully.")
}
