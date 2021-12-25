package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"picture_community/global"
	"picture_community/initialize"
)

func MysqlDateBaseInit() error {
	db, err := gorm.Open("mysql", global.DbUrl)
	if err == nil {
		db.DB().SetMaxIdleConns(200)
		db.DB().SetMaxIdleConns(200)
	}
	fmt.Println(db)
	global.MYSQL_DB = db
	return err
}

func main() {
	fmt.Println("----------System Start-----------")
	if err := MysqlDateBaseInit(); err != nil {
		fmt.Printf("database error : %v\n", err)
		return
	}
	fmt.Println(global.MYSQL_DB)
	initialize.RunSystemWithGIN()
}
