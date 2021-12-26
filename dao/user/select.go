package user

import (
	"database/sql"
	"picture_community/global"
)

var db *sql.DB

func QueryIDAndPasswordByUsername(username string) (id int64, password string) {
	global.MYSQL_DB.DB().QueryRow("select u_id,password from user where username=?", username).Scan(&id, &password)
	return
}

func QueryIDAndPasswordByEmail(email string) (id int64, password string) {
	global.MYSQL_DB.DB().QueryRow("select u_id,password from user where email=?", email).Scan(&id, &password)
	return
}

func QueryIDAndPasswordByTelephone(telephone string) (id int64, password string) {
	global.MYSQL_DB.DB().QueryRow("select u_id,password from user where telephone=?", telephone).Scan(&id, &password)
	return
}
