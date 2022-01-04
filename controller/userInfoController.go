package controller

import (
	"github.com/gin-gonic/gin"
	"picture_community/entity/_request"
	"picture_community/response"
	"picture_community/service"
)

func UpdateUserDetailController(c *gin.Context) {
	var u _request.UpdateUserDetailInfo

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

func UpdateUserEmailController(c *gin.Context) {
	var u _request.UpdateUserEmailInfo
	if err := c.ShouldBind(&u); err != nil {
		response.CheckFail(c, nil, "更新email参数错误")
		return
	}
	uid, _ := c.Get("uid")
	isOK, message := service.UpdateUserEmailService(u, uid.(uint))
	if isOK {
		response.Success(c, nil, message)
	} else {
		response.Fail(c, nil, message)
	}
	return
}

func UpdateUserTelephoneController(c *gin.Context) {
	var u _request.UpdateUserTelephoneInfo
	if err := c.ShouldBind(&u); err != nil {
		response.CheckFail(c, nil, "更新电话号码参数错误")
		return
	}

	uid, _ := c.Get("uid")
	isOK, message := service.UpdateUserTelephoneService(u, uid.(uint))
	if isOK {
		response.Success(c, nil, message)
	} else {
		response.Fail(c, nil, message)
	}
	return
}
