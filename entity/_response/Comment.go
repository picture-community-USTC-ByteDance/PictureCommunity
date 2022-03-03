package _response

import (
	"time"
)

type QueryCommentBack struct {
	CID         uint //评论id
	UserID      uint //(评论)用户id\
	NickName    string
	Profile     string    //略缩图
	Content     string    //内容
	UpdateTime  time.Time //评论时间
	LikeNumber  uint      //评论的点赞数
	ChildNumber int       //子评论个数
	Username    string
}
type QueryCommentBack2 struct {
	CID uint //评论id

	UserID   uint //(评论)用户id
	NickName string
	Profile  string //略缩图

	Content    string    //内容
	UpdateTime time.Time //评论时间
	LikeNumber uint      //评论的点赞数
	ParentId   uint
	Username   string
}

type CreateCommentBack struct {
	CID          uint `gorm:"primaryKey"`
	PostID       uint
	ParentID     uint      //如果为一级评论则为空
	ChildNumber  int       //子评论个数
	Content      string    //内容
	UserID       uint      //评论作者id
	LikeNumber   uint      //评论的点赞数
	DeleteStatus bool      //评论是否被删除 true表示删除
	UpdateTime   time.Time `gorm:"autoUpdateTime:milli"`
	CreateTime   time.Time `gorm:"autoCreateTime:milli"`
	NickName     string
	Profile      string //略缩图
}

//type DeleteCommentBack struct {
//	PID 		uint	`json:"post_id"`
//	ParentID 	uint	`json:"parent_id"`
//}
