package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"picture_community/initialize"
)

func main() {
	fmt.Println("----------System Start-----------")
	if err := initialize.MysqlDateBaseInit(); err != nil {
		fmt.Printf("database error : %v\n", err)
		return
	}
	initialize.RunSystemWithGIN()
}
