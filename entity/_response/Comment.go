package _response

type QueryCommentBack struct {

	UID      int    `json:"userId"`
	NickName string `json:"nickName"`
	//	Profile     string //头像略缩图
	ChildNumber int    //子评论个数
	Content     string //内容
}
type QueryCommentBackTemp struct {

	ChildNumber int //子评论个数
	LikeNumber  int
	Content     string //内容
	UserID      uint   //评论作者id
}
type QueryCommentBackTemp2 struct {
	//	ChildNumber int    //子评论个数
	LikeNumber int
	Content    string //内容
	UserID     uint   //评论作者id
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
