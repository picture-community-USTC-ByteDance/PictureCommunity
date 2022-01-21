package db

import "time"

type Collection struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	UID        uint
	PID        uint
	State      bool
	UpdateTime time.Time `gorm:"autoUpdateTime:milli"`
	CreateTime time.Time `gorm:"autoCreateTime:milli"`
}
