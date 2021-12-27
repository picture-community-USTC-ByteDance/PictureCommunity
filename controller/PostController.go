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
		response.Fail(c, nil, "Invalid parameter")
		return
	}

	res := service.CreatePost(u)
	response.HandleResponse(c, res)
}
