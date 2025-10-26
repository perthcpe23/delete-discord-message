package services

import (
	"delete-discord-message/consts"
	"delete-discord-message/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"time"
)

// The sleep time is carefully selected. Reducing it can cause rate limit error.
const sleepBetweenDeleteBatch = time.Duration(10) * time.Second

func DeleteAllMessageInChannel(username string, deleteRequestBodyModel model.DeleteRequestBodyModel) error {
	messages, err := getAllMessagesInChannel(deleteRequestBodyModel)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("%s: total messages = %d\n", deleteRequestBodyModel.ChannelId, len(messages))

	totalMessageCount := 0
	for _, message := range messages {
		if message.Author.Username == username {
			totalMessageCount += 1
		}
	}

	deletedMessageCount := 0
	for _, message := range messages {
		if message.Author.Username != username {
			continue
		}

		_, err := deleteMessage(deleteRequestBodyModel.ChannelId, message.ID, deleteRequestBodyModel.Token)
		if err != nil {
			log.Println(err)
			return err
		}

		deletedMessageCount += 1

		if deletedMessageCount > 0 && deletedMessageCount%6 == 0 {
			log.Printf(
				"[%s] Deleted %d/%d messages, sleep %.0f seconds...\n",
				deleteRequestBodyModel.ChannelId,
				deletedMessageCount,
				totalMessageCount,
				sleepBetweenDeleteBatch.Seconds(),
			)
			time.Sleep(sleepBetweenDeleteBatch)
		}
	}

	log.Println("Done delete all messages")
	return nil
}

func getAllMessagesInChannel(deleteRequestBodyModel model.DeleteRequestBodyModel) ([]model.DiscordMessageModel, error) {
	beforeID := ""
	allMessages := []model.DiscordMessageModel{}

	for {
		url := fmt.Sprintf(
			"%s/channels/%s/messages?limit=50",
			consts.DiscordApiBaseUrl,
			deleteRequestBodyModel.ChannelId,
		)

		if beforeID != "" {
			url = fmt.Sprintf("%s&before=%s", url, beforeID)
		}

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("Authorization", deleteRequestBodyModel.Token)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		var messages []model.DiscordMessageModel
		err = json.Unmarshal(data, &messages)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		allMessages = append(allMessages, messages...)
		log.Printf("Fetched %d messages, total = %d", len(messages), len(allMessages))

		// TODO: do not rely on returned message count, use pagination info instead
		if len(messages) < 50 {
			break
		} else {
			beforeID = messages[len(messages)-1].ID
		}
	}

	// return the oldest message first
	slices.Reverse(allMessages)
	return allMessages, nil
}

func deleteMessage(channelID string, messageID string, token string) (statusCode int, err error) {
	url := fmt.Sprintf("%s/channels/%s/messages/%s", consts.DiscordApiBaseUrl, channelID, messageID)
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Add("Authorization", token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	data, _ := io.ReadAll(res.Body)

	statusCode = res.StatusCode
	log.Printf("Delete %s with %d: %s\n", messageID, statusCode, string(data))
	return
}
