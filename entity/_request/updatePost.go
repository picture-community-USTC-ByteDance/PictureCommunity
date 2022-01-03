package _request

type UpdatePost struct {
	ID      int64    `form:"u_id" json:"u_id" binding:"required"`
	Content UserInfo `form:"content" json:"content" binding:"required"`
}
