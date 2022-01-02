package db

import "time"

type Liked struct {
	ID           uint `gorm:"primaryKey;autoIncrement"`
	ToLikePostID uint
	FromUserID   uint
	State        bool
	UpdateTime   time.Time `gorm:"autoUpdateTime:milli"`
	CreateTime   time.Time `gorm:"autoCreateTime:milli"`
}
