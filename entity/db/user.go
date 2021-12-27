package db

import "time"

type User struct {
	UID          int64  `gorm:"primaryKey"`
	Username     string `gorm:"unique"`
	Password     string
	Telephone    int    `gorm:"unique"`
	Email        string `gorm:"unique"`
	Status       int8
	RegisterDate time.Time
}
