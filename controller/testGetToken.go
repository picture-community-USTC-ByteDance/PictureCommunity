package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picture_community/utils"
	"strconv"
)

func GetToken(c *gin.Context) {
	id := c.Query("id")
	uid, _ := strconv.ParseInt(id, 10, 64)
	c.JSON(http.StatusOK, gin.H{
		"token": utils.CreateToken(uid),
	})
}
