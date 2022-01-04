package _request

type CreatePost struct {
	Content string `form:"content" json:"content"`
}

type NewForward struct {
	PID     uint   `form:"post_id" json:"post_id" binding:"required"`
	Content string `form:"content" json:"content"`
}
