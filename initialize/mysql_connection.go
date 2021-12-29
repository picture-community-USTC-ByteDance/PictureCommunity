package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"picture_community/entity/db"
	"picture_community/global"
)

func MysqlDateBaseInit() error {
	database, err := gorm.Open(mysql.Open(global.DbUrl), &gorm.Config{})

	database.AutoMigrate(&db.Post{})
	database.AutoMigrate(&db.UserDetail{})
	global.MysqlDB = database

	return err
}
