package main

import (
	"delete-discord-message/model"
	"delete-discord-message/services"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: delete-discord-message <token> <channel_id1> [channel_id2 channel_id3 ...]")
		return
	}

	token := os.Args[1]
	channelIds := os.Args[2:]

	user, err := services.GetUser(token)
	if err != nil {
		fmt.Println("Error fetching user:", err)
		return
	}

	for _, channelId := range channelIds {
		err := services.DeleteAllMessageInChannel(user.Username, model.DeleteRequestBodyModel{
			ChannelId: channelId,
			Token:     token,
		})

		if err != nil {
			fmt.Printf("Error deleting messages in channel %s: %s\n", channelId, err.Error())
		}
	}
}
