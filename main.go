package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/AdamAlberty/discord-engagement-tweeter/config"
	"github.com/AdamAlberty/discord-engagement-tweeter/handlers"
	"github.com/AdamAlberty/discord-engagement-tweeter/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("ERROR", err)
	}

	// Load config
	configPath := flag.String("config", "config.json", "Path to config file, defaults to config.json")
	config.LoadConfig(configPath)
	utils.InitMessageDB()

	// Initialize Discord bot
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Fatalln("ERROR: ", err)
	}

	// Create handler that listens to reactions
	discord.AddHandler(handlers.HandleNewReaction)

	// Start listening
	if err := discord.Open(); err != nil {
		log.Fatal("ERROR: ", err)
	}
	defer discord.Close()
	fmt.Println("Bot is running...")

	// Exit on interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
