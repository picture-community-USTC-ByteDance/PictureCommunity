package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"picture_community/dao/user"
	"picture_community/entity/_request"
	"picture_community/entity/db"
	"picture_community/global"
)

func RegisterService(param _request.RegisterUser) (isOK bool, message string) {
	_, _, err := user.QueryIDAndPasswordByUsername(param.Username)
	if err == nil {
		return false, "用户名重复"
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println(err)
		return false, "数据库错误"
	}
	newUser := db.User{
		UID:      global.UserIDGenerator.NewID(),
		Username: param.Username,
		Password: param.Password,
		Status:   0,
	}
	err = user.InsertUser(newUser)
	if err != nil {
		fmt.Println(err)
		return false, "注册失败"
	} else {
		return true, "注册成功"
	}
}

func UsernameIsUniqueService(param _request.UsernameIsUniqueInfo) (isOK bool, message string) {
	//todo 用户名合法性检查
	_, _, err := user.QueryIDAndPasswordByUsername(param.Username)

	if err == nil {
		return false, "用户名重复"
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return true, "用户名可用"
	} else {
		fmt.Println(err)
		return false, "数据库错误"
	}
}
