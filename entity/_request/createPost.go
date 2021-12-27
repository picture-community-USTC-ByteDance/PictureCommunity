package _request

type CreatePost struct {
	ID       int64  `form:"u_id" json:"u_id" binding:"required"`
	ImageUrl string `form:"image_url" json:"image_url" binding:"required"`
	Content  string `form:"content" json:"content" binding:"required"`
	Token    string `form:"token" json:"token" binding:"required"`
}

type SearchUsers struct {
	NickName string `form:"nickname"  binding:"required"`
	PageSize int    `form:"pagesize"  binding:"required"`
	Page     int    `form:"page"  binding:"required"`
}
