package _request

type UpdateUserDetailInfo struct {
	Nickname string `form:"nickname" json:"nickname"`
	//false为女  true为男
	Sex      bool   `form:"sex" json:"sex"`
	Birthday string `form:"birthday" json:"birthday"`
	Address  string `form:"address" json:"address"`
	Motto    string `form:"motto" json:"motto"`
	//略缩图
	Profile string `form:"profile" json:"profile"`
	//详细头像url
	OriginProfile string `form:"origin_profile" json:"origin_profile"`
}

type EmailIsUniqueInfo struct {
	Email string `form:"email" json:"email" binding:"required"`
}

type TelephoneIsUniqueInfo struct {
	Telephone uint `form:"telephone" json:"telephone" binding:"required"`
}

type UpdateUserEmailInfo struct {
	Email string `form:"email" json:"email" binding:"required"`
}

type UpdateUserTelephoneInfo struct {
	Telephone uint `form:"telephone" json:"telephone" binding:"required"`
}
