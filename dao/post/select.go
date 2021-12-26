package post

import (
	"fmt"
	"picture_community/entity"
	"picture_community/global"
)

func InsertPostByUserID(newPost entity.Post) (int64, error) {
	fmt.Println(newPost.TitlePhotoUrl)
	err := global.MYSQL_DB.Create(&newPost).Error
	return newPost.ID, err
}
