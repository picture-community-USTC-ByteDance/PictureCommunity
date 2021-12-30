package db

import "time"

//这是粉丝表
type Fans struct {
	ID        uint  `gorm:"autoIncrement"`
	UID       int64 `gorm:"not null"`
	follower  int64 `gorm:"not null"`
	State     int8
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName 会将 User 的表名重写为 `profiles`
func (Fans) TableName() string {
	return "fan"
}
