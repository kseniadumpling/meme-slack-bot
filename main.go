package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {

	// Load env variables
	godotenv.Load(".env")

	token := os.Getenv("SLACK_AUTH_TOKEN")
	channelID := os.Getenv("SLACK_CHANNEL_ID")

	// Create a new client by passing the token, also set Debug mode
	client := slack.New(token, slack.OptionDebug(true))

	// A temp message that will be sent
	attachement := slack.Attachment{
		Pretext: "It's almost Wednesday, my dudes!",
		Text:	 "Look at fabulous green line near the text!",
		Color: 	 "#36a64f",
		Fields: []slack.AttachmentField{
			{
				Title: "Today is",
				Value: time.Now().Format(time.UnixDate),
			},
		},
	}

	// Post a message
	// The first param is a channelID, so skipping it
	_, timestamp, err := client.PostMessage(
		channelID,
		slack.MsgOptionAttachments(attachement),
	)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Message was sent at %s", timestamp)
}