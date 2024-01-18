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
	IgnoreAfterHours int `json:"ignore_after_hours"`

	DiscordBotToken string
	XConsumerKey string
	XConsumerSecret string
	XAcessToken string
	XTokenSecret string
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
	Config.DiscordBotToken = os.Getenv("DISCORD_BOT_TOKEN")
	Config.XConsumerKey = os.Getenv("X_CONSUMER_KEY")
	Config.XConsumerSecret = os.Getenv("X_CONSUMER_SECRET")
	Config.XAcessToken = os.Getenv("X_ACCESS_TOKEN")
	Config.XTokenSecret = os.Getenv("X_TOKEN_SECRET")
}
