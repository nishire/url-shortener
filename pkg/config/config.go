package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"url-shortener/pkg/constants"

	"github.com/spf13/viper"
)

//Init :
func Init(service, env, path string) {
	body, err := fetchConfiguration(service, path, env)
	if err != nil {
		fmt.Println("Couldn't load configuration, cannot start. Terminating. Error: " + err.Error())
	}
	parseConfiguration(body)
}

// read config json and return byte data
func fetchConfiguration(service, path, env string) ([]byte, error) {
	var bodyBytes []byte
	var err error
	result := strings.Compare(env, constants.DevEnvironment)
	if result == 0 {
		bodyBytes, err = ioutil.ReadFile(path + "/config/config.json")
		if err != nil {
			fmt.Println("Couldn't read local configuration file.", err)
		} else {
			log.Print("using local config.")
		}
	} else {
		fmt.Println("No configurations found for other environments...")
	}
	return bodyBytes, err
}

// Pass JSON bytes into struct and then into Viper
func parseConfiguration(body []byte) {
	var configData configObj
	err := json.Unmarshal(body, &configData)
	if err != nil {
		fmt.Println("Cannot parse configuration, message: " + err.Error())
	}
	for key, value := range configData.Property {
		viper.Set(key, value)
		fmt.Printf("Loading config property > %s - %s \n", key, value)
	}
}

type configObj struct {
	Name     string                 `json:"name"`
	Property map[string]interface{} `json:"property"`
}
