package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"picture_community/entity/_request"
	"picture_community/response"
	"picture_community/service"
)

func UpdateUserInfo(c *gin.Context) {
	var u _request.UpdatePost

	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -2,
			"message": err.Error(),
		})
		return
	}
	fmt.Println("controller: ", u.ID, u.Content)
	status, message := service.UpdateUserInfo(u)
	c.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": message,
	})
}

func EmailIsUniqueController(c *gin.Context) {
	var u _request.EmailIsUniqueInfo
	if err := c.ShouldBind(&u); err != nil {
		response.CheckFail(c, nil, "检查email参数错误")
		return
	}

	isOK, message := service.EmailIsUnique(u)
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

	isOK, message := service.TelephoneIsUnique(u)
	if isOK {
		response.Success(c, nil, message)
	} else {
		response.Fail(c, nil, message)
	}
}
