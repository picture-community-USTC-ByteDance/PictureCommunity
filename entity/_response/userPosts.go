package _response

// 用户的帖子数组信息
type UserPosts struct {
	PID           uint   `json:"pid"`
	TitlePhotoUrl string `json:"title_photo_url"`
	LikeNumber    int    `json:"like_number"`
	CommentNumber int    `json:"comment_number"`
	ForwardNumber int    `json:"forward_number"`
}
