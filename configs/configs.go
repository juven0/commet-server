package configs

import (
	"encoding/json"
	"log"
	"os"
)

type Configs struct {
	Port      string `json:"port"`
	Host      string `json:"host"`
	BaseUrl   string `json:"baseurl"`
	SecretKey string `jsno:"secretkey"`
}

func (Configs) LoadConfigs() *Configs {
	var configs Configs
	confFile, err := os.Open("./configs.json")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer confFile.Close()
	jsonDecode := json.NewDecoder(confFile)
	jsonDecode.Decode(&configs)
	return &configs
}

func NewConfig() (Configs, error) {
	var cf Configs
	confFile, err := os.Open("./configs.json")
	defer confFile.Close()

	if err != nil {
		return cf, err
	}
	jsonParser := json.NewDecoder(confFile)
	jsonParser.Decode(&cf)
	if err != nil {
		return cf, err
	}

	return cf, nil
}
