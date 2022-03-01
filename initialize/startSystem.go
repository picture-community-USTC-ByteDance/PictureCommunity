package initialize

import (
	"github.com/gin-gonic/gin"
	"picture_community/controller"
	"picture_community/global"
	"picture_community/router"
)

func RunSystemWithGIN() {
	go controller.Manager.Start()
	r := gin.Default()
	global.GinEngine = r
	router.SetRouter()
	r.Run(":8080")
}
