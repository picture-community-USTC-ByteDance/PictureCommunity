package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"picture_community/entity/_request"
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
