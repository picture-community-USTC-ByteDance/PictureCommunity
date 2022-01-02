package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"picture_community/entity/db"
	"picture_community/global"
)

func MysqlDateBaseInit() error {
	database, err := gorm.Open(mysql.Open(global.DbUrl), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	database.AutoMigrate(&db.Post{})
	database.AutoMigrate(&db.UserDetail{})
	database.AutoMigrate(&db.Fans{})
	database.AutoMigrate(&db.Comment{})
	database.AutoMigrate(&db.User{})
	database.AutoMigrate(&db.Forward{})
	database.AutoMigrate(&db.Liked{})
	database.AutoMigrate(&db.Follow{})
	database.AutoMigrate(&db.UserData{})
	database.AutoMigrate(&db.Collection{})

	global.MysqlDB = database
	return err
}
