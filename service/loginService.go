package service

import (
	"picture_community/dao/user"
	"picture_community/entity/_request"
	"picture_community/utils"
)

func VerifyLogin(param _request.LoginUser) (isValid int, message string, token string) {
	var id int64
	var password string

	message = "密码错误"
	isValid = -1
	token = ""

	switch param.Method {
	case 0:
		id, password = user.QueryIDAndPasswordByUsername(param.Info)
	case 1:
		id, password = user.QueryIDAndPasswordByEmail(param.Info)
	case 2:
		id, password = user.QueryIDAndPasswordByTelephone(param.Info)
	default:
		message = "格式错误"
		return
	}
	if password != "" && password == param.Password {
		isValid = 0
		message = "登录成功"
		token = utils.CreateToken(id)
	}
	return
}
