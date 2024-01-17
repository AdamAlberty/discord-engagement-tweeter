package handlers

import (
	"fmt"
	"time"

	"slices"

	"github.com/AdamAlberty/discord-engagement-tweeter/config"
	"github.com/AdamAlberty/discord-engagement-tweeter/database"
	"github.com/AdamAlberty/discord-engagement-tweeter/utils"
	"github.com/bwmarrin/discordgo"
)

// Listen to reactions being added
func HandleNewReaction(discord *discordgo.Session, message *discordgo.MessageReactionAdd) {
	if !(slices.Contains(config.Config.ChannelIds, message.ChannelID)) {
		return
	}
	
	// Get the reacted message
	messageDetail, err := discord.ChannelMessage(message.ChannelID, message.MessageID)
	if err != nil {
		fmt.Println("ERROR: `could not get message`")
		return
	}

	// If the message is older than 72 hours ignore it
	if time.Now().Sub(messageDetail.Timestamp).Hours() > 72 || database.IsMessagePosted(messageDetail.ID) {
		return
	}

	// Count total reactions
	total_count := 0
	for _, reaction := range messageDetail.Reactions {
		total_count += reaction.Count
	}

	// Post tweet when there are more reactions than the threshold specified
	if total_count < config.Config.ReactionThreshold {
		return
	}

	// If the tweet was posted successfully
	if utils.CreateTweet(messageDetail) {
		fmt.Println("TWEET POSTED SUCCESSFULLY")
		database.SaveMessage(messageDetail.ID)
	} else {
		fmt.Println("THERE HAS BEEN AN ERROR POSTING TWEET")
	}
}
