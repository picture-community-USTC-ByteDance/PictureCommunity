package _request

import "time"

type UserInfo struct {
	Nickname  string    `json:"nickname"`
	Sex       bool      `json:"sex"`
	Motto     string    `json:"motto"`
	Email     string    `json:"email"`
	Telephone int       `json:"telephone"`
	Birthday  time.Time `json:"birthday"`
	Address   string    `json:"address"`
	Profile   string    `json:"profile"` //头像url
}