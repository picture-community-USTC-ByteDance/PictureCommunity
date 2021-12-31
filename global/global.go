package global

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DbUrl   = "root:123456@(121.5.1.73:3306)/pic?charset=utf8mb4&parseTime=True"
	MysqlDB *gorm.DB

	GinEngine *gin.Engine
)
