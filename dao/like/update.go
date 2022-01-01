package like

import (
	"fmt"
	"picture_community/entity/db"
	"picture_community/global"
)

func InsertLikeByUserID(newLike db.Liked) (int64, error) {
	fmt.Println(newLike.ToLikePostID)
	err := global.MysqlDB.Create(&newLike).Error
	return int64(newLike.FromUserID), err
}
func CancelLikeByUserID(u_id uint, post_id uint) bool {
	var like db.Liked
	fmt.Println(u_id, post_id)
	global.MysqlDB.Where("ToLikePostID=? and FromUserID=?", post_id, u_id).First(&like)
	global.MysqlDB.Update("State", true)
	return like.State
}
