package entity

import "time"

type LoginUser struct {
	Info     string `form:"info" json:"info" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Method   int8   `form:"method" json:"method"`
}

type CreatePost struct {
	ID       int64  `form:"u_id" json:"u_id" binding:"required"`
	ImageUrl string `form:"image_url" json:"image_url" binding:"required"`
	Content  string `form:"content" json:"content" binding:"required"`
	Token    string `form:"token" json:"token" binding:"required"`
}

type SearchUsers struct {
	NickName string `form:"nickname"  binding:"required"`
	PageSize int    `form:"pagesize"  binding:"required"`
	Page     int    `form:"page"  binding:"required"`
}
type UpdatePost struct {
	ID int64 `form:"u_id" json:"u_id" binding:"required"`
	//UpdateTime time.Time `form:"update_time" json:"update_time" binding:"required"`
	Content UserInfo `form:"content" json:"content" binding:"required"`
}
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
