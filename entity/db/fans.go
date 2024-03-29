package db

import "time"

type Fans struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	UID        uint
	FansID     uint
	state      bool      //true为fansID是UID的粉丝
	UpdateTime time.Time `gorm:"autoUpdateTime:milli"`
	CreateTime time.Time `gorm:"autoCreateTime:milli"`
}
