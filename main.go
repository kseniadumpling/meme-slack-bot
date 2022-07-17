package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

func main() {

	// Load env variables
	godotenv.Load(".env")

	token := os.Getenv("SLACK_AUTH_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")

	// Create a new client by passing the bot token, debug opt & app level token (for websockets)
	client := slack.New(token, slack.OptionDebug(true), slack.OptionAppLevelToken(appToken))

	// Config Socket Mode 
	socketClient := socketmode.New(
		client,
		socketmode.OptionDebug(true),
		// Some custom logger
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	// gotoutine context
	ctx, cancel := context.WithCancel(context.Background())

	// setup defer for cleaning up resources at the end
	defer cancel()

	go func(ctx context.Context, client *slack.Client, socketClient *socketmode.Client) {
		for {
			select {
			case <-ctx.Done():
				log.Println("Shutting down socketmode listener")
				return

			case event := <-socketClient.Events:
				switch event.Type {
				case socketmode.EventTypeConnecting:
					log.Println("Connecting to Slack with Socket Mode...")

				case socketmode.EventTypeConnectionError:
					log.Println("Connection failed. Retrying later...")

				case socketmode.EventTypeConnected:
					log.Println("Connected to Slack with Socket Mode.")

				case socketmode.EventTypeEventsAPI:
					eventsAPIEvent, ok := event.Data.(slackevents.EventsAPIEvent)
					if !ok {
						log.Printf("Could not type cast the event to the EventsAPIEvent: %v\n", event)
						continue
					}
					socketClient.Ack(*event.Request)
					log.Println(eventsAPIEvent)
				}
			}
		}
	}(ctx, client, socketClient)

	socketClient.Run()
}