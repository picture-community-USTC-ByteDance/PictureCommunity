package _request

type CreatePost struct {
	Url     string `form:"title_photo_url" json:"title_photo_url" binding:"required"`
	Content string `form:"content" json:"content"`
}

type NewForward struct {
	PID     uint   `form:"post_id" json:"post_id" binding:"required"`
	Content string `form:"content" json:"content"`
}
