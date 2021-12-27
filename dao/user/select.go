package user

import (
	"picture_community/entity/db"
	"picture_community/global"
)

func QueryIDAndPasswordByUsername(username string) (id int64, password string) {
	var user db.User
	global.MysqlDB.Select("uid", "password").Where("username=?", username).First(&user)
	return user.UID, user.Password
}

func QueryIDAndPasswordByEmail(email string) (id int64, password string) {
	var user db.User
	global.MysqlDB.Select("uid", "password").Where("email=?", email).First(&user)
	return user.UID, user.Password
}

func QueryIDAndPasswordByTelephone(telephone string) (id int64, password string) {
	var user db.User
	global.MysqlDB.Select("uid", "password").Where("telephone=?", telephone).First(&user)
	return user.UID, user.Password
}
