package db

import (
	"time"
)

type UserDetail struct {
	UID           uint
	Nickname      string
	Sex           bool      //false为女  true为男
	Birthday      time.Time `gorm:"autoCreateTime:milli"`
	Address       string
	Motto         string
	Profile       string    //略缩图
	OriginProfile string    //详细头像url
	UpdateTime    time.Time `gorm:"autoUpdateTime:milli"`
	CreateTime    time.Time `gorm:"autoCreateTime:milli"`
}
