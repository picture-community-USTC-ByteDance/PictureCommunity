package main

import (
	"fmt"
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
