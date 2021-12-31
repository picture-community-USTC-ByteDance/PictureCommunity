package db

import "time"

type UserData struct {
	UID              uint
	FollowersNumber  int
	FansNumber       int
	PostsNumber      int
	CollectionNumber int
	ForwardNumber    int
	UpdateTime       time.Time `gorm:"autoUpdateTime:milli"`
	CreateTime       time.Time `gorm:"autoCreateTime:milli"`
}
