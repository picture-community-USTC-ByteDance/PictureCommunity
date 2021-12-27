package db

import "time"

type UserDetail struct {
	UID        int64 `gorm:"primaryKey"`
	Nickname   string
	Sex        bool
	UpdateDate time.Time
	Birthday   time.Time
	Address    string
	Motto      string
	Profile    string
}
