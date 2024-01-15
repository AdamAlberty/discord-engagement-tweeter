package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// Listen to reactions being added
func HandleNewReaction(discord *discordgo.Session, message *discordgo.MessageReactionAdd) {
	if message.ChannelID != os.Getenv("CHANNEL_ID") {
		return
	}

	// Get the reacted message
	messageDetail, err := discord.ChannelMessage(os.Getenv("CHANNEL_ID"), message.MessageID)
	if err != nil {
		fmt.Println("ERROR: `could not get message`")
	}

	// Count total reactions
	total_count := 0
	for _, reaction := range messageDetail.Reactions {
		total_count += reaction.Count
	}

	reaction_threshold, err := strconv.Atoi(os.Getenv("SEND_TWEET_AFTER"))
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	// Post tweet when there are more reactions than the threshold specified
	if total_count >= reaction_threshold {
		CreateTweet(messageDetail)
	}
}
