package _request

type UserPosts struct {
	PageSize int `form:"pagesize" json:"pagesize" binding:"required"`
	Page     int `form:"page"     json:"page"      binding:"required"`
}

type UserInfo struct {
	Username string `form:"username" json:"username" binding:"required"`
}
