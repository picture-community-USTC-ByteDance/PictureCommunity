package _response

import "time"

type QueryCommentBack struct {
	UserID   int    `json:"userId"`
	NickName string `json:"nickName"`
	//	Profile     string //头像略缩图
	ChildNumber int    //子评论个数
	Content     string //内容
}
type QueryCommentBackTemp struct {
	CID         uint `form:"c_id" json:"c_id" binding:"required"` //评论id
	UserID      uint //(评论)用户id
	NickName    string
	Profile     string    //略缩图
	Content     string    //内容
	UpdateTime  time.Time //评论时间
	LikeNumber  uint      //评论的点赞数
	ChildNumber int       //子评论个数
	LikeStatus  bool      //true表示已点赞
}
type QueryCommentBackTemp2 struct {
	CID uint `form:"c_id" json:"c_id" binding:"required"` //评论id

	UserID   uint //(评论)用户id
	NickName string
	Profile  string //略缩图

	Content    string    //内容
	UpdateTime time.Time //评论时间
	LikeNumber uint      //评论的点赞数
	LikeStatus bool      //true表示已点赞
	ParentId   uint      `form:"parentid" json:"parentid"`
}

type CreateFirstLevelCommentBack struct {
	PID     uint   `json:"post_id"`
	Content string //评论内容
}
type CreateSecondLevelCommentBack struct {
	PID      uint   `json:"post_id"`
	ParentID uint   `json:"parent_id"`
	Content  string //评论内容
}

//type DeleteCommentBack struct {
//	PID 		uint	`json:"post_id"`
//	ParentID 	uint	`json:"parent_id"`
//}
