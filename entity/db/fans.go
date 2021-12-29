package db

import "time"

//这是粉丝表
type Fans struct {
	ID        int64 `gorm:"autoIncrement"`
	UID       int64 `gorm:"not null"`
	follower  int64 `gorm:"not null"`
	State     int8
	CreatedAt time.Time
	UpdatedAt time.Time
}
