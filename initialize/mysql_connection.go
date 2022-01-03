package initialize

import (
	"fmt"
	"picture_community/entity/db"
	"picture_community/global"
	"picture_community/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func MysqlDateBaseInit() error {
	dbUrl := utils.GetDbUrl()
	fmt.Println("using DATABASE: ", dbUrl)
	database, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}

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
