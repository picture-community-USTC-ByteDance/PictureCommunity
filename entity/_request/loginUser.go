package _request

type LoginUser struct {
	Info     string `form:"info" json:"info" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Method   int8   `form:"method" json:"method"`
}

type ModifyPasswd struct {
	OldPassword string `json:"old_password" form:"old_password" binding:"required"`
	NewPassword string `json:"new_password" form:"new_password" binding:"required"`
}
