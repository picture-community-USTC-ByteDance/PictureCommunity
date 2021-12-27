package entity

import (
	"time"
)

//用户
type User struct {
	ID           int64
	Username     string
	Password     string
	Telephone    int
	Email        string
	Status       int8
	RegisterDate time.Time `gorm:"register_date"`
}

type Post struct {
	ID               int64 `gorm:"autoIncrement"`
	UID              int64
	PhotoNumber      uint
	Content          string `gorm:"size:400"`
	TitlePhotoUrl    string `gorm:"size:1000"`
	CommentNumber    int
	ForwardNumber    int
	LikeNumber       int
	CollectionNumber int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type UserDetail struct {
	ID        int64 `gorm:"u_id"`
	NickName  string
	Sex       int8
	UpdatedAt time.Time
	Birthday  time.Time
	Address   string
	Motto     string
	Profile   string
}
