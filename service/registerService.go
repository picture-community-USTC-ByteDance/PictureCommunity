package service

import (
	"fmt"
	"picture_community/dao/user"
	"picture_community/entity/_request"
	"picture_community/entity/db"
	"picture_community/global"
)

func RegisterService(param _request.RegisterUser) (isOK bool, message string) {
	UID, _ := user.QueryIDAndPasswordByUsername(param.Username)
	if UID != 0 {
		return false, "用户名重复"
	}
	newUser := db.User{
		UID:      global.UserIDGenerator.NewID(),
		Username: param.Username,
		Password: param.Password,
		Status:   0,
	}
	err := user.InsertUser(newUser)
	if err != nil {
		fmt.Println(err)
		return false, "注册失败"
	} else {
		return true, "注册成功"
	}
}

func UsernameIsUnique(param _request.UsernameIsUniqueInfo) (isOK bool, message string) {
	//todo 用户名合法性检查
	ID, _ := user.QueryIDAndPasswordByUsername(param.Username)

	if ID != 0 {
		return false, "用户名重复"
	} else {
		return true, "用户名可用"
	}
}
