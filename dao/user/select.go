package user

import (
	"picture_community/entity/_response"
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

func QuerySearchUsersByNickname(keywords string, page int, pageSize int) (int64, []_response.ResponseSearchUsers) {
	var searchUsers []_response.ResponseSearchUsers
	var count int64
	global.MysqlDB.Offset((page-1)*pageSize).Limit(pageSize).Model(db.UserDetail{}).
		Select("profile,uid,nickname,motto").
		Joins("inner join user on user.uid = user_detail.uid").
		Where("nickname like ?", keywords).Count(&count).Scan(&searchUsers)
	return count, searchUsers
}
