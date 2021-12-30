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
