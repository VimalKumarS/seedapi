package util

import (
	"encoding/json"
	"fmt"
	"os"
	"seedapi/model"
)

// GetConfig read the
func GetConfig(fileName string) *model.AppSetting {
	if fileName == "" {
		fileName = "./appsetting/appsetting.json"
	} else {
		fileName = fmt.Sprint("./appsetting/appsetting.", fileName, ".json")
	}
	var config model.AppSetting
	configFile, err := os.Open(fileName)

	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
		return &model.AppSetting{
			Connection: "db connection string",
			Env:        "dev",
			Port:       "8000",
		}
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return &config
}
