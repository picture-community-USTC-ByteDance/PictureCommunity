package initialize

import (
	"github.com/jinzhu/gorm"
	"picture_community/global"
)

func MysqlDateBaseInit() error {
	db, err := gorm.Open("mysql", global.DbUrl)
	if err == nil {
		db.DB().SetMaxIdleConns(200)
		db.DB().SetMaxIdleConns(200)
	}
	global.MYSQL_DB = db

	return err
}
