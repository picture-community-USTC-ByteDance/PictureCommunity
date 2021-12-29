package db

import "time"

type Follow struct {
	ID           int64 `gorm:"autoIncrement"`
	UID          int64 `gorm:"not null"`
	FollowedUser int64 `gorm:"not null"`
	State        int8
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
