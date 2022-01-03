package user

import (
	"picture_community/entity/db"
	"picture_community/global"
)

func QueryIDAndPasswordByUsername(username string) (int64, string, error) {
	var user db.User
	err := global.MysqlDB.Select("uid", "password").Where("username=?", username).First(&user).Error
	return int64(user.UID), user.Password, err
}

func QueryIDAndPasswordByEmail(email string) (int64, string, error) {
	var user db.User
	err := global.MysqlDB.Select("uid", "password").Where("email=?", email).First(&user).Error
	return int64(user.UID), user.Password, err
}

func QueryIDAndPasswordByTelephone(telephone string) (int64, string, error) {
	var user db.User
	err := global.MysqlDB.Select("uid", "password").Where("telephone=?", telephone).First(&user).Error
	return int64(user.UID), user.Password, err
}
