package db

import (
	"time"
)

type PostPhoto struct {
	ID           uint
	PID          uint
	UID          uint
	Url          string    `gorm:"size:400"`
	SerialNumber uint      //photo 的编号，title_photo的编号为0，每个post最多10个photo
	UpdateTime   time.Time `gorm:"autoUpdateTime:milli"`
	CreateTime   time.Time `gorm:"autoCreateTime:milli"`
}
