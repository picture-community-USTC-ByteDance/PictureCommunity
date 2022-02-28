package _request

type MessageRequest struct {
	PageSize int `form:"pagesize" json:"pagesize" binding:"required"`
	Page     int `form:"page"     json:"page"      binding:"required"`
}
