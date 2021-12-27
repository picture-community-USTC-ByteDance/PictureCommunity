package db

import "time"

type Post struct {
	PID              int64 `gorm:"primaryKey"`
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
