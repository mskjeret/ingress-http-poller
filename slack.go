package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

//NotifySlack send message to slack with the given key to the givem channel
func NotifySlack(channel string, key string, message string) bool {

	api := slack.New(key)

	channelID, timestamp, err := api.PostMessage(channel, slack.MsgOptionText(message, false))
	if err != nil {
		fmt.Printf("%s\n", err)
		return false
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

	return true
}
