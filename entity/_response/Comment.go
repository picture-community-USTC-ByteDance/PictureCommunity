package _response

import (
	"picture_community/entity/db"
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

type CreateFirstLevelCommentBack = db.Comment
type CreateSecondLevelCommentBack = db.Comment

//type DeleteCommentBack struct {
//	PID 		uint	`json:"post_id"`
//	ParentID 	uint	`json:"parent_id"`
//}
