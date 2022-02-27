package _response

import "time"

type ResponseChatUsers struct {
	Uid      uint   `json:"uid"`
	Profile  string `json:"profile"`
	Nickname string `json:"nickname"`
	Unread   uint   `json:"unread_msg"`
}

type ResponseHistoryMsg struct {
	FromID    uint      `json:"from_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
