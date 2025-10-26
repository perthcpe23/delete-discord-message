package services

import (
	"delete-discord-message/consts"
	"delete-discord-message/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetUser(token string) (user model.UserModel, err error) {
	url := fmt.Sprintf("%s/users/@me", consts.DiscordApiBaseUrl)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Add("Authorization", token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		return
	}

	return
}
