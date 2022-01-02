package global

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"picture_community/utils"
)

var (
	DbUrl = "root:123456@(121.5.1.73:3306)/pic?charset=utf8mb4&parseTime=True"
	//DbUrl   = "root:root@(127.0.0.1:3306)/test2?charset=utf8mb4&parseTime=True"
	MysqlDB *gorm.DB

	GinEngine *gin.Engine

	UserIDGenerator    *utils.IDGenerator
	PostIDGenerator    *utils.IDGenerator
	ForwardIDGenerator *utils.IDGenerator
	CommentIDGenerator *utils.IDGenerator
)
