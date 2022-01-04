package _request

type LikeRequest struct {
	PostID uint `json:"post_id" form:"post_id"`
	//UID    uint `json:"uid" form:"uid"`
}
