package _request

type RegisterUserInfo struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UsernameIsUniqueInfo struct {
	Username string `form:"username" json:"username" binding:"required"`
}
