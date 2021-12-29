package initialize

import (
	"picture_community/entity"
	"picture_community/global"

	"github.com/jinzhu/gorm"
)

func MysqlDateBaseInit() error {
	db, err := gorm.Open("mysql", global.DbUrl)
	if err == nil {
		db.DB().SetMaxIdleConns(200)
	}
	db.AutoMigrate(entity.Post{})
	db.AutoMigrate(&entity.Follow{})
	global.MYSQL_DB = db

	return err
}
