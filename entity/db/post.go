package db

import (
	"time"
)

type Post struct {
	PID              uint `gorm:"primaryKey"`
	UID              uint
	PhotoNumber      int
	Content          string `gorm:"size:400"`
	TitlePhotoUrl    string `gorm:"size:80"`
	CommentNumber    int
	ForwardNumber    int
	LikeNumber       int
	CollectionNumber int
	UpdateTime       time.Time `gorm:"autoUpdateTime:milli"`
	CreateTime       time.Time `gorm:"autoCreateTime:milli"`
}
