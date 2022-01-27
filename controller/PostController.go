package controller

import (
	"picture_community/entity/_request"
	"picture_community/response"
	"picture_community/service"

	"github.com/gin-gonic/gin"
)

func CreatePostController(c *gin.Context) {
	var u _request.CreatePost
	err := c.ShouldBind(&u)
	if err != nil {
		response.CheckFail(c, nil, "Invalid parameter")
		return
	}

	uid, _ := c.Get("uid")
	res := service.CreatePost(c, uid.(uint), u.Url, u.Content)
	response.HandleResponse(c, res)
}

func NewForwardController(c *gin.Context) {
	var u _request.NewForward
	err := c.ShouldBind(&u)
	if err != nil {
		response.CheckFail(c, nil, "Invalid parameter")
		return
	}
	uid, _ := c.Get("uid")
	res := service.NewForward(uid.(uint), u.PID, u.Content)
	response.HandleResponse(c, res)
}
