package db

import "time"

/*ToForwardPostID表示帖子id
AuthorUserID表示转发了这个帖子的用户id*/
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
