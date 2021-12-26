package entity

type LoginUser struct {
	Info     string `form:"info" json:"info" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Method   int8   `form:"method" json:"method"`
}

type CreatePost struct {
	ID       int64  `form:"u_id" json:"u_id" binding:"required"`
	ImageUrl string `form:"image_url" json:"image_url" binding:"required"`
	Content  string `form:"content" json:"content" binding:"required"`
}
