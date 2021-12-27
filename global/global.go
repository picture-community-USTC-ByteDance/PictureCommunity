package global

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	//DbUrl    = "root:root@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True" 本地测试数据库
	DbUrl    = "root:123456@(121.5.1.73:3306)/pic?charset=utf8mb4&parseTime=True"
	MYSQL_DB *gorm.DB

	GinEngin *gin.Engine
)
