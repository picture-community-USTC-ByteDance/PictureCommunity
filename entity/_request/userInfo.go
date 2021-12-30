package _request

import "time"

type UserInfo struct {
	NickName  string    `json:"nickname"`
	Sex       int8      `json:"sex"`
	Motto     string    `json:"motto"`
	Email     string    `json:"email"`
	Telephone int       `json:"telephone"`
	Birthday  time.Time `json:"birthday"`
	Address   string    `json:"address"`
	Profile   string    `json:"profile"` //头像url
}
