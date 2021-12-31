package db

import "time"

type Follow struct {
	UID        uint
	FollowedID uint
	state      bool      //true为已关注
	UpdateTime time.Time `gorm:"autoUpdateTime:milli"`
	CreateTime time.Time `gorm:"autoCreateTime:milli"`
}
