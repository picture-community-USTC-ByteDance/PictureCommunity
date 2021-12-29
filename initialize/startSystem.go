package initialize

import (
	"github.com/gin-gonic/gin"
	"picture_community/global"
	"picture_community/router"
)

func RunSystemWithGIN() {
	r := gin.Default()
	global.GinEngine = r
	router.SetRouter()
	r.Run(":8080")
}
