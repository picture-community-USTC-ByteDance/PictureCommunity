package db

import "time"

type Comment struct {
	CID          uint `gorm:"primaryKey"`
	PostID       uint
	ParentID     uint      //如果为一级评论则为空
	ChildNumber  int       //子评论个数
	Content      string    //内容
	UserID       uint      //评论作者id
	Profile      string    //评论者头像
	NickName     string    //评论者昵称
	LikeNumber   uint      //评论的点赞数
	LikeStatus   bool      //评论是否被点赞 true表示已点赞
	DeleteStatus bool      //评论是否被删除 true表示删除
	UpdateTime   time.Time `gorm:"autoUpdateTime:milli"`
	CreateTime   time.Time `gorm:"autoCreateTime:milli"`
}
