package initialize

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"picture_community/controller"
	"picture_community/global"
	"picture_community/router"
)

func RunSystemWithGIN() {
	go controller.Manager.Start()
	app := gin.Default()
	global.GinEngine = app

	pprof.Register(app) //性能监控

	router.SetRouter()
	app.Run(":8080")
}
