package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("ERROR", err)
	}

	// Initialize the bot
	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalln("ERROR: ", err)
	}

	// Create handler that listens to reactions
	discord.AddHandler(HandleNewReaction)

	// Start listening
	discord.Open()
	defer discord.Close()
	fmt.Println("Bot is running...")

	// Exit on interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
