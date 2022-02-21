package main

import (
	"os"
	"url-shortener/internal/app/server"

	"url-shortener/pkg/config"
	"url-shortener/pkg/constants"
)

func main() {
	service := "url-shortener"
	environment := os.Getenv("BOOT_CUR_ENV")
	if environment == "" {
		environment = constants.DevEnvironment
	}

	// initialize config
	config.Init(service, environment, constants.ConfigFilePath)

	// initialize server
	server.Init()
}
