package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	oauth1 "github.com/klaidas/go-oauth1"

	_ "image/jpeg"
	_ "image/png"
)

func CreateTweet(message *discordgo.Message) bool {
	// Upload media of the tweet
	mediaId := ""
	for _, att := range message.Attachments {
		if strings.Contains(att.ContentType, "image") {
			mediaId = uploadMedia(att)
			if mediaId == "" {
				fmt.Println("ERROR uploading image")
				return false
			}
			break
		}
	}
	var postBody []byte
	if mediaId != "" {
		// Tweet payload
		var err error
		postBody, err = json.Marshal(map[string]any{
			"text": message.Content,
			"media": map[string]any{
				"media_ids": [1]string{mediaId},
			},
		})
		if err != nil {
			fmt.Println("ERROR: ", err)
			return false
		}
	} else {
		var err error
		postBody, err = json.Marshal(map[string]any{
			"text": message.Content,
		})
		if err != nil {
			fmt.Println("ERROR: ", err)
			return false
		}
	}

	// Construct request
	req, err := http.NewRequest("POST", "https://api.twitter.com/2/tweets", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println("ERROR: ", err)
		return false
	}

	// Construct oauth 1 header for authentication
	auth := oauth1.OAuth1{
		ConsumerKey:    os.Getenv("X_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("X_CONSUMER_SECRET"),
		AccessToken:    os.Getenv("X_ACCESS_TOKEN"),
		AccessSecret:   os.Getenv("X_TOKEN_SECRET"),
	}
	authHeader := auth.BuildOAuth1Header("POST", "https://api.twitter.com/2/tweets",
		map[string]string{})
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authHeader)

	// Post the tweet
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return false
	}

	// Parse response
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

func uploadMedia(attachment *discordgo.MessageAttachment) string {
	fmt.Println("Uploading media")

	// Get image from discord
	res, err := http.Get(attachment.URL)
	if err != nil || res.StatusCode != 200 {
		fmt.Println("ERROR:", err)
		return ""
	}
	defer res.Body.Close()
	imageData, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ERROR:", err)
		return ""
	}

	// Create a buffer to store the multipart form data
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	// Add the base64-encoded image as a file field in the form
	imageField, err := writer.CreateFormField("media_data")
	if err != nil {
		fmt.Println("Error creating form field:", err)
		return ""
	}
	base64Encoded := base64.StdEncoding.EncodeToString(imageData)
	imageField.Write([]byte(base64Encoded))

	// Close the multipart writer to finalize the form data
	writer.Close()

	// Post the tweet
	req, err := http.NewRequest("POST", "https://upload.twitter.com/1.1/media/upload.json", &buffer)
	if err != nil {
		fmt.Println("ERROR:", err)
		return ""
	}

	// Construct oauth 1 header for authentication
	auth := oauth1.OAuth1{
		ConsumerKey:    os.Getenv("X_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("X_CONSUMER_SECRET"),
		AccessToken:    os.Getenv("X_ACCESS_TOKEN"),
		AccessSecret:   os.Getenv("X_TOKEN_SECRET"),
	}
	authHeader := auth.BuildOAuth1Header("POST", "https://upload.twitter.com/1.1/media/upload.json",
		map[string]string{})
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authHeader)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	res, err = client.Do(req)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return ""
	}

	var data map[string]any
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return ""
	}

	fmt.Println(data)
	mediaId, exists := data["media_id_string"].(string)
	if !exists {
		fmt.Println("ERROR: Could not upload media tweet")
		return ""
	} else {
		return mediaId
	}

}
