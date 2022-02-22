package controller

import (
	"github.com/gin-gonic/gin"
	"picture_community/entity/_request"
	"picture_community/entity/_response"
	"picture_community/response"
	"picture_community/service"
)

func LoginController(c *gin.Context) {
	var u _request.LoginUser
	if err := c.ShouldBind(&u); err != nil {
		response.CheckFail(c, nil, "参数错误")
		return
	}
	status, message, token := service.VerifyLogin(u)
	if status {
		response.Success(c, _response.Token{Token: token}, message)
	} else {
		response.Fail(c, nil, message)
	}
}

func UpdatePassword(c *gin.Context) {
	uid, exists := c.Get("uid")
	if exists == false {
		response.Fail(c, nil, "用户不存在！")
	}

	var u _request.ModifyPasswd
	if err := c.ShouldBind(&u); err != nil {
		response.CheckFail(c, nil, "参数错误")
		return
	}

	flag, msg := service.ModifyPassword(uid.(uint), u)

	if flag {
		response.Success(c, nil, msg)
	} else {
		response.Fail(c, nil, msg)
	}
}
