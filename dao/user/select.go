package user

import (
	"picture_community/entity/db"
	"picture_community/global"
)

func QueryIDAndPasswordByUsername(username string) (id int64, password string) {
	var user db.User
	global.MysqlDB.Select("id", "password").Where("username=?", username).First(&user)
	return int64(user.UID), user.Password
}

func QueryIDAndPasswordByEmail(email string) (id int64, password string) {
	var user db.User
	global.MysqlDB.Select("id", "password").Where("email=?", email).First(&user)
	return int64(user.UID), user.Password
}

func QueryIDAndPasswordByTelephone(telephone string) (id int64, password string) {
	var user db.User
	global.MysqlDB.Select("id", "password").Where("telephone=?", telephone).First(&user)
	return int64(user.UID), user.Password
}

func QueryUserData(uid int) {
	var user_data db.UserData
	global.MysqlDB.Where("uid=?", uid).First(&user_data)

}
