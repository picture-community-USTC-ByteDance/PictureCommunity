package entity

import (
	"time"
)

//用户
type User struct {
	ID        int64  `db: u_id`
	Username  string `db: username`
	Password  string `db: password`
	Telephone string `db: telephone`
	Email     string `db: email`
	Status    int8   `db: status`
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
