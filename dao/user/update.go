package user

import (
	"picture_community/entity/db"
	"picture_community/global"
)

func UpdateUserDetailByID(uid uint, userDetail db.UserDetail) error {
	err := global.MysqlDB.Model(&userDetail).Where("uid = ?", uid).Updates(userDetail).Error
	return err
}
