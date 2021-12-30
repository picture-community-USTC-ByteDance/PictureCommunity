package db

import "time"

type Forward struct {
	FID             uint `gorm:"primaryKey"`
	ToForwardPostID uint
	AuthorUserID    uint
	Content         string
	CommentNumber   int
	LikeNumber      int
	UpdateTime      time.Time `gorm:"autoUpdateTime:milli"`
	CreateTime      time.Time `gorm:"autoCreateTime:milli"`
}
