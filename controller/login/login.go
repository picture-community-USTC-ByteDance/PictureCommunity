package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picture_community/entity"
	"picture_community/service/login"
)

func LoginController(c *gin.Context) {
	var u entity.LoginUser
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -2,
			"message": err.Error(),
			"token":   "",
		})
		return
	}
	status, message, token := login.VerifyLogin(u)
	c.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": message,
		"token":   token,
	})
}
