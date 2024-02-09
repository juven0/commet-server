package configs

import (
	"encoding/json"
	"log"
	"os"
)

type Configs struct {
	Server struct {
		Port string `json:"port"`
		Host string `json:"host"`
	} `json:"server"`
	Api struct {
		BaseUrl   string `json:"baseurl"`
		SecretKey string `jsno:"secretkey"`
	} `json:"api"`
}

func (Configs) LoadConfigs() *Configs {
	var configs Configs
	confFile, err := os.Open("/configs/configs.json")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer confFile.Close()
	jsonDecode := json.NewDecoder(confFile)
	jsonDecode.Decode(&configs)
	return &configs
}
