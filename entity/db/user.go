package db

import "time"

type User struct {
	UID        uint   `gorm:"primaryKey"`
	Username   string `gorm:"unique"`
	Password   string
	Telephone  uint   `gorm:"unique"`
	Email      string `gorm:"unique"`
	Status     int8
	UpdateTime time.Time `gorm:"autoUpdateTime:milli"`
	CreateTime time.Time `gorm:"autoCreateTime:milli"`
}
