package db

import "time"

type Liked struct {
	ToLikePostID uint
	FromUserID   uint
	State        bool
	UpdateTime   time.Time `gorm:"autoUpdateTime:milli"`
	CreateTime   time.Time `gorm:"autoCreateTime:milli"`
}
