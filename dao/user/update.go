package user

import (
	"picture_community/entity/db"
	"picture_community/global"
)

func UpdateUserDetailByID(userDetail db.UserDetail, Id int64) error {
	res := global.MysqlDB.Create(userDetail)
	return res.Error
}
