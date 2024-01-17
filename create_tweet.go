package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	oauth1 "github.com/klaidas/go-oauth1"
)

func CreateTweet(message *discordgo.Message) bool {
	// Tweet payload
	postBody, err := json.Marshal(map[string]string{
		"text": message.Content,
	})
	if err != nil {
		fmt.Println("ERROR: ", err)
		return false
	}

	// Post the tweet
	req, err := http.NewRequest("POST", "https://api.twitter.com/2/tweets", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println("ERROR: ", err)
		return false
	}

	// Construct oauth 1 header for authentication
	auth := oauth1.OAuth1{
		ConsumerKey: os.Getenv("X_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("X_CONSUMER_SECRET"),
		AccessToken: os.Getenv("X_ACCESS_TOKEN"),
		AccessSecret: os.Getenv("X_TOKEN_SECRET"),
	}
	authHeader := auth.BuildOAuth1Header("POST", "https://api.twitter.com/2/tweets", 
	map[string]string {})
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authHeader)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return false
	}

	var data map[string]any
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return false
	}

	_, exists := data["data"].(map[string]interface{})
	if !exists {
    	fmt.Println("ERROR: Could not post tweet")

		_, exists = data["detail"].(string)
		if exists {
    		fmt.Println("ERROR: ", data["detail"])
    		return false
		}
	}

	return true
}
