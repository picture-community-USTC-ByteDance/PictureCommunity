package _response

import "time"

type ResponseChatUsers struct {
	Uid      uint   `json:"uid"`
	Profile  string `json:"profile"`
	Nickname string `json:"nickname"`
	Unread   uint   `json:"unread_msg"`
}

type ResponseHistoryMsg struct {
	FromID    uint
	Content   string
	CreatedAt time.Time
}
