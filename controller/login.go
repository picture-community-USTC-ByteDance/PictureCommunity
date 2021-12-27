package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picture_community/entity/_request"
	"picture_community/service"
)

func LoginController(c *gin.Context) {
	var u _request.LoginUser
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -2,
			"message": err.Error(),
			"token":   "",
		})
		return
	}
	status, message, token := service.VerifyLogin(u)
	c.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": message,
		"token":   token,
	})
}
