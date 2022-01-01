package like

import (
	"fmt"
	"picture_community/entity/db"
	"picture_community/global"
)

func QueryLikeByUserID(u_id uint, post_id uint) (bool, error) {
	var like db.Liked
	fmt.Println(u_id, post_id)
	err := global.MysqlDB.Select("State").Where("ToLikePostID=? and FromUserID=?", post_id, u_id).First(&like).Error
	return like.State, err
}
