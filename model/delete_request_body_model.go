package model

type DeleteRequestBodyModel struct {
	ChannelId string `json:"channel_id"`
	Token     string `json:"token"`
}
