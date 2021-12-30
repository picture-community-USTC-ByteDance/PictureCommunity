package post

import (
	"fmt"
	"picture_community/entity/db"
	"picture_community/global"
)

func InsertPostByUserID(newPost db.Post) (int64, error) {
	fmt.Println(newPost.TitlePhotoUrl)
	err := global.MysqlDB.Create(&newPost).Error
	return int64(newPost.UID), err
}
