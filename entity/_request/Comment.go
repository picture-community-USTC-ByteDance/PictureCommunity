package _request

type CreateFirstLevelComment struct {
	PID      uint   `form:"p_id" json:"p_id" binding:"required"`
	UserID   uint   `form:"u_id" json:"u_id" binding:"required"`
	Content  string `form:"content" json:"content" binding:"required"`
	Nickname string `form:"nickname" json:"nickename"`
	Profile  string `form:"profile" json:"profile"`
}
type CreateSecondLevelComment struct {
	PID      uint   `form:"p_id" json:"p_id" binding:"required"`
	UserID   uint   `form:"u_id" json:"u_id" binding:"required"`
	ParentId uint   `form:"parent_id" json:"parent_id" binding:"required"`
	Content  string `form:"content" json:"content" binding:"required"`
	Nickname string `form:"nickname" json:"nickename"`
	Profile  string `form:"profile" json:"profile"`
}
type DeleteComment struct {
	CID uint `form:"c_id" json:"c_id" binding:"required"`
	//PID      uint	 	`form:"p_id" json:"p_id" binding:"required"`
	//ParentId uint 		`form:"parent_id" json:"parent_id" binding:"required"`
}
type QueryComment struct {
	PID      uint `form:"p_id" json:"p_id" binding:"required"` //帖子id
	PageSize int  `form:"pagesize"  binding:"required"`
	Page     int  `form:"page"  binding:"required"`
}

type QueryComment2 struct {
	PID      uint `form:"p_id" json:"p_id" binding:"required"`
	ParentId uint `form:"parentid" json:"parentid"`
	PageSize int  `form:"pagesize"  binding:"required"`
	Page     int  `form:"page"  binding:"required"`
}
