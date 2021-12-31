package controller

import (
	"github.com/gin-gonic/gin"
	"picture_community/entity/_request"
	"picture_community/response"
	"picture_community/service"
)

func CreatePostController(c *gin.Context) {
	var u _request.CreatePost
	err := c.ShouldBind(&u)
	if err != nil {
		response.CheckFail(c, nil, "Invalid parameter")
		return
	}
	//if !VerifyIDByToken(u.ID, u.Token) {
	//	response.CheckFail(c, nil, "Invalid Token")
	//	return
	//}
	res := service.CreatePost(u)
	response.HandleResponse(c, res)
}
