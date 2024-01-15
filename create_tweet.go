package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

func CreateTweet(message *discordgo.Message) {
	// Tweet payload
	postBody, err := json.Marshal(map[string]string{
		"text": "something",
	})
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	// Post the tweet
	req, err := http.NewRequest("POST", "https://api.twitter.com/2/tweets", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("TWITTER_BEARER")))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	var data map[string]any;
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	// Print the response
	fmt.Printf("%+v\n", data)
}
