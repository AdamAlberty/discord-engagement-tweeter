package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type ConfigType struct {
	DbPath string `json:"db_path"`
	ReactionThreshold int `json:"reaction_threshold"`
	ChannelIds []string `json:"channel_ids"`
}

var Config = ConfigType{}

func LoadConfig(configPath *string) {
	fp, err := os.Open(*configPath)
	if err != nil {
		fmt.Println("ERROR: ", err)
		log.Fatalf("Could not find config file %s", *configPath)
	}
	defer fp.Close()

	bytes, _ := io.ReadAll(fp)

	if err := json.Unmarshal(bytes, &Config); err != nil {
		fmt.Println("ERROR in config file:", err)
		log.Fatal("Could not parse config file")
	}
}
