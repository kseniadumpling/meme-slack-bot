package utils

import (
	"errors"
	"strings"
	"fmt"
	"time"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func handleAppMentionEvent(event *slackevents.AppMentionEvent, client *slack.Client) error {
	// Get user info
	user, err := client.GetUserInfo(event.User)
	if err != nil {
		return err
	}

	// Initialize the attachment to send it back to the slack
	attachment := slack.Attachment{}

	// Get the text
	text := strings.ToLower(event.Text)

	if strings.Contains(text, "hello") {
		attachment.Text = fmt.Sprintf("Hello %s!", user.Name)

	} else if strings.Contains(text, "wednesday") || strings.Contains(text, "today"){
		weekday := time.Now().Weekday()

		if int(weekday) == 3 {
			attachment.Text = fmt.Sprintf("It's Wednesday, my dude!")
			attachment.ImageURL = "https://i.kym-cdn.com/entries/icons/mobile/000/020/016/wednesdaymydudeswide.jpg"
		} else {
			attachment.Text = fmt.Sprintf("It's almost Wednesday, my dude!")
			attachment.ImageURL = "https://pbs.twimg.com/media/C92ykXiVYAMrgaT.jpg"
		}

	} else if strings.Contains(text, "random meme") {
		attachment.Text = fmt.Sprintf("Here you go, my dude!")
		// TODO: add a handler for random meme
		
	} else {
		wednesdayMsg := 	"\t - Ask me if it's Wednesday"
		todayMsg := 		"\t - Ask me what day is it today"
		randomMemeMsg := 	"\t - Ask me to send a random meme, I'll kindly do that"

		attachment.Text = fmt.Sprintf("Dear %s, this is what I can do for you: \n%s\n%s\n%s", user.Name, 
										wednesdayMsg, todayMsg, randomMemeMsg)
	}

	// Send the message to the channel
	_, _, err = client.PostMessage(event.Channel, slack.MsgOptionAttachments(attachment))
	if err != nil {
		return fmt.Errorf("Failed to post message: %w", err)
	}

	return nil
}

func HandleEventMessage(event slackevents.EventsAPIEvent, client *slack.Client) error {
	switch event.Type {
	case slackevents.CallbackEvent:
		innerEvent := event.InnerEvent

		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			err := handleAppMentionEvent(ev, client)
			if err != nil {
				return err
			}
		}

	default:
		return errors.New("Unsupported event type")
	}

	return nil
}