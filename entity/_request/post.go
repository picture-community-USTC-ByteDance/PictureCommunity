package _request

type CreatePost struct {
	Url     []string `form:"photo_url" json:"photo_url" binding:"required"`
	Content string   `form:"content" json:"content"`
}

type DeletePost struct {
	PID uint `form:"post_id" json:"post_id" binding:"required"`
}

type NewForward struct {
	PID     uint   `form:"post_id" json:"post_id" binding:"required"`
	Content string `form:"content" json:"content"`
}

type DeleteForward struct {
	FID uint `form:"forward_id" json:"forward_id" binding:"required"`
}
