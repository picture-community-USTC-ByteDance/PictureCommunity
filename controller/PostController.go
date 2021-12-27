package controller

import (
	"picture_community/entity"
	"picture_community/response"
	"picture_community/service"

	"github.com/gin-gonic/gin"
)

func CreatePostController(c *gin.Context) {
	var u entity.CreatePost
	err := c.ShouldBind(&u)
	if err != nil {
		response.CheckFail(c, nil, "Invalid parameter")
		return
	}
	if !VerifyIDByToken(u.ID, u.Token) {
		response.CheckFail(c, nil, "Invalid Token")
		return
	}
	res := service.CreatePost(u)
	response.HandleResponse(c, res)
}
