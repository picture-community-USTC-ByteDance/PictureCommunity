package _request

type UpdatePost struct {
	ID int64 `form:"u_id" json:"u_id" binding:"required"`
	//UpdateTime time.Time `form:"update_time" json:"update_time" binding:"required"`
	Content UserInfo `form:"content" json:"content" binding:"required"`
}
