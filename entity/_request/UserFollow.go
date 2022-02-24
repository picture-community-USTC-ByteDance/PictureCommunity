package _request

type FollowUserRequest struct {
	UID uint `form:"u_id" json:"u_id" binding:"required"`
}
