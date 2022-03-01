package _response

import (
	"picture_community/entity/db"
	"time"
)

type QueryCommentBackTemp struct {
	CID         uint //评论id
	UserID      uint //(评论)用户id
	NickName    string
	Profile     string    //略缩图
	Content     string    //内容
	UpdateTime  time.Time //评论时间
	LikeNumber  uint      //评论的点赞数
	ChildNumber int       //子评论个数
	LikeStatus  bool      //true表示已点赞
	Username    string
}
type QueryCommentBackTemp2 struct {
	CID uint //评论id

	UserID   uint //(评论)用户id
	NickName string
	Profile  string //略缩图

	Content    string    //内容
	UpdateTime time.Time //评论时间
	LikeNumber uint      //评论的点赞数
	LikeStatus bool      //true表示已点赞
	ParentId   uint
	Username   string

}

type CreateFirstLevelCommentBack = db.Comment
type CreateSecondLevelCommentBack = db.Comment

//type DeleteCommentBack struct {
//	PID 		uint	`json:"post_id"`
//	ParentID 	uint	`json:"parent_id"`
//}
