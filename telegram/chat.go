package telegram

import (
	"encoding/json"
	"fmt"
)

type ChatModel struct {
	ID                          int    `json:"id"`
	Title                       string `json:"title"`
	Type                        string `json:"type"`
	AllMembersAreAdministrators bool   `json:"all_members_are_administrators"`
}

type FromModel struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type EntitieModel struct {
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}

type MessageModel struct {
	MessageID int            `json:"message_id"`
	From      FromModel      `json:"from"`
	Chat      ChatModel      `json:"chat"`
	Date      int            `json:"date"`
	Text      string         `json:"text"`
	Entities  []EntitieModel `json:"entities"`
}

type ResultModel struct {
	UpdateID int          `json:"update_id"`
	Message  MessageModel `json:"message"`
}

func Chat(message []byte) {
	dat := ResultModel{}
	if err := json.Unmarshal(message, &dat); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(dat)
}
