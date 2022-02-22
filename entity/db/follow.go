package db

import "time"

type Follow struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	UID        uint
	FollowedID uint
	State      bool      //true为已关注
	UpdateTime time.Time `gorm:"autoUpdateTime:milli"`
	CreateTime time.Time `gorm:"autoCreateTime:milli"`
}
