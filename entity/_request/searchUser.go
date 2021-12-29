package _request

type SearchUsers struct {
	NickName string `form:"nickname"  binding:"required"`
	PageSize int    `form:"pagesize"  binding:"required"`
	Page     int    `form:"page"  binding:"required"`
}
