package _request

type Firstpage struct {
	Uid      uint `form:"uid"      json:"uid"       binding:"required"`
	PageSize int  `form:"pagesize" json:"pagesize" binding:"required"`
	Page     int  `form:"page"     json:"page"      binding:"required"`
}

type SinglePost struct {
	Uid uint `form:"uid"      json:"uid"       binding:"required"`
	Pid uint `form:"pid"      json:"pid"       binding:"required"`
}
