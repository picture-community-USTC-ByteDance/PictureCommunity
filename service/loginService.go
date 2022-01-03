package service

import (
	"errors"
	"gorm.io/gorm"
	"picture_community/dao/user"
	"picture_community/entity/_request"
	"picture_community/utils"
)

func VerifyLogin(param _request.LoginUser) (isValid bool, message string, token string) {
	var id int64
	var password string
	var err error

	message = "密码错误"
	isValid = false
	token = ""

	switch param.Method {
	case 0:
		id, password, err = user.QueryIDAndPasswordByUsername(param.Info)
	case 1:
		id, password, err = user.QueryIDAndPasswordByEmail(param.Info)
	case 2:
		id, password, err = user.QueryIDAndPasswordByTelephone(param.Info)
	default:
		message = "格式错误"
		return
	}

	if err != nil {
		isValid = false
		if errors.Is(err, gorm.ErrRecordNotFound) {
			message = "账号不存在"
			return
		} else {
			message = "数据库错误"
			return
		}
	}

	if password == param.Password {
		isValid = true
		message = "登录成功"
		token, _ = utils.CreateToken(id)
	}
	return
}
