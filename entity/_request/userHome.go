package _request

type UserHomeRequest struct {
	Uid      uint `form:"uid"      json:"uid"       binding:"required"`
	PageSize int  `form:"pagesize" json:"pagesize" binding:"required"`
	Page     int  `form:"page"     json:"page"      binding:"required"`
}
