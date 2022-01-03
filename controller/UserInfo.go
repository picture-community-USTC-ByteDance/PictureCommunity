package controller

import (
	"github.com/gin-gonic/gin"
	"picture_community/response"
	"picture_community/service"
)

// 获取用户主页个人信息
func QueryUserData(c *gin.Context) {
	// 根据用户id查找用户数据信息
	uid, exists := c.Get("uid")
	if exists == false {
		response.Fail(c, nil, "用户不存在！")
	}

	res := service.GetUserData(uid.(int)) // 实际处理业务逻辑的函数

	response.HandleResponse(c, res)

}

// 获取用户的帖子数组信息
func QueryUserPosts(c *gin.Context) {
	uid, exists := c.Get("uid")
	if exists == false {
		response.Fail(c, nil, "用户不存在！")
	}

	res := service.GetUserPosts(uid.(int)) // 实际处理业务逻辑的函数

	response.HandleResponse(c, res)
}
