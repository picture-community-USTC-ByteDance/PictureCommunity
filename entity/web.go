package entity

type LoginUser struct {
	Info     string `form:"info" json:"info" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Method   int8   `form:"method" json:"method"`
}
