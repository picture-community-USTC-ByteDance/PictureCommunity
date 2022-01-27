package user

import (
	"picture_community/entity/db"
	"picture_community/global"
)

func UpdateUserDetailByID(uid uint, userDetail db.UserDetail) error {
	err := global.MysqlDB.Model(&userDetail).Where("uid = ?", uid).Updates(userDetail).Error
	return err
}

func UpdateEmailByID(uid uint, email string) error {
	err := global.MysqlDB.Model(&db.User{}).Where("uid = ?", uid).Update("email", email).Error
	return err
}

func UpdateTelephoneByID(uid uint, telephone uint) error {
	err := global.MysqlDB.Model(&db.User{}).Where("uid = ?", uid).Update("telephone", telephone).Error
	return err
}

func UpdatePasswordByID(uid uint, password string) error {
	return global.MysqlDB.Model(&db.User{}).Where("uid = ?", uid).Update("password", password).Error
}
