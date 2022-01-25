package controller

import (
	"github.com/gin-gonic/gin"
	"picture_community/entity/_request"
	"picture_community/response"
	"picture_community/service"
)

// 获取用户自己的主页个人信息
func QueryUserData(c *gin.Context) {
	// 根据用户id查找用户数据信息
	uid, exists := c.Get("uid")
	if exists == false {
		response.Fail(c, nil, "用户不存在！")
	}

	err, res := service.GetUserData(uid.(uint)) // 实际处理业务逻辑的函数
	if err != nil {
		response.Fail(c, nil, "查找用户主页信息失败")
	} else {
		response.Success(c, gin.H{"userdata": res}, "ok")
	}

	//response.HandleResponse(c, res)

}

// 获取用户自己个人主页的帖子数组信息
func QueryUserPosts(c *gin.Context) {
	var u _request.UserPosts
	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "请求错误")
		return
	}
	if u.Page <= 0 || u.PageSize <= 0 {
		response.CheckFail(c, nil, "页码或数量有误")
		return
	}

	uid, exists := c.Get("uid")
	if exists == false {
		response.Fail(c, nil, "用户不存在！")
	}

	err, count, totalPage, res := service.GetUserPosts(uid.(uint), u) // 实际处理业务逻辑的函数
	if count == 0 {
		response.Success(c, nil, "当前未发表帖子")
	}
	if err != nil {
		response.Fail(c, nil, "查看用户主页帖子数组失败")
	} else {
		response.Success(c, gin.H{"totalpage": totalPage, "userposts": res}, "ok")
	}

	//response.HandleResponse(c, res)
}
