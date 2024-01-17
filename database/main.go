package database

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func SaveMessage(messageId string) {
	fp, err := os.OpenFile("messages.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("ERROR writing to file:", err)
		return
	}
	defer fp.Close()
	fp.WriteString(fmt.Sprintf("%s %d\n", messageId, time.Now().Unix()))
}

func IsMessagePosted(messageId string) bool {
	fp, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("ERROR reading file:", err)
		return false
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		mid := strings.Split(line, " ")[0]

		if messageId == mid {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return false
}
