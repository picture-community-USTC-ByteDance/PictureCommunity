package post

import (
	"picture_community/entity/db"
	"picture_community/global"
)

func InsertPostByUserID(newPost db.Post) (int64, error) {
	err := global.MysqlDB.Create(&newPost).Error
	return int64(newPost.PID), err
}
