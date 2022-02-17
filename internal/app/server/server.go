package server

import (
	"fmt"
	"url-shortener/internal/app/router"

	// "url-shortener/go-helpers/logger"

	"github.com/spf13/viper"
)

func Init() {
	r := router.NewRouter()
	err := r.Run(viper.GetString("server.port"))
	if err != nil {
		fmt.Println("Server not able to startup with error: ", err) // need to add custom log
	}
}
