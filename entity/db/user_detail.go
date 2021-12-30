package db

import "time"

type UserDetail struct {
	UID       int64 `gorm:"primaryKey"`
	NickName  string
	Sex       int8
	UpdatedAt time.Time
	Birthday  time.Time
	Address   string
	Motto     string
	Profile   string
}
