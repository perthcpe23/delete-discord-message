package model

type DiscordMessageModel struct {
	ID     string `json:"id"`
	Author struct {
		Username string `json:"username"`
	}
}
