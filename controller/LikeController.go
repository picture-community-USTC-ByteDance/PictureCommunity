package controller

import (
	"github.com/gin-gonic/gin"
	"picture_community/entity/db"
	"picture_community/response"
	"picture_community/service"
)

func CreateLikeController(c *gin.Context) {
	var u db.Liked
	err := c.ShouldBind(&u)
	if err != nil {
		response.CheckFail(c, nil, "Invalid parameter")
		return
	}
	res := service.CreateLike(u.FromUserID, u.ToLikePostID)
	response.HandleResponse(c, res)
}
func QueryLikeController(c *gin.Context) {
	var u db.Liked
	err := c.ShouldBind(&u)
	if err != nil {
		response.CheckFail(c, nil, "Invalid parameter")
		return
	}
	res := service.QueryLike(u.FromUserID, u.ToLikePostID)
	response.HandleResponse(c, res)
}

func CancelLikeController(c *gin.Context) {
	var u db.Liked
	err := c.ShouldBind(&u)
	if err != nil {
		response.CheckFail(c, nil, "Invalid parameter")
		return
	}
	res := service.CancelLike(u.FromUserID, u.ToLikePostID)
	response.HandleResponse(c, res)
}
