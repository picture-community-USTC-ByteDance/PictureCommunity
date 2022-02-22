package controller

import (
	"picture_community/entity/_request"
	"picture_community/response"
	"picture_community/service"

	"github.com/gin-gonic/gin"
)

func UserFollowController(c *gin.Context) {
	var u _request.FollowUserRequest
	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "Invalid parameter")
		return
	}
	uid, _ := c.Get("uid")
	res := service.FollowUser(uid.(uint), u.UID)
	response.HandleResponse(c, res)
}

func UserUnfollowController(c *gin.Context) {
	var u _request.FollowUserRequest
	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "Invalid parameter")
		return
	}
	uid, _ := c.Get("uid")
	res := service.UnfollowUser(uid.(uint), u.UID)
	response.HandleResponse(c, res)
}
