package db

import "time"

//这是关注者表
type Follow struct {
	ID           uint  `gorm:"autoIncrement"`
	UID          int64 `gorm:"not null"`
	FollowedUser int64 `gorm:"not null"`
	State        int8
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// TableName 会将 User 的表名重写为 `profiles`
func (Follow) TableName() string {
	return "follow"
}
