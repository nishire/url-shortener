package server

import (
	"url-shortener/internal/app/router"

	"github.com/golang/glog"

	"github.com/spf13/viper"
)

func Init() {
	r := router.NewRouter()
	err := r.Run(viper.GetString("ServerPort"))
	if err != nil {
		glog.Error("Server not able to startup with error: ", err)
	}
}
