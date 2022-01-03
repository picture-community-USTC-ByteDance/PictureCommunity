package controller

import (
	"github.com/gin-gonic/gin"
	"picture_community/entity/_request"
	"picture_community/response"
	"picture_community/service"
)

func CreateLikeController(c *gin.Context) {
	//var u db.Liked
	var u _request.LikeRequest
	err := c.ShouldBind(&u)
	uid, _ := c.Get("uid")
	if err != nil {
		response.CheckFail(c, nil, "Invalid parameter")
		return
	}
	res := service.CreateLike(uid.(uint), u.PostID)
	if res {
		response.Success(c, nil, "ok")
		return
	}
	response.Fail(c, nil, "creat error")
}
func QueryLikeController(c *gin.Context) {
	var u _request.LikeRequest
	err := c.ShouldBind(&u)
	uid, _ := c.Get("uid")
	if err != nil {
		response.CheckFail(c, nil, "Invalid parameter")
		return
	}
	res := service.QueryLike(uid.(uint), u.PostID)
	//res := service.QueryLike(u.UID, u.PostID)
	if res {
		response.Success(c, nil, "ok")
		return
	}
	response.Fail(c, nil, "query like error")
}

func CancelLikeController(c *gin.Context) {
	//var u db.Liked
	var u _request.LikeRequest
	err := c.ShouldBind(&u)
	uid, _ := c.Get("uid")
	if err != nil {
		response.CheckFail(c, nil, "Invalid parameter")
		return
	}
	res := service.CancelLike(uid.(uint), u.PostID)
	if res {
		response.Success(c, nil, "ok")
		return
	}
	response.Fail(c, nil, "cancel like error")
}
