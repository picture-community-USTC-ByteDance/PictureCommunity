package controller

import (
	"github.com/gin-gonic/gin"
	"picture_community/entity/_request"
	"picture_community/response"
	"picture_community/service"
)

func UpdateUserController(c *gin.Context) {
	var u _request.UpdateUserDetail

	if err := c.ShouldBind(&u); err != nil {
		response.CheckFail(c, nil, "更新用户信息参数错误")
	}
	uid, _ := c.Get("uid")
	isOK, message := service.UpdateUserDetailService(u, uid.(uint))
	if isOK {
		response.Success(c, nil, message)
	} else {
		response.Fail(c, nil, message)
	}
}

func EmailIsUniqueController(c *gin.Context) {
	var u _request.EmailIsUniqueInfo
	if err := c.ShouldBind(&u); err != nil {
		response.CheckFail(c, nil, "检查email参数错误")
		return
	}

	isOK, message := service.EmailIsUniqueService(u)
	if isOK {
		response.Success(c, nil, message)
	} else {
		response.Fail(c, nil, message)
	}
	return
}

func TelephoneIsUniqueController(c *gin.Context) {
	var u _request.TelephoneIsUniqueInfo
	if err := c.ShouldBind(&u); err != nil {
		response.CheckFail(c, nil, "检查电话号参数错误")
		return
	}

	isOK, message := service.TelephoneIsUniqueService(u)
	if isOK {
		response.Success(c, nil, message)
	} else {
		response.Fail(c, nil, message)
	}
}
