package _request

type Firstpage struct {
	PageSize int `form:"pagesize" json:"pagesize" binding:"required"`
	Page     int `form:"page"     json:"page"      binding:"required"`
}

type SinglePost struct {
	Pid uint `form:"pid"      json:"pid"       binding:"required"`
}
