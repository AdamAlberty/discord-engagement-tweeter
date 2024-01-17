package utils

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/AdamAlberty/discord-engagement-tweeter/config"
)

func InitMessageDB() {
	// Create file (if it doesn't exist) that holds messages that are already posted
	_, error := os.Stat(config.Config.DbPath)
	if errors.Is(error, os.ErrNotExist) {
		fp, err := os.Create(config.Config.DbPath)
		if err != nil {
			log.Fatal("ERROR: ", err)
		}
		if err = fp.Close(); err != nil {
			fmt.Println("ERROR: ", err)
		}
	}
}
