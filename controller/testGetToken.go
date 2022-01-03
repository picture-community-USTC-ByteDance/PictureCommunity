package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picture_community/utils"
	"strconv"
)

func GetToken(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"token": utils.CreateToken(int64(id)),
	})
}
