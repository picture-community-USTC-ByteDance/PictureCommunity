package db

import "time"

type Comment struct {
	CID         uint `gorm:"primaryKey"`
	PostID      uint
	ParentID    uint      //如果为一级评论则为空
	ChildNumber int       //子评论个数
	Content     string    //内容
	UserID      uint      //评论作者id
	UpdateTime  time.Time `gorm:"autoUpdateTime:milli"`
	CreateTime  time.Time `gorm:"autoCreateTime:milli"`
}
