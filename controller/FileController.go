package controller

import (
	"picture_community/response"
	"picture_community/service"

	"github.com/gin-gonic/gin"
)

func FileUploadController(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.CheckFail(c, nil, "File not found")
		return
	}
	uid, _ := c.Get("uid")
	res := service.FileUpload(c, uid.(uint), file)
	response.HandleResponse(c, res)
}
