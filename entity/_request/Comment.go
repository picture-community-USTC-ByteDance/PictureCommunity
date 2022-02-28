package _request

type CreateFirstLevelComment struct {
	PID     uint   `form:"p_id" json:"p_id" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}
type CreateSecondLevelComment struct {
	PID      uint   `form:"p_id" json:"p_id" binding:"required"`
	ParentId uint   `form:"parent_id" json:"parent_id" binding:"required"`
	Content  string `form:"content" json:"content" binding:"required"`

}
type DeleteComment struct {
	CID uint `form:"c_id" json:"c_id" binding:"required"`
}
type QueryComment struct {
	PID      uint `form:"p_id" json:"p_id" binding:"required"` //帖子id
	PageSize int  `form:"pagesize" json:"page_size"  binding:"required"`
	Page     int  `form:"page" json:"page" binding:"required"`
}

type QueryComment2 struct {
	PID      uint `form:"p_id" json:"p_id" binding:"required"`
	ParentId uint `form:"parent_id" json:"parent_id"`
	PageSize int  `form:"pagesize" json:"page_size" binding:"required"`
	Page     int  `form:"page" json:"page" binding:"required"`
}
